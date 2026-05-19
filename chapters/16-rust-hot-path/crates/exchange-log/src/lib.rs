use std::fs::{File, OpenOptions};
use std::io::{BufRead, BufReader, BufWriter, Write};
use std::path::{Path, PathBuf};

use exchange_types::CommandEnvelope;
use thiserror::Error;

#[derive(Debug)]
pub struct CommandLog {
    path: PathBuf,
}

impl CommandLog {
    pub fn open(path: impl Into<PathBuf>) -> Self {
        Self { path: path.into() }
    }

    pub fn append(&self, command: &CommandEnvelope) -> Result<(), LogError> {
        let file = OpenOptions::new()
            .create(true)
            .append(true)
            .open(&self.path)?;
        let mut writer = BufWriter::new(file);
        serde_json::to_writer(&mut writer, command)?;
        writer.write_all(b"\n")?;
        writer.flush()?;
        Ok(())
    }

    pub fn read_all(&self) -> Result<Vec<CommandEnvelope>, LogError> {
        read_commands(&self.path)
    }
}

pub fn read_commands(path: impl AsRef<Path>) -> Result<Vec<CommandEnvelope>, LogError> {
    let file = File::open(path)?;
    let reader = BufReader::new(file);
    let mut commands = Vec::new();

    for line in reader.lines() {
        let line = line?;
        if line.trim().is_empty() {
            continue;
        }
        commands.push(serde_json::from_str(&line)?);
    }

    Ok(commands)
}

#[derive(Debug, Error)]
pub enum LogError {
    #[error("io error: {0}")]
    Io(#[from] std::io::Error),
    #[error("json error: {0}")]
    Json(#[from] serde_json::Error),
}

#[cfg(test)]
mod tests {
    use super::*;
    use exchange_types::{
        AccountId, Amount, Asset, Command, CommandEnvelope, Price, Quantity, Seq,
    };

    #[test]
    fn append_only_log_round_trips_commands_in_order() {
        let dir = tempfile::tempdir().unwrap();
        let log = CommandLog::open(dir.path().join("commands.jsonl"));

        let first = CommandEnvelope {
            seq: Seq(1),
            command: Command::Deposit {
                account_id: AccountId(1),
                asset: Asset::Usd,
                amount: Amount(1_000_000),
            },
        };
        let second = CommandEnvelope {
            seq: Seq(2),
            command: Command::PlaceLimit {
                client_order_id: exchange_types::ClientOrderId(1001),
                account_id: AccountId(1),
                instrument_id: exchange_types::InstrumentId(1),
                side: exchange_types::Side::Buy,
                price: Price(10_000),
                quantity: Quantity(1),
            },
        };

        log.append(&first).unwrap();
        log.append(&second).unwrap();

        assert_eq!(log.read_all().unwrap(), vec![first, second]);
    }
}
