# Xero-Knowledge XRPL

**Xero-Knowledge** is a modular zero-knowledge toolkit designed to simplify zk-based identity, payroll, and compliance use cases. Built on the XRPL EVM Sidechain, this framework provides developer-friendly zk credential verification using Solidity, Circom, and modern web3 tooling.

---

## 🌐 Overview

This repository provides a scaffold for building zk-enabled apps on XRPL using Solidity, Circom, and modern web3 tooling.

---

## 🔧 Setup Instructions

### Requirements

- Node.js (v18+)
- Circom + SnarkJS
- Hardhat
- MetaMask

---

### 1. Compile Circuit

```bash
cd circuits
circom credential_verification.circom --r1cs --wasm --sym
snarkjs groth16 setup credential_verification.r1cs pot12_final.ptau circuit_final.zkey
snarkjs zkey export verificationkey circuit_final.zkey verification_key.json

> You can generate a dummy .zkey if you just want to test frontend flows.




---

2. Deploy Smart Contract

cd contracts
npx hardhat compile
npx hardhat run scripts/deploy.js --network xrplEVM


---

3. Run Frontend

cd frontend
npm install
npm start

Visit http://localhost:5173 in your browser.


---

✨ Features

Plug-and-play zk credential verification

Interoperable with XRPL EVM sidechain

Easily extendable to payroll, voting, or compliance use cases



---

📄 LICENSE

This project is proprietary software and may not be copied, modified, distributed, or used without explicit written permission from the Proven National Workers Foundation. See LICENSE.txt.


---

© COPYRIGHT

See COPYRIGHT.txt for trademark and naming protections.


---

💬 Contact

To request a commercial license or partnership: Email: provennationalworkers@gmail.com
X: coming soon
Site: coming soon


---

Built for transparency. Powered by zero-knowledge.



