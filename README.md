# Maxis Flip Bot

Maxis Flip Bot is an open-source automation script written in Go for automating flips on the Maxis platform. The script
supports two modes of operation: Basic and Advanced, catering to different user needs and expertise levels.

## Features

- `Basic Mode`: Allows running the bot with single wallet configuration directly via command line arguments.
- `Advanced Mode`: Supports multiple wallet configurations through a JSON file for complex operations and multiple bets.

## Compiled Binaries For Download

The script is pre-compiled for multiple platforms. You can find the appropriate binary for your system:

- `flip_bot_amd64` for Linux 64-bit.
- `flip_bot_amd64.exe` for Windows 64-bit.
- `flip_bot_mac-arm64` for macOS on Apple M-series chips.

## Usage

### Basic Mode

Allows running the bot with single wallet configuration directly via command line arguments.

#### Example Command

For macOS (Apple M-series chips):

```shell
$ ./flip_bot_mac-arm64 -pk yourWalletAddress -sk yourPrivateKey -amount 0.5 -bet 0.002
```

For Linux (x86-64):

```shell
$ flip_bot_amd64.exe -pk yourWalletAddress -sk yourPrivateKey -amount 0.5 -bet 0.002
```

For Windows (x86-64):

```shell
$ ./flip_bot_amd64 -pk yourWalletAddress -sk yourPrivateKey -amount 0.5 -bet 0.002
```

In Basic Mode, you can directly use the compiled binary with the following parameters:

- `pk`: Wallet address to perform the flips.
- `sk`: Private key of the wallet.
- `amount`: Total amount to run (script stops after exceeding this amount).
- `bet`: Amount to bet per flip.

### Advanced Mode

Advanced Mode reads configurations from a config.json file, allowing detailed customization and running multiple
wallets. You can specify the configuration file location with the -config parameter. If not specified, the default
location ./config.json is used.

#### Example Command

For macOS (Apple M-series chips):

```shell
$ ./flip_bot_mac-arm64 -config /Users/xxx/config.json
```

For Linux (x86-64):

```shell
$ flip_bot_amd64.exe -config /Users/xxx/config.json
```

For Windows (x86-64):

```shell
$ ./flip_bot_amd64 -config /Users/xxx/config.json
```

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
      "pk": "yourWalletAddress1",
      "sk": "yourPrivateKey1",
      "payAmount": 5
    },
    {
      "pk": "yourWalletAddress2",
      "sk": "yourPrivateKey2",
      "payAmount": 10
    }
  ]
}
```

- `betSelectProbability` allows setting the probability for each betting amount, ensuring the sum of all prob values
  equals 100.
- `wallets` allows configuration of multiple wallets, where pk is the public key, sk is the private key, and payAmount
  is the total amount intended for the run.

## Compiling Binaries From Source

If you want to compile the Maxis Flip Bot binaries yourself to ensure they are up to date or to make modifications, you
can do so by using the Go compiler with specific environment settings for each target platform. Here’s how you can
compile the binaries for different operating systems:

### Requirements

- Go installed on your system (version 1.15 or higher is recommended).
- Access to a command-line interface.

### Compilation Commands

Navigate to the directory containing the source code and run the following commands based on your target operating
system:

For macOS (Apple M-series chips):

```shell
$ GOOS=darwin GOARCH=arm64 go build -o flip_bot_mac-arm64 main.go
```

For Linux (x86-64):

```shell
$ GOOS=linux GOARCH=amd64 go build -o flip_bot_amd64 main.go
```

For Windows (x86-64):

```shell
$ GOOS=windows GOARCH=amd64 go build -o flip_bot_amd64.exe main.go
```

## Installation

To use Maxis Flip Bot, download the appropriate binary for your operating system and run it using the command line as
demonstrated in the usage section.

## Contributing

Contributions to Maxis Flip Bot are welcome. Please ensure to follow the best practices and code of conduct when making
contributions.