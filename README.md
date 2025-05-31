# Xero
Xero Knowledge system for xrpl

# Xero-Knowledge XRPL

**Xero-Knowledge** is a modular zero-knowledge toolkit designed to simplify zk-based identity, payroll, and compliance use cases. Built on the XRPL EVM Sidechain, this framework provides developer-friendly zk credential verification using Solidity, Circom, and modern web3 tooling.

---

## 🌐 What It Does

- ✅ Verifies zk credentials with on-chain smart contracts  
- 🔐 Supports privacy-first proof systems (SnarkJS, Circom)  
- ⚙️ Bridges traditional data into zk formats (identity, payroll)  
- 🧩 Interoperable with XRPL mainnet via Axelar bridge  

---

## 🛠️ Requirements

- Node.js (v18+)
- Circom + SnarkJS
- Hardhat (for contract deployment)
- MetaMask (for XRPL EVM Sidechain)

---

## ⚙️ Setup Instructions

### 1. Install and Compile Circom Circuit

```bash
cd circuits
circom credential_verification.circom --r1cs --wasm --sym
snarkjs groth16 setup credential_verification.r1cs pot12_final.ptau circuit_final.zkey
snarkjs zkey export verificationkey circuit_final.zkey verification_key.json

> You can generate a dummy .zkey if you just want to test frontend flows.




---

2. Deploy Smart Contract to XRPL EVM Sidechain

cd contracts
npx hardhat compile
npx hardhat run scripts/deploy.js --network xrplEVM


---

3. Launch Frontend

cd frontend
npm install
npm run start

Visit http://localhost:5173 in your browser.


---

🧠 Use Cases

zkPayroll: proof of hours worked without exposing identity

zkKYC: validate age, citizenship, residency via circuit

zkCompliance: cryptographic attestations without leaks



---

🔒 License

This project is licensed under the MIT License.


---

© Trademark Notice

XERO-KNOWLEDGE™ is a protected title under common law trademark.

All rights to the name "Xero-Knowledge" in relation to zero-knowledge systems are claimed by the Proven National Workers Foundation, 2025.

See COPYRIGHT.txt for more details.


---

💬 Questions?

Contact: founder@pnw.foundation
Website: coming soon
Twitter: @PNW_MVP (tentative)


---

Built for transparency. Powered by zero-knowledge.




