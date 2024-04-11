// src/components/Signup.js
import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import "./signup.css";
import axios from "axios";

const Signup = () => {
  const navigate = useNavigate();
  const [formData, setFormData] = useState({
    fullName: "",
    email: "",
    password: "",
    role: "",
    company: "",
    mobile: "",
    confirmPassword: "",
    location: "",
  });

  const handleInputChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post("/api/signup", formData);
      console.log("Signup successful:", response.data);
      console.log("New data entry created:", formData);
      // Redirect to login page after successful signup
      navigate("/login");
    } catch (error) {
      console.error("Signup failed:", error);
      // Handle signup error
    }
  };

  const handleLoginClick = () => {
    // Redirect to the login page
    navigate("/login");
  };

  return (
    <div className="signup-container">
      <div className="signup-card">
        <h2 className="signup-heading">Signup</h2>

        {/* Details Container */}
        <div className="details-container">
          {/* Row 1 */}
          <div className="form-group">
            <label htmlFor="fullName">Full Name</label>
            <input
              type="text"
              id="fullName"
              name="fullName"
              value={formData.fullName}
              onChange={handleInputChange}
            />
          </div>

          <div className="form-group">
            <label htmlFor="email">Email</label>
            <input
              type="text"
              id="email"
              name="email"
              value={formData.email}
              onChange={handleInputChange}
            />
          </div>

          <div className="form-group">
            <label htmlFor="password">Set Password</label>
            <input
              type="password"
              id="password"
              name="password"
              value={formData.password}
              onChange={handleInputChange}
            />
          </div>

          <div className="form-group">
            <label htmlFor="role">Role in the Company</label>
            <input
              type="text"
              id="role"
              name="role"
              value={formData.role}
              onChange={handleInputChange}
            />
          </div>

          {/* Row 2 */}
          <div className="form-group">
            <label htmlFor="company">Company</label>
            <input
              type="text"
              id="company"
              name="company"
              value={formData.company}
              onChange={handleInputChange}
            />
          </div>

          <div className="form-group">
            <label htmlFor="mobile">Mobile</label>
            <input
              type="text"
              id="mobile"
              name="mobile"
              value={formData.mobile}
              onChange={handleInputChange}
            />
          </div>

          <div className="form-group">
            <label htmlFor="confirmPassword">Confirm Password</label>
            <input
              type="password"
              id="confirmPassword"
              name="confirmPassword"
              value={formData.confirmPassword}
              onChange={handleInputChange}
            />
          </div>

          <div className="form-group">
            <label htmlFor="location">Location</label>
            <input
              type="text"
              id="location"
              name="location"
              value={formData.location}
              onChange={handleInputChange}
            />
          </div>
        </div>

        {/* Submit Button */}
        <button type="submit" onClick={handleSubmit}>
          Submit
        </button>
        <div className="login-container">
          <h4 className="signup-heading" onClick={handleLoginClick}>
            Have an Account? Login
          </h4>
        </div>
      </div>
    </div>
  );
};

export default Signup;
