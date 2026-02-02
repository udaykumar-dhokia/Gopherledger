# Gopherledger

Gopherledger is a simple, CLI-based expense tracker built with Go. It helps you manage your daily expenses directly from your terminal.

## Installation

To install Gopherledger, ensure you have Go installed on your machine. Then, clone the repository and build the binary:

```bash
git clone https://github.com/udaykumar-dhokia/Gopherledger.git
cd Gopherledger
go build -o gl
```

You can move the `gl` binary to a directory in your system's PATH to use it globally.

## Usage

### Add an Expense

Use the `add` command to record a new expense. You must specify the amount using the `--amount` (or `-a`) flag. Optionally, you can add a note with `--note` (or `-n`).

```bash
# Add an expense of 100
./gl add --amount 100

# Add an expense with a note
./gl add --amount 50 --note "Coffee"

# Using shorthand flags
./gl add -a 25.50 -n "Taxi"
```

### Delete an Expense

Use the `delete` command to remove an expense by its ID. You can find the ID of an expense when you add it or by (future feature: listing expenses).

```bash
# Delete transaction with ID 123
./gl delete --id 123

# Using shorthand flag
./gl delete -i 123
```

## Storage

Expenses are stored locally in a JSON file.
