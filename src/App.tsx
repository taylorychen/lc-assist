import { Routes, Route } from "react-router-dom";
import Dashboard from "./dashboard";
import Question from "./question";

function App() {
  return (
    <>
      <Routes>
        <Route path="/" element={<Dashboard />} />
        <Route path="/question" element={<Question />} />
      </Routes>
    </>
  );
}

export default App;
