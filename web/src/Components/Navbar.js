import React, { useEffect, useState } from "react";
import '../css/Navbar.css';

function Navbar() {
    const [userName, setUserName] = useState("");
    const [userAvatarUrl, setUserAvatarUrl] = useState("");
    const [isLogged, setIsLogged] = useState(false);

    useEffect(() => {
        (async () => {
            try {
                const res = await fetch("http://localhost:8080/user-data", {
                    method: "GET",
                    credentials: "include"
                });

                if (res.status === 200) {
                    setIsLogged(true);
                    const result = await res.json();
                    
                    setUserName(result.name);
                    setUserAvatarUrl(result.avatar_url);
                } else if (res.status === 401) {
                    setIsLogged(false);
                } else {
                    throw new Error("Failed to fetch user data");
                }
            } catch (error) {
                console.error("Error fetching user data:", error);
            }
        })();
    }, []);

    if (!isLogged) {
        return (
            <nav class="navbar">
                <div class="navbar-brand">
                    <span class="company-name">YourCompany</span>
                </div>
                <div class="navbar-links">
                    <a href="http://localhost:8080/login" class="login-button">Login</a>
                </div>
            </nav>
        );
    }

    return (
        <nav class="navbar">
            <div class="navbar-brand">
                <span class="company-name">YourCompany</span>
            </div>
            <div class="navbar-links">
                <div class="user-info">
                    <img src={userAvatarUrl} alt="User Icon" class="user-icon" />
                    <span class="username">{userName}</span>
                </div>
                <a href="http://localhost:8080/logout" class="logout-button">Logout</a>
            </div>
        </nav>
    );
}

export default Navbar;