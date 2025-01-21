# LedgerEase

### Introduction
LedgerEase is a cutting-edge solution aimed at transforming invoice management for small and medium-sized enterprises (SMEs). By leveraging advanced technologies such as Azure Document Intelligence AI and Cosmos Blockchain, it automates invoice processing, ensures data integrity, and provides a secure, user-friendly platform for seamless management.

---

## Conceptual Study of the Project

In today’s digital era, traditional paper-based invoices and manual transaction records pose significant operational challenges for businesses. Manual data entry is not only time-consuming but also error-prone, with risks of data tampering and fraud. 

**LedgerEase** offers a transformative approach by:
- Automating invoice data extraction using **Azure Document Intelligence AI**.
- Securing and decentralizing data storage with **Cosmos Blockchain**.
- Enhancing transparency and oversight through user-friendly web interfaces with tailored access controls.

---

## **Features**

1. **Automated Invoice Processing:**
   - Extracts essential fields such as invoice number, date, customer name, due date, and total price using Azure AI.

2. **Decentralized Data Storage:**
   - Stores sensitive invoice data on Cosmos Blockchain for enhanced security and immutability.

3. **Access Control:**
   - Admins have full access to view and manage all invoices.
   - Employees can upload invoices with restricted access.

4. **Real-Time Alerts:**
   - Enables quick decision-making with timely notifications.

5. **Web-Based Interface:**
   - Provides an intuitive and seamless user experience for managing invoices.

---

## **Folder Structure**

```plaintext
LedgerEase/
|-- ignite_chain/             # Cosmos Blockchain module
|   |-- invoice/              # Blockchain implementation
|-- ml-model/                 # Azure AI integration
|   |-- main.py               # Main script for invoice processing
|   |-- images/               # Sample invoice images
|-- server/                   # Backend (Node.js & Express)
|   |-- controllers/          # API controllers
|   |-- models/               # Database models
|   |-- routes/               # API routes
|   |-- server.js             # Backend entry point
|-- UI/                       # Frontend (ReactJS)
    |-- src/                  # Frontend source code
        |-- components/       # React components
```

---

## **Technologies Used**

### **Languages and Frameworks**
- **Frontend:** ReactJS, JavaScript
- **Backend:** Node.js, Golang
- **Blockchain:** Cosmos SDK

### **Dependencies**
- **Docker:** For containerization and deployment.
- **Azure AI:** For invoice data extraction.
- **MongoDB:** For database storage.
- **Nginx:** For reverse proxy and load balancing.
- **Express.js:** For building backend APIs.

---

## **Setup Instructions**

Follow the steps below to run the project on your local machine:

### **1. Clone the Repository**
```bash
git clone <repository-url>
cd LedgerEase
```

### **2. Install Dependencies**

#### **Backend**
```bash
cd server
npm install
```

#### **Frontend**
```bash
cd UI
npm install
```

### **3. Configure Azure AI and Cosmos Blockchain**

- Set up Azure Document Intelligence AI ([Documentation](https://learn.microsoft.com/en-us/azure/synapse-analytics/machine-learning/tutorial-form-recognizer-use-mmlspark)) and add API keys in `ml-model/main.py`.
- Configure Cosmos Blockchain by navigating to `ignite_chain/invoice/config.yml` and updating the required settings.

### **4. Run the Project**

#### **Start Backend**
```bash
cd server
npm start
```

#### **Start Frontend**
```bash
cd UI/src
npm start
```

#### **Run Machine Learning Module**
```bash
cd ml-model
python3 main.py
```

#### **Start Cosmos Blockchain**
```bash
cd ignite_chain/invoice
ignite chain serve
```

### **5. Access the Application**
- Open your browser and navigate to `http://localhost:3000` for the frontend.
- Backend APIs are available at `http://localhost:5000`.

---

## **Usage**

1. **Admin Login:**
   - Log in using admin credentials to view and manage all invoices.

2. **Upload Invoices:**
   - Employees can log in and upload invoice files. The system extracts key data automatically.

3. **View Transaction History:**
   - Admins can view the history of all processed invoices on the dashboard.

4. **Data Security:**
   - All invoice data is securely stored on the Cosmos Blockchain, ensuring integrity and transparency.

---

## **Future Scope**

- **Automated Payment Reminders:**
  Send alerts for invoice due dates to ensure timely payments.
- **Profit and Loss Analysis:**
  Incorporate financial decision-making algorithms.
- **Integration with ERP Systems:**
  Enable seamless interoperability with existing business platforms.

---
