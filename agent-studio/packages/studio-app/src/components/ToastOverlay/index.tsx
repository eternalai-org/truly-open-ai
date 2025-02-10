import React from "react";
import { Toaster } from "react-hot-toast";

const ToastOverlay = () => {
  return (
    <Toaster
      containerStyle={{
        top: 80,
        left: 20,
        bottom: 20,
        right: 20,
        zIndex: 9999999999999,
      }}
      toastOptions={{
        success: {
          className: "toast-success",
          style: {
            padding: "4px 12px",
            border: "1px solid #00AA6C",
            color: "#00AA6C",
            background: "#FFFFFF",
            width: "fit-content",
            maxWidth: "100%",
            whiteSpace: "pre-wrap",
            wordBreak: "break-word",
            gap: "8px",
          },
        },
        error: {
          className: "toast-error",
          style: {
            border: "1px solid #FF4747",
            color: "#FF4747",
            background: "#FFFFFF",
            justifyContent: "center",
          },
        },
        style: {
          padding: "10px 32px",
          boxShadow: "8px 7px 24px rgba(0, 0, 0, 0.15)",
          borderRadius: "4px",
          width: "100%",
          minWidth: "100px",
          justifyContent: "center",
        },
        duration: 2000,
      }}
    />
  );
};

export default ToastOverlay;
