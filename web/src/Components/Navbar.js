import React from "react";
import '../css/Navbar.css';

function Navbar() {
    const handleLogin = async () => {
        try {
            const response = await fetch('http://localhost:8080/login');
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            const data = await response.json();
            window.location.href = data.redirect_url;
        } catch (error) {
            console.error('Error logging in:', error);
        }
    };

    return (
        <nav>
            <button onClick={handleLogin}>Login with Google</button>
        </nav>
    );
}

export default Navbar;