import React from "react";
import "./modal.css";

const Modal = ({ invoice, onClose }) => {
  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-content" onClick={(e) => e.stopPropagation()}>
        <button className="modal-close-btn" onClick={onClose}>
          X
        </button>
        {invoice && invoice.imageSrc && (
          <img src={invoice.imageSrc} alt="Invoice" />
        )}
      </div>
    </div>
  );
};

export default Modal;
