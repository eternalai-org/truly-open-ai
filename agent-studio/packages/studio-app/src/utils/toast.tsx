import toast from "react-hot-toast";
import ToastMessage from "../components/ToastMessage";

export const showValidateError = (msg: string) => {
  toast.error(msg, {
    icon: null,
    style: {
      borderColor: "blue",
      color: "blue",
    },
    duration: 3000,
    position: "bottom-center",
  });
};

export const showError = ({
  message,
}: {
  url?: string;
  linkText?: string;
  message: string;
  description?: string;
  theme?: "dark" | "light";
}) => {
  toast.error(() => <ToastMessage message={message} />, {
    duration: 5000,
    style: {
      borderLeft: "4px solid rgb(255, 255, 255, 0.4)",
      maxWidth: "300px",
      padding: `10px`,
      justifyContent: "space-between",
      background: " #ff4747",
      borderRadius: "8px",
      alignItems: "flex-start",
      gap: "8px",
      zIndex: "var(--index-above-all)",
    },
    position: "bottom-right",
    ariaProps: {
      role: "status",
      "aria-live": "polite",
    },
  });
};

export const showSuccess = ({
  message,
}: {
  url?: string;
  linkText?: string;
  description?: string;
  message: string;
  theme?: "dark" | "light";
}) => {
  toast.success(() => <ToastMessage message={message} />, {
    duration: 5000,
    style: {
      maxWidth: "300px",
      borderLeft: "4px solid rgb(255, 255, 255, 0.4)",
      padding: `10px`,
      justifyContent: "space-between",
      background: "#62d344",
      borderRadius: "8px",
      alignItems: "flex-start",
      gap: "8px",
      zIndex: "var(--index-above-all)",
    },
    ariaProps: {
      role: "status",
      "aria-live": "polite",
    },
    position: "bottom-right",
  });
};

export const removeToast = () => {
  toast.remove();
};
