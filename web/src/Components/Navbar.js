import React, { useEffect, useState } from "react";
import '../css/Navbar.css';

function Navbar() {
    const [userName, setUserName] = useState("");
    const [userAvatarUrl, setUserAvatarUrl] = useState("");
    const [isLogged, setIsLogged] = useState(false)

    useEffect(() => {
        (async () => {
            try {

                const res = await fetch("http://localhost:8080/user-data", {
                    method: "GET",
                    credentials: "include"
                })

                if (res.status == 200) {
                    setIsLogged(true)
                } else if (res.status == 402) {
                    setIsLogged(false)
                } else {
                    throw new Error("Failed to fetch user data");
                }
                
                const result = await res.json()
                setUserName(result.name)
                setUserAvatarUrl(result.avatar_url)
            } catch (error) {
                console.error("Error fetching user data:", error);
            }    
        })();
    })

    if (isLogged) {
        return (
            <nav>
                <p>{userName}</p>
                <img src={userAvatarUrl} alt="Logo"/>
            </nav>
        );
    } else {
        <nav>
            <a href="http://localhost:8080/login">login</a>
        </nav>
    }

   
}

export default Navbar;