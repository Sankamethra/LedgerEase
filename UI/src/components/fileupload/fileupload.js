import React, { useState } from "react";
import Header from "../header/header";
import { useNavigate } from "react-router-dom"; // Import useNavigate from react-router-dom
import "./fileupload.css";

const FileUpload = () => {
  const [selectedFile, setSelectedFile] = useState(null);
  const navigate = useNavigate(); // Initialize useNavigate

  const handleFileChange = (e) => {
    setSelectedFile(e.target.files[0]);
  };

  const handleUpload = () => {
    if (selectedFile) {
      const formData = new FormData();
      formData.append("file", selectedFile);
      fetch("http://localhost:5002/upload", {
        method: "POST", // Specify the method as POST
        body: formData,
      })
        .then((response) => {
          if (response.ok) {
            console.log("File uploaded successfully");
            // Navigate to showtransactions page upon successful upload
            navigate("/showtransactions");
          } else {
            console.error("Failed to upload file");
          }
        })
        .catch((error) => {
          console.error("Error uploading file:", error);
        });
    } else {
      console.log("No file selected");
    }
  };

  return (
    <div>
      <Header />
      <div className="file-upload-container">
        <div className="file-upload-content">
          <h2>Upload Invoice</h2>
          <div className="file-input-group">
            <input
              type="file"
              id="file"
              name="file"
              className="file-input"
              onChange={handleFileChange}
            />
            <label htmlFor="file" className="choose-file-button">
              Choose File
            </label>
          </div>
          {selectedFile && <p>Selected File: {selectedFile.name}</p>}
          <button className="upload-button" onClick={handleUpload}>
            Upload
          </button>
        </div>
      </div>
    </div>
  );
};

export default FileUpload;
