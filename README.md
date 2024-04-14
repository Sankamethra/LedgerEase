# LedgerEase Project


LedgerEase is a comprehensive invoice management system that integrates a ReactJS front-end, a Python machine learning model, a Cosmos blockchain, and a NodeJS back-end. This system allows users to securely store and manage data in a tamper-proof blockchain ledger.



## Project Overview

The LedgerEase project consists of several components:

- *ReactJS Front-end:* The user interface for interacting with the system.

- *Azure's Document Intelligence AI:* Responsible for the extraction of necessary fields from the invoices. 

- *Cosmos Blockchain:* Provides a secure and immutable ledger for storing data.

- *NodeJS Back-end:* Handles server-side logic and communication with the blockchain.



## Prerequisites

Before getting started with LedgerEase, ensure you have the following installed:

- NodeJS (for running the React front-end and the back-end)

- Python (for running the machine learning model)

- Ignite CLI (for the ledger)

- Golang v1.21 or greater.


### To use the Invoice prebuilt model of Azure, 

- Create an Azure account if not already done.

- Proceed to the Azure portal and create the necessary resources for Document Intelligence AI.

- Follow the documentation and tutorials provided by Azure for setting up and configuring the Document Intelligence service.

      - Reference: https://learn.microsoft.com/en-us/azure/synapse-analytics/machine-learning/tutorial-form-recognizer-use-mmlspark

- Explore the Document Intelligence Studio, we've used the Invoice prebuilt model which is designed specifically for processing invoices.

- You can add / remove fields as per your requirements.

- It supports many languages, we've utilized Python.

- Utilize the provided API endpoint and secret key to integrate it to the system.



## Getting Started

Follow these steps to set up and run the LedgerEase project:

1. *Machine Learning Model*

   - Navigate to LedgerEase/ml-model.

   - Run python3 file.py to start the ML model. This model will handle data extraction and processing.


2. *Cosmos Blockchain*

   - Start the blockchain by navigating to LedgerEase/ignite-chain/invoice.

   - Build the chain binaries using make build-linux.

   - Move the binaries to bin directory: sudo cp -r invoiced /usr/local/bin.

   - Run the blockchain using the command ignite chain serve.


3. *React Front-end*

   - Navigate to LedgerEase/UI.

   - Install necessary dependencies and packages with npm install.

   - Start the React app with npm run start. This will launch the user interface for interacting with the ledger.


4. *Node.js Back-end*

   - Start the MongoDB service with sudo systemctl start mongod.service.

   - Start the NodeJS server by navigating to LedgerEase/server and running node server.js. 

   This server handles communication with the blockchain and the machine learning model.



## Usage

1. Open the React app in your browser at [http://localhost:3000](http://localhost:3000).

2. Sign up or log in with valid credentials to access the ledger functionalities.

3. Navigate to the file upload page and upload any invoice image. 

4. Click the upload button to trigger the data extraction process using the ML model.

5. Confirm to store the extracted data in the blockchain when prompted.

6. Once the transaction is successful, view the stored data in the blockchain on the "Show Transactions" page.



## Viewing Stored Data

To verify the data stored in the blockchain, use the following command:

```bash
cd LedgerEase/ignite-chain/invoice
invoiced q invoice list-invoice --chain-id=invoice
```

## Note

-  The React front-end runs on port 3000 by default.

-  Ensure all necessary services (MongoDB, Ignite chain, ML model) are running before using the application.



## Future works

- Access control mechanism has to be implemented so that only authorized administrators should be able to view the entire transaction history.

- Invoices from same company has to be grouped together for easy accessibility.

- Reminder notifications can be sent to the admin according to the due date of the invoices.

- The transaction data can be used to perform Profit and Loss, or other data analysis to benefit the company.

- Get feedbacks from real users and tailor the application according to the needs of Indian market.