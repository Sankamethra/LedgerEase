const mongoose = require("mongoose");

// Define the schema for the invoice
const invoiceSchema = new mongoose.Schema({
  "Invoice Number": {
    type: String,
  },
  "Customer Name": {
    type: String,
  },
  "Invoice Date": {
    type: Date,
  },
  "Issue Date": {
    type: Date,
  },
  "Total Amount": {
    type: String,
  },
  "Due Date": {
    type: Date,
  },
  image: {
    data: Buffer, // Store image data as a Buffer
    contentType: String, // Store image content type
  },
});

// Create a model from the schema
const Invoice = mongoose.model("Invoice", invoiceSchema);

module.exports = Invoice;
