import { Routes, Route } from "react-router-dom"
import App from './App';
import { BrowserRouter } from "react-router-dom";

function AppRouter() {
    return (
        <BrowserRouter>
          <Routes>
            <Route path="/" element={ <App/> } />
            {/* <Route path="about" element={ <Users/> } /> */}
          </Routes>
        </BrowserRouter>
    )
}
  

export default AppRouter;