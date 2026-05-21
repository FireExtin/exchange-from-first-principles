# Change Brief 06: Credit, Margin, And Funding Extension

Date: 2026-05-21

## Summary

Added a planned credit, margin, and funding extension before the trading desk
extension.

The project still does not become a Modern Lending product. The new note only
borrows lending concepts that strengthen exchange semantics:

- collateral and borrow-liability accounts;
- funding and interest accrual boundaries;
- repayment and collateral release;
- liquidation settlement;
- bad-debt and insurance-fund transfers.

## Key Changes

- Added `docs/09-credit-margin-funding-extension.md`.
- Moved the trading desk note to `docs/10-trading-desk-extension.md`.
- Reordered planned notes so credit/margin/funding occupies 18-20 and trading
  desk ideas occupy 21-25.
- Updated the documentation map, chapter roadmap, root README, and design
  paper to match the new order.

## Guardrails

- No business implementation was added.
- No new chapter directories were created for the planned notes.
- Credit/margin/funding remains an exchange-core extension, not a separate
  lending marketplace track.
- Older change briefs remain historical and were not rewritten.
