import "@/App.css";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Home from "@/pages/home";
import SignupPage from "@/pages/signup";
import LoginPage from "@/pages/login";

function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="signup" element={<SignupPage />} />
          <Route path="login" element={<LoginPage />} />
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
