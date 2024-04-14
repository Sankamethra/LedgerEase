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

    fetchInvoices();
    getInvoiceImage();
  }, []);


  const handleViewDetails = async (invoice) => {

    setSelectedInvoice(invoice);
    setIsModalOpen(true);
  };

  const handleCloseModal = () => {
    setIsModalOpen(false); // Close the modal
  };


   // Filter out rows with empty or "N/A" fields
   const filteredInvoices = invoices.filter((invoice) => {
    const {
      Invoice_Number,
      Customer_Name,
      Invoice_Date,
      Total_Amount,
      Due_Date,
    } = invoice;
    return (
      Invoice_Number?.trim() !== "N/A" &&
      Customer_Name?.trim() !== "N/A" &&
      Invoice_Date?.trim() !== "N/A" &&
      Total_Amount?.trim() !== "N/A" &&
      Due_Date?.trim() !== "N/A"
    );
  });
  

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
            {filteredInvoices.map((invoice, index) => (
                <tr key={index}>
                <td>{invoice.Invoice_Number || "N/A"}</td>
                <td>{invoice.Customer_Name || "N/A"}</td>
                <td>{invoice.Invoice_Date || "N/A"}</td>
                <td>{invoice.Total_Amount || "N/A"}</td>
                <td>{invoice.Due_Date || "N/A"}</td>
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
;

export default ShowTransactions;
