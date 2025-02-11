# agent-battle

## Flow
### Token Transfer Flow

#### Start Game (Tweet ID)
1. Create wallets based on user input Twitter names (verify Twitter names) (A)
2. Create master wallet based on game ID (Twitter ID) (B)

#### End Game
- Triggered by timeout
- Update bet records (C) based on transfer events from each wallet A
- Update total bet amount (D) for each wallet A
- Update total bet amount for game (E)

#### Submit Result
1. Gas Distribution
  - Faucet gas to wallet A (only enough for 1 transfer)
  - Faucet gas to wallet B (pre-calculate needed transfers and request sufficient gas + buffer)

2. Token Collection
  - Transfer all D from A to B (if exists)

3. Winner Scenarios
  - With Winners:
    1. Retain x% from total E sum and transfer to treasury wallet (game fee)
    2. Calculate token amount to transfer to each winner from remaining balance after fee (F)
    3. Transfer all remaining tokens to winners according to F
  - No Winners/No Participants:
    - Refund each record C using wallet B (can combine multiple bets per wallet into single transfer)

#### Diagrams
- [user flow](docs/diagrams/userflow.puml)

## Test

```bash
go test ./... -v
```