import s from "./App.module.scss";
import Create from "./pages/Create";
import { Routes, Route } from "react-router";
import Update from "./pages/Update";

function App() {
  return (
    <div className={s.studioApp}>
      <Routes>
        <Route path="/" element={<Create />} />
        <Route path=":id" element={<Update />} />
      </Routes>
    </div>
  );
}

export default App;
