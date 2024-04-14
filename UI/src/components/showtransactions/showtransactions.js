import React, { useState, useEffect } from "react";
import Header from "../header/header";
import Modal from "../modal/modal";
import "./showtransactions.css";

const ShowTransactions = () => {
  const [invoices, setInvoices] = useState([]);
  const [error, setError] = useState(null);
  const [selectedInvoice, setSelectedInvoice] = useState(null); // State to store selected invoice
  const [isModalOpen, setIsModalOpen] = useState(false); // State to control modal visibility

  useEffect(() => {
    const fetchInvoices = async () => {
      try {
        const response = await fetch("http://localhost:5002/query");
        if (response.ok) {
          const data = await response.json();
          console.log("Fetched data:", data);
          setInvoices(data.invoice);
        } else {
          const errorData = await response.json();
          setError(errorData.error || "Failed to fetch invoices");
        }
      } catch (error) {
        console.error("Fetch Error:", error);
        setError("Error fetching invoices: " + error.message);
      }
    };

    fetchInvoices();
  }, []);

  const handleViewDetails = async (invoice) => {
    console.log("Clicked View Details for Invoice:", invoice);
    if (!invoice || !invoice["_id"]) { // Check for "_id" existence
      console.error("Invalid invoice:", invoice);
      return;
    }
    
    const invoiceId = invoice["_id"]; // Get the invoice ID
    console.log(invoiceId)
    try {
      // Fetch the image data using the correct endpoint with invoiceId
      const response = await fetch(`http://localhost:5002/invoices/${invoiceId}`);
      // Rest of the function...
    } catch (error) {
      console.error("Error handling view details:", error);
    }
  };

  const handleCloseModal = () => {
    setIsModalOpen(false); // Close the modal
  };

  const getInvoiceImage = async (invoiceId) => {
    try {
      const response = await fetch(`http://localhost:5002/invoices`);
      if (response.ok) {
        const imageData = await response.blob(); // Get image data as Blob
        return URL.createObjectURL(imageData); // Create object URL for the image blob
      } else {
        console.error("Failed to fetch invoice image:", response.status);
        return null;
      }
    } catch (error) {
      console.error("Error fetching invoice image:", error);
      return null;
    }
  };

  return (
    <div>
      <Header />
      <div className="transaction-details-container">
        <h2 className="transaction-details-heading">Transaction Details</h2>
        {error ? (
          <div className="error-message">{error}</div>
        ) : (
          <table className="transaction-table">
            <thead>
              <tr>
                <th>Invoice Number</th>
                <th>Customer Name</th>
                <th>Invoice Date</th>
                <th>Total Amount</th>
                <th>Due Date</th>
                <th>Action</th>
              </tr>
            </thead>
            <tbody>
              {invoices.map((invoice, index) => (
                <tr key={index}>
                  <td>{invoice["Invoice_Number"]}</td>
                  <td>{invoice["Customer_Name"]}</td>
                  <td>{invoice["Invoice_Date"]}</td>
                  <td>{invoice["Total_Amount"]}</td>
                  <td>{invoice["Due_Date"]}</td>
                  <td>
                    <button onClick={() => handleViewDetails(invoice)}>
                      View Details
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
      </div>
      {/* Render the modal only when isModalOpen is true */}
      {isModalOpen && (
        <Modal invoice={selectedInvoice} onClose={handleCloseModal} />
      )}
    </div>
  );
};

export default ShowTransactions;
