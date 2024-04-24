import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";

function ExerciseTypes(){
    return (
        <div>
            <Link to="/exercises/audio-exercises/">Audio exercise</Link>
            <br />
            <Link to="/exercises/vowels-exercises/">Vowels exercise</Link>
        </div>
    )
}

export default ExerciseTypes;