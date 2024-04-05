import React, { useEffect, useState } from "react";
import '../css/Navbar.css';

function Navbar() {
    const [userAvatarUrl, setUserAvatarUrl] = useState("");

    useEffect(() => {
        fetch("http://localhost:8080/avatar", {
            method: "GET",
            credentials: "include"

        })
        .then(res => res.json())
        .then(data => {
            setUserAvatarUrl(data.AvatarURL)
        })
    
    })

    return (
        <nav>
            <a href="http://localhost:8080/login">login</a>
        </nav>
    );
}

export default Navbar;