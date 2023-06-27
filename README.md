# community-assistant-cli

A CLI tool to help with community event organisation.

## Build the CLI binary

```bash
go build -o community-assist
```

## Run the CLI

```bash
./community-assist <command> <subcommand> [flags]
```

### Commands

#### `help`

```bash
./community-assist help
```

#### `tweet`

Allows you to create a tweet formulated text from an input text file and save it out to a file.

```bash
./community-assist tweet --input <input-file> --output <output-file>
```

## Tests

```bash
go test ./...
```

