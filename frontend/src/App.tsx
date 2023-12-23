import React from "react";
import {
  BrowserRouter,
  Routes,
  Route
} from "react-router-dom";
import { 
  AppIndexPage, 
  LoginPage, 
  SignUpPage 
} from "./pages";
import "./App.css";

export default (): React.JSX.Element => {
  return (
    <React.StrictMode>
      <BrowserRouter>
        <Routes>
          <Route index={true} element={<AppIndexPage/>}/>
          <Route path="/login" element={<LoginPage/>}/>
          <Route path="/signup" element={<SignUpPage/>}/>
        </Routes>
      </BrowserRouter>
    </React.StrictMode>
  );
};
