import React from 'react';
import {
    BrowserRouter,
    Switch,
    Route,
    Link,
    Routes
  } from "react-router-dom";
import App from './App';


function AppRouter() {
    <BrowserRouter>
        <Routes>
            <Route path="/" element={<App />}>
            </Route>
        </Routes>
    </BrowserRouter>   
}

export default AppRouter;