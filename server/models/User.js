// models/User.js

const mongoose = require("mongoose");

// Define the schema for the user
const userSchema = new mongoose.Schema({
  fullName: {
    type: String,
    required: true,
  },
  email: {
    type: String,
    required: true,
    unique: true,
  },
  password: {
    type: String,
    required: true,
  },
  role: {
    type: String,
  },
  company: {
    type: String,
  },
  mobile: {
    type: String,
  },
  location: {
    type: String,
  },
});

// Create a model from the schema
const User = mongoose.model("User", userSchema);

module.exports = User;
