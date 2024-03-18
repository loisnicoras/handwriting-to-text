import React from 'react';
import ReactDOM from 'react-dom/client';
// import './index.css';
import App from './Components/App';
import Navbar from './Components/Navbar';
// import reportWebVitals from './reportWebVitals';


const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <Navbar />
    <App />
  </React.StrictMode>
);

