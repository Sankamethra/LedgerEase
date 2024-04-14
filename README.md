# LedgerEase Project

LedgerEase is a comprehensive ledger system that integrates a React front-end, a Python machine learning model, an Ignite blockchain, and a Node.js back-end. This system allows users to securely store and manage data in a tamper-proof blockchain ledger.

## Project Overview

The LedgerEase project consists of several components:

- **React Front-end:** The user interface for interacting with the ledger system.
- **Python Machine Learning Model:** Responsible for data extraction and processing.
- **Ignite Blockchain:** Provides a secure and immutable ledger for storing data.
- **Node.js Back-end:** Handles server-side logic and communication with the blockchain.

## Prerequisites

Before getting started with LedgerEase, ensure you have the following installed:

- Node.js (for running the React front-end and the back-end)
- Python (for running the machine learning model)
- Ignite Blockchain (for the ledger)

## Getting Started

Follow these steps to set up and run the LedgerEase project:

1. **Machine Learning Model**
   - Navigate to `LedgerEase/ml-model`.
   - Run `python file.py` to start the ML model. This model will handle data extraction and processing.

2. **Ignite Blockchain**
   - Start the blockchain by navigating to `LedgerEase/ignite-chain/invoice`.
   - Build the binaries using `make build-linux`.
   - Move the binaries to a local directory: `sudo cp -r invoiced /usr/local/bin`.
   - Run the blockchain with `ignite chain serve`.

3. **React Front-end**
   - Navigate to `LedgerEase/UI`.
   - Install dependencies with `npm install`.
   - Start the React app with `npm run start`. This will launch the user interface for interacting with the ledger.

4. **Node.js Back-end**
   - Restart the MongoDB service with `sudo systemctl restart mongod`.
   - Start the Node.js server by navigating to `LedgerEase/server` and running `node server.js`. This server handles communication with the blockchain and the machine learning model.

## Usage

1. Open the React app in your browser at [http://localhost:3000](http://localhost:3000).
2. Sign up or log in with valid credentials to access the ledger functionalities.
3. Navigate to the file upload page and upload your desired file.
4. Click the upload button to trigger the data extraction process using the ML model.
5. Confirm to store the extracted data in the blockchain when prompted.
6. Once the transaction is successful, view the stored data in the blockchain on the "Show Transactions" page.

## Viewing Stored Data

To check the data stored in the blockchain, use the following command:

```bash
cd LedgerEase/ignite-chain/invoice
invoiced q invoice list-invoice --chain-id=invoice
```

## Notes

-  The React front-end runs on port 3000 by default.
-  Ensure all necessary services (MongoDB, Ignite, ML model) are running before using the application.
-  Make sure to have the correct credentials for signing up and logging in to access the application features.