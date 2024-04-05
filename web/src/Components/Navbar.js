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
            <nav>
                <a href="http://localhost:8080/login">Login</a>
            </nav>
        );
    }

    return (
        <nav>
            <p>{userName}</p>
            <img src={userAvatarUrl} alt="Logo" />
            <br />
            <a href="http://localhost:8080/logout">Logout</a>
        </nav>
    );
}

export default Navbar;