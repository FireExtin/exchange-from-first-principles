use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord, Hash, Serialize, Deserialize)]
#[serde(transparent)]
pub struct Seq(pub u64);

#[derive(Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord, Hash, Serialize, Deserialize)]
#[serde(transparent)]
pub struct AccountId(pub u64);

#[derive(Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord, Hash, Serialize, Deserialize)]
#[serde(transparent)]
pub struct InstrumentId(pub u32);

#[derive(Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord, Hash, Serialize, Deserialize)]
#[serde(transparent)]
pub struct OrderId(pub u64);

#[derive(Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord, Hash, Serialize, Deserialize)]
#[serde(transparent)]
pub struct ClientOrderId(pub u64);

#[derive(Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord, Hash, Serialize, Deserialize)]
#[serde(transparent)]
pub struct Price(pub i64);

#[derive(Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord, Hash, Serialize, Deserialize)]
#[serde(transparent)]
pub struct Quantity(pub i64);

#[derive(Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord, Hash, Serialize, Deserialize)]
#[serde(transparent)]
pub struct Amount(pub i64);

#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash, Serialize, Deserialize)]
#[serde(rename_all = "snake_case")]
pub enum Side {
    Buy,
    Sell,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash, Serialize, Deserialize)]
#[serde(rename_all = "UPPERCASE")]
pub enum Asset {
    Usd,
    Btc,
}

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
#[serde(tag = "type", rename_all = "snake_case")]
pub enum Command {
    Deposit {
        account_id: AccountId,
        asset: Asset,
        amount: Amount,
    },
    PlaceLimit {
        client_order_id: ClientOrderId,
        account_id: AccountId,
        instrument_id: InstrumentId,
        side: Side,
        price: Price,
        quantity: Quantity,
    },
    Cancel {
        account_id: AccountId,
        client_order_id: ClientOrderId,
    },
}

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct CommandEnvelope {
    pub seq: Seq,
    pub command: Command,
}

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
#[serde(tag = "type", rename_all = "snake_case")]
pub enum Event {
    Deposited {
        account_id: AccountId,
        asset: Asset,
        amount: Amount,
    },
    OrderAccepted {
        order_id: OrderId,
        client_order_id: ClientOrderId,
        account_id: AccountId,
    },
    OrderRejected {
        client_order_id: ClientOrderId,
        reason: RejectReason,
    },
    TradeExecuted {
        maker_order_id: OrderId,
        taker_order_id: OrderId,
        price: Price,
        quantity: Quantity,
    },
    OrderRested {
        order_id: OrderId,
        remaining: Quantity,
    },
    OrderCancelled {
        client_order_id: ClientOrderId,
    },
    CancelRejected {
        client_order_id: ClientOrderId,
        reason: RejectReason,
    },
}

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct EventEnvelope {
    pub seq: Seq,
    pub event: Event,
}

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
#[serde(rename_all = "snake_case")]
pub enum RejectReason {
    DuplicateClientOrderId,
    InsufficientBalance,
    UnknownOrder,
    InvalidQuantity,
    InvalidPrice,
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn command_contract_round_trips_as_json() {
        let command = CommandEnvelope {
            seq: Seq(42),
            command: Command::PlaceLimit {
                client_order_id: ClientOrderId(1001),
                account_id: AccountId(7),
                instrument_id: InstrumentId(1),
                side: Side::Buy,
                price: Price(10_000),
                quantity: Quantity(3),
            },
        };

        let encoded = serde_json::to_string(&command).unwrap();
        let decoded: CommandEnvelope = serde_json::from_str(&encoded).unwrap();

        assert_eq!(decoded, command);
    }

    #[test]
    fn event_contract_round_trips_as_json() {
        let event = EventEnvelope {
            seq: Seq(43),
            event: Event::TradeExecuted {
                maker_order_id: OrderId(1),
                taker_order_id: OrderId(2),
                price: Price(10_000),
                quantity: Quantity(1),
            },
        };

        let encoded = serde_json::to_string(&event).unwrap();
        let decoded: EventEnvelope = serde_json::from_str(&encoded).unwrap();

        assert_eq!(decoded, event);
    }
}
