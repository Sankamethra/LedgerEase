// src/components/Login.js
import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import "../signup/signup.css";
import axios from "axios";

const Login = () => {
  const navigate = useNavigate();
  const [formData, setFormData] = useState({
    email: "",
    password: "",
  });
  const [error, setError] = useState("");

  const handleInputChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSignupClick = () => {
    navigate("/signup");
  };

  const handleLogin = async () => {
    try {
      const response = await axios.post("/api/login", formData);
      localStorage.setItem("token", response.data.token);
      navigate("/fileupload");
    } catch (error) {
      setError("Invalid email or password");
    }
  };

  return (
    <div className="signup-container">
      <div className="signup-card">
        <h2 className="signup-heading">Login</h2>
        <div className="error">{error}</div>
        <div className="details-container">
          <div className="form-group">
            <label htmlFor="email">Email</label>
            <input
              type="email"
              id="email"
              name="email"
              value={formData.email}
              onChange={handleInputChange}
            />
          </div>
          <div className="form-group">
            <label htmlFor="password">Password</label>
            <input
              type="password"
              id="password"
              name="password"
              value={formData.password}
              onChange={handleInputChange}
            />
          </div>
        </div>
        <button type="submit" onClick={handleLogin}>
          Login
        </button>
        <div className="login-container">
          <h4 className="signup-heading" onClick={handleSignupClick}>
            Don't Have an Account? Signup
          </h4>
        </div>
      </div>
    </div>
  );
};

export default Login;
