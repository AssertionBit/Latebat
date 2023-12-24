import React from "react";
import {
  BrowserRouter,
  Routes,
  Route
} from "react-router-dom";
import { 
  AppBatchPage,
  AppIndexPage, 
  LoginPage, 
  SignUpPage 
} from "./pages";
import "./App.css";
import { InternalOutlet } from "./outlets";
import { Provider } from "react-redux";
import store from "./store";

export default (): React.JSX.Element => {
  return (
    <React.StrictMode>
      <Provider store={store}>
        <BrowserRouter>
          <Routes>
            <Route path="/login" element={<LoginPage/>}/>
            <Route path="/signup" element={<SignUpPage/>}/>
            <Route path="/" element={<InternalOutlet/>}>
              <Route index={true} element={<AppIndexPage/>}/>
              <Route path="/batch" element={<AppBatchPage/>}/>
            </Route>
          </Routes>
        </BrowserRouter>
      </Provider>
    </React.StrictMode>
  );
};
