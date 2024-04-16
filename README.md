# Maxis Flip Bot
Maxis Flip Bot is an open-source automation script written in Go for automating flips on the Maxis platform. The script supports two modes of operation: Basic and Advanced, catering to different user needs and expertise levels.

## Features

- `Basic Mode`: Allows running the bot with single wallet configuration directly via command line arguments.
- `Advanced Mode`: Supports multiple wallet configurations through a JSON file for complex operations and multiple bets.

## Compiled Binaries
The script is pre-compiled for multiple platforms. You can find the appropriate binary for your system:

- `flip_bot_amd64` for Linux 64-bit.
- `flip_bot_amd64` exe for Windows 64-bit.
- `flip_bot_mac-arm64` for macOS on Apple M-series chips.

## Usage

### Basic Mode
In Basic Mode, you can directly use the compiled binary with the following parameters:
- `pk`: Wallet address to perform the flips.
- `sk`: Private key of the wallet.
- `amount`: Total amount to run (script stops after exceeding this amount).
- `bet`: Amount to bet per flip.
#### Example Command

```
$ flip_bot_mac-arm64 -pk 0x000 -sk xxxx -amount 0.5 -bet 0.002
```

### Advanced Mode
Advanced Mode reads configurations from a config.json file, allowing detailed customization and running multiple wallets.

#### Configuration File: `config.json`
#### Below is the structure of the `config.json` file:

```json
{
  "betSelectProbability": [
    {
      "betSelect": 0.002,
      "prob": 100
    },
    {
      "betSelect": 0.005,
      "prob": 0
    },
    {
      "betSelect": 0.01,
      "prob": 0
    },
    {
      "betSelect": 0.02,
      "prob": 0
    },
    {
      "betSelect": 0.05,
      "prob": 0
    },
    {
      "betSelect": 0.1,
      "prob": 0
    }
  ],
  "wallets": [
    {
      "pk": "",
      "sk": "",
      "payAmount": 10
    }
  ]
}
```

- `betSelectProbability` allows setting the probability for each betting amount, ensuring the sum of all prob values equals 100.
- `wallets` allows configuration of multiple wallets, where pk is the public key, sk is the private key, and payAmount is the total amount intended for the run.

#### Example Command
```
$ flip_bot_mac-arm64
```

## Installation
To use Maxis Flip Bot, download the appropriate binary for your operating system and run it using the command line as demonstrated in the usage section.

## Contributing
Contributions to Maxis Flip Bot are welcome. Please ensure to follow the best practices and code of conduct when making contributions.