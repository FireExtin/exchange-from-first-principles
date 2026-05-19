use exchange_types::{CommandEnvelope, EventEnvelope};
use thiserror::Error;

pub trait ExchangeStateMachine {
    fn apply(&mut self, command: CommandEnvelope) -> Result<Vec<EventEnvelope>, CoreError>;

    fn snapshot(&self) -> Result<Vec<u8>, CoreError>;

    fn load_snapshot(&mut self, snapshot: &[u8]) -> Result<(), CoreError>;
}

#[derive(Debug, Error)]
pub enum CoreError {
    #[error("invalid command: {0}")]
    InvalidCommand(&'static str),
    #[error("snapshot error: {0}")]
    Snapshot(&'static str),
}

#[cfg(test)]
mod tests {
    use super::*;
    use exchange_types::{Command, Seq};

    struct RecordingMachine {
        seen: Vec<Seq>,
    }

    impl ExchangeStateMachine for RecordingMachine {
        fn apply(&mut self, command: CommandEnvelope) -> Result<Vec<EventEnvelope>, CoreError> {
            self.seen.push(command.seq);
            Ok(Vec::new())
        }

        fn snapshot(&self) -> Result<Vec<u8>, CoreError> {
            Ok(self
                .seen
                .iter()
                .flat_map(|seq| seq.0.to_le_bytes())
                .collect())
        }

        fn load_snapshot(&mut self, _snapshot: &[u8]) -> Result<(), CoreError> {
            Ok(())
        }
    }

    #[test]
    fn state_machine_contract_accepts_ordered_commands() {
        let mut machine = RecordingMachine { seen: Vec::new() };

        machine
            .apply(CommandEnvelope {
                seq: Seq(1),
                command: Command::Cancel {
                    account_id: exchange_types::AccountId(1),
                    client_order_id: exchange_types::ClientOrderId(10),
                },
            })
            .unwrap();

        assert_eq!(machine.seen, vec![Seq(1)]);
    }
}
