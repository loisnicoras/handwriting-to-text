import React from 'react';
import {
    BrowserRouter,
    Switch,
    Route,
    Link,
    Routes
  } from "react-router-dom";
import App from './App';
import ExerciseName from './ExerciseName';


function AppRouter() {
    <BrowserRouter>
        <Routes>
            <Route path="/" element={<App />}>
            </Route>
            <Route path="/exercise" element={<ExerciseName/>}>
            </Route>
        </Routes>
    </BrowserRouter>   
}

export default AppRouter;