import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import App from "./App";
import reportWebVitals from "./reportWebVitals";

import { ChakraProvider } from "@chakra-ui/react";
import chakraThemes from "./chakra-themes";
import { BrowserRouter } from "react-router";
import GlobalDataProvider from "./providers/GlobalDataProvider";
import ToastOverlay from "./components/ToastOverlay";

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);
root.render(
  <React.StrictMode>
    <GlobalDataProvider>
      <ChakraProvider theme={chakraThemes}>
        <BrowserRouter>
          <App />
        </BrowserRouter>
      </ChakraProvider>
    </GlobalDataProvider>
    <ToastOverlay />
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
