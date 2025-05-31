# pnwxkd — Proven National Workers Xero-Knowledge Daemon (aka "XERØ")

**pnwxkd**, also known as **"XERØ"**, is a Cosmos SDK–based sovereign sidechain that powers governance, compliance, and zk-audit infrastructure for the Proven National Workers (PNW) ecosystem. It complements the Aleo-based `pnw_mvp` backbone, serving as a modular layer for cross-chain zero-knowledge credential validation, SubDAO oversight, and IBC interoperability.

---

## 🌐 Architecture Overview

       [ Aleo: pnw_mvp ]
        - zk payroll proofs
        - credential NFTs
        - .pnw identity registry
               ||
               ||  (zk proof, hash, or IBC relay)
               \/
   [ Cosmos SDK: pnwxkd ("Xero") ]
    - zk credential indexer
    - payroll audit log mirror
    - DAO + SubDAO governance
    - .pnw naming registry checkpoint
    - Treasury + penalty enforcement

---

## 🔧 Key Modules

### ✅ zkcredential
- Validates and stores zk credential hashes from Aleo
- Supports public hash indexing and revocation mapping
- Allows third-party verification via REST/gRPC

### ✅ namingregistry
- Mirrors `.pnw` identity mappings from Aleo
- Enables IBC-readable lookup of human-readable identities

### ✅ payrollaudit
- Cross-chain logging of payroll hashes, timestamps, and states
- Designed for audit compliance and SubDAO traceability

### ✅ dao
- SubDAO proposal voting and penalty system
- Funds allocation and reward enforcement logic

---

## 🪙 Future Utility: The XERØ Token (`XKRO`)

The **XERØ token** (ticker: `XKRO`) will be introduced in future phases as a programmable, pegged-value wage token.

- Designed as a **CBDC-style instrument**, pegged to local fiat currencies (e.g., USD, MXN, BRL)
- Used as a **cross-border wage token** for Proven National Workers
- Enables **stable-value payroll payouts** in a privacy-preserving manner
- May be used for staking, DAO incentives, and remittance disbursement

Governance of the XERØ token will reside within the `pnwxkd` sidechain, with issuance and peg mechanics transparent and regulated by SubDAOs.

---

## 🧠 Design Principles

- ✅ **Sovereign** — Full control and upgradeability using Cosmos SDK  
- 🔐 **Privacy-Aware** — Accepts and validates zk artifacts from Aleo  
- 🔗 **Interoperable** — Designed to bridge Aleo with XRPL, Ethereum, Solana, etc.  
- 📊 **Transparent** — On-chain audit logs and credential registries  

---

## 📄 Licensing

This software is proprietary and owned by the Proven National Workers Foundation.  
Use of the name `Xero-Knowledge™` and any component of this software requires prior written approval.  
See `LICENSE.txt` and `COPYRIGHT.txt` for details.

---

## ✉️ Contact

- Email: `provennationalworkers@gmail.com`  
- Aleo Project: [`pnw_mvp`](https://github.com/PNWBNW/pnw_mvp)  
- Cosmos Chain ID: `pnwxkd-1` (testnet coming soon)

---

Built for zk governance.  
Secured by Aleo.  
Extended by Cosmos.  
Powered by XERØ (`XKRO`).
