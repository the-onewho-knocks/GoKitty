# GoKitty

> Your Go code has bugs. GoKitty has opinions.

GoKitty is a static analysis tool for Go that detects code smells, concurrency issues, architecture problems, security risks, and testing gaps before they reach production.

## Features

- Ignored error detection
- Goroutine leak detection
- Channel deadlock warnings
- WaitGroup misuse detection
- Mutex misuse detection
- Architecture health scoring
- Dead code detection
- SQL injection risk detection
- Missing test detection
- Panic stack trace explanation
- Developer ranking
- Code roasting

## Installation

```bash
go install github.com/yourusername/gokitty/cmd/gokitty@latest
```

## Usage

```bash
gokitty check .

gokitty score .

gokitty judge .

gokitty roast .

gokitty duck

gokitty explain panic.log
```

## Example

```text
$ gokitty check .

GoKitty Health Report

Health Score: 82/100
Rank: Go Adventurer

Ignored Errors: 12
Concurrency Risks: 2
Large Functions: 4
Missing Tests: 7

Verdict:
Your code is brave.
Perhaps too brave.
```

## Philosophy

Good code ships.
Better code survives production.

##
