import React, { useState, useEffect } from "react";
import Header from "../header/header";
import "./showtransactions.css";

const ShowTransactions = () => {
  const [invoices, setInvoices] = useState([]);
  const [selectedInvoice, setSelectedInvoice] = useState(null); // State to store selected invoice

  useEffect(() => {
    const fetchInvoices = async () => {
      try {
        const response = await fetch("http://localhost:5002/invoices");
        if (response.ok) {
          const data = await response.json();
          setInvoices(data);
        } else {
          console.error("Failed to fetch invoices");
        }
      } catch (error) {
        console.error("Error fetching invoices:" + error.message);
      }
    };

    fetchInvoices(); // Call fetchInvoices directly inside useEffect
  }, []); // Add an empty dependency array to run the effect only once

  const handleViewDetails = (invoice) => {
    // Set the selected invoice when View Details button is clicked
    setSelectedInvoice(invoice);
  };

  return (
    <div>
      <Header />
      <div className="transaction-details-container">
        <h2 className="transaction-details-heading">Transaction Details</h2>
        <table className="transaction-table">
          <thead>
            <tr>
              <th>Invoice Number</th>
              <th>Customer Name</th>
              <th>Invoice Date</th>
              <th>Total Amount</th>
              <th>Due Date</th>
              <th>View Details</th>
            </tr>
          </thead>
          <tbody>
            {invoices.map((invoice, index) => (
              <tr key={index}>
                <td>{invoice["Invoice Number"]}</td>
                <td>{invoice["Customer Name"]}</td>
                <td>{invoice["Invoice Date"]}</td>
                <td>{invoice["Total Amount"]}</td>
                <td>{invoice["Due Date"]}</td>
                <td>
                  <button onClick={() => handleViewDetails(invoice)}>
                    View Details
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
        {/* Display selected invoice details, including image */}
        {selectedInvoice && (
          <div className="selected-invoice-details">
            {selectedInvoice.image && (
              <img
                src={`data:image/jpeg;base64,${selectedInvoice.image.data}`}
                alt="Invoice"
              />
            )}
          </div>
        )}
      </div>
    </div>
  );
};

export default ShowTransactions;
