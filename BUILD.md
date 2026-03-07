# 🛠️ How to Build QuantumChain

This guide explains how to compile the QuantumChain Layer-1 node from source.

## Prerequisites
- Go 1.21+
- Linux (Ubuntu/Debian recommended) or macOS
- GCC Compiler

## Build Instructions
1. Clone the repository:
   ```bash
   git clone [https://github.com/irlan7/quantumpay-docs.git](https://github.com/irlan7/quantumpay-docs.git)
   cd quantumpay-docs
Download dependencies:

Bash
go mod tidy
Compile the node binary:

Bash
go build -o quantum-node ./cmd/settlement-engine
Verify installation:

Bash
./quantum-node version
