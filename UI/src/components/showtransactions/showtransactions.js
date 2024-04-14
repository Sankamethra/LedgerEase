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

    setSelectedInvoice(invoice);
    setIsModalOpen(true);
  };

  const handleCloseModal = () => {
    setIsModalOpen(false); // Close the modal
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
              </tr>
            </thead>
            <tbody>
            {invoices.map((invoice, index) => (
                <tr key={index}>
                <td>{invoice.Invoice_Number}</td>
                <td>{invoice.Customer_Name}</td>
                <td>{invoice.Invoice_Date}</td>
                <td>{invoice.Total_Amount}</td>
                <td>{invoice.Due_Date}</td>
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
