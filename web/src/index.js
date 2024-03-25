import React from 'react';
import ReactDOM from 'react-dom/client';
// import './index.css';
// import App from './Components/App';
import Navbar from './Components/Navbar';
// import reportWebVitals from './reportWebVitals';
// import { RouterProvider } from "react-router-dom";
import AppRouter from './Components/Router';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <Navbar />
    <AppRouter />
  </React.StrictMode>
);

