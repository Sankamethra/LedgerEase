import React from "react";
import "./modal.css";

const Modal = ({ invoice, onClose }) => {
  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-content" onClick={(e) => e.stopPropagation()}>
        <button className="modal-close-btn" onClick={onClose}>
          X
        </button>
        {invoice && invoice.image && (
          <img
            src={`data:image/jpeg;base64,${invoice.image.data}`}
            alt="Invoice"
          />
        )}
      </div>
    </div>
  );
};

export default Modal;