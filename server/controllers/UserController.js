const User = require("../models/User");
const jwt = require("jsonwebtoken");
const bcrypt = require("bcrypt");
require("dotenv").config();

exports.signup = async (req, res) => {
  try {
    // Extract user data from request body
    const { fullName, email, password, role, company, mobile, location } =
      req.body;

    // Check if user already exists
    const existingUser = await User.findOne({ email });
    if (existingUser) {
      return res.status(400).json({ error: "User already exists" });
    }

    // Create a new user instance
    const newUser = new User({
      fullName,
      email,
      password,
      role,
      company,
      mobile,
      location,
    });

    // Save the user to the database
    await newUser.save();

    // Return success response
    res.status(201).json({ message: "Signup successful" });
  } catch (error) {
    console.error("Signup failed:", error);
    res.status(500).json({ error: "Signup failed" });
  }
};

exports.login = async (req, res) => {
  try {
    const { email, password } = req.body;

    // Check if user exists in the database
    const user = await User.findOne({ email });
    if (!user) {
      return res.status(404).json({ error: "User not found" });
    }

    // Custom password comparison function
    const isMatch = comparePasswords(password, user.password);

    if (!isMatch) {
      console.log("Incorrect password:", password);
      return res.status(401).json({ error: "Incorrect password" });
    }

    // Generate JWT token
    if (!process.env.JWT_SECRET) {
      throw new Error("JWT secret key not found");
    }
    const token = jwt.sign({ id: user._id }, process.env.JWT_SECRET, {
      expiresIn: "1h",
    });

    // Send token as response
    res.status(200).json({ token });
  } catch (error) {
    console.error("Login failed:", error);
    res.status(500).json({ error: "Login failed" });
  }
};

// Custom password comparison function
function comparePasswords(inputPassword, hashedPassword) {
  // Compare the input password with the hashed password using your own logic
  if (inputPassword.length !== hashedPassword.length) {
    return false;
  }
  for (let i = 0; i < inputPassword.length; i++) {
    if (inputPassword[i] !== hashedPassword[i]) {
      return false;
    }
  }
  return true;
}
