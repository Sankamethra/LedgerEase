// Header.js
import React from "react";
import { useNavigate } from "react-router-dom";
import logo from "../assets/logo.png"; // Import your logo image
import "./header.css"; // Import header styles

const Header = () => {
  const navigate = useNavigate();

  const handleFileUploadClick = () => {
    // Redirect to the add user page
    navigate("/fileupload");
  };

  const handleShowTransactionsClick = () => {
    // Redirect to the add user page
    navigate("/showtransactions");
  };

  return (
    <header className="header">
      <div className="logo-container">
        <img src={logo} alt="Logo" className="logo" />
      </div>
      <div className="heading-container">
        <h1 className="heading">Ledger Ease</h1>
      </div>
      <nav classname="menu-bar">
        <ul className="menu">
          <li onClick={handleFileUploadClick}>Upload Invoice</li>
          <li onClick={handleShowTransactionsClick}>Show Transactions</li>
        </ul>
      </nav>
      <div className="profile-icon">
        {/* Add your profile icon component or image here */}
      </div>
    </header>
  );
};

export default Header;
