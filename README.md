# pnwxkd — Proven National Workers Xero-Knowledge Daemon

**pnwxkd** is a Cosmos SDK–based sovereign sidechain that powers the governance, compliance, and zk-audit infrastructure for the Proven National Workers (PNW) ecosystem. It complements the Aleo-based `pnw_mvp` backbone, serving as a modular layer for cross-chain zero-knowledge credential validation, SubDAO oversight, and IBC interoperability.

## 🌐 Architecture Overview

       [ Aleo: pnw_mvp ]
        - zk payroll proofs
        - credential NFTs
        - .pnw identity registry
               ||
               ||  (zk proof, hash, or IBC relay)
               \/
   [ Cosmos SDK: pnwxkd ]
    - zk credential indexer
    - payroll audit log mirror
    - DAO + SubDAO governance
    - ANS registry checkpoint
    - Treasury + penalty enforcement

## 📁 Directory Structure

pnwxkd/
├── app/                          # Cosmos app wiring
│   └── app.go
├── cmd/                          # CLI entrypoint
│   └── pnwxkd/
│       └── main.go
├── x/                            # Custom Cosmos modules
│   ├── zkcredential/             # ZK credential verifier + Aleo sync
│   ├── namingregistry/              # Mirrors .pnw Aleo name service
│   ├── payrollaudit/             # Indexes payroll hashes from Aleo
│   ├── dao/                      # SubDAO and oversight logic
│   └── governance/               # Proposal engine (optional)
├── proto/                        # Protobuf definitions
│   └── pnwxkd/
├── scripts/                      # IBC tools, relayers, genesis setup
├── test/                         # Integration test environment
├── docs/                         # Developer and IBC guides
├── LICENSE.txt                   # Proprietary or hybrid license
├── COPYRIGHT.txt                 # Trademark protection
├── README.md                     # This file
├── go.mod                        # Go module setup
└── go.sum                        # Go module lock

## 🔧 Key Modules

### ✅ zkcredential
- Validates and stores zk credential hashes from Aleo
- Supports public hash indexing and revocation mapping
- Allows third-party verification via REST/gRPC

### ✅ namingregistry
- Mirrors .pnw identity mappings from Aleo
- Enables IBC-readable lookup of human-readable identities

### ✅ paytrrollaudit
- Cross-chain logging of payroll hashes, timestamps, and states
- Designed for audit compliance and SubDAO traceability

### ✅ dao
- SubDAO proposal voting and penalty system
- Funds allocation and reward enforcement logic

## 🧠 Design Principles

- ✅ Sovereign: Built using Cosmos SDK for full control  
- 🔐 Privacy-Aware: Accepts and validates zk artifacts from Aleo  
- 🔗 Interoperable: Designed to bridge across XRPL, Ethereum, Solana, and more  
- 📊 Transparent: On-chain audit support for zero-knowledge proofs  

## 📄 Licensing

This software is proprietary and owned by the Proven National Workers Foundation.  
Use of the name `Xero-Knowledge™` and any component of this software requires prior written approval.  
See `LICENSE.txt` and `COPYRIGHT.txt` for details.

## ✉️ Contact

- Email: `provennationalworkers@gmail.com`  
- Aleo Project: [`pnw_mvp`](https://github.com/PNWBNW/pnw_mvp)  
- Cosmos Chain ID: `pnwxkd-1` (testnet coming soon)

---

Built for zk governance.  
Secured by Aleo.  
Extended by Cosmos.
