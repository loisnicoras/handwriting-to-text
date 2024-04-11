import React, { useEffect, useState } from "react";
import { useParams, Link } from "react-router-dom";

const SingleExercise = () => {
    const { id } = useParams();
    const [exercise, setExercise] = useState([]);
    const [genText, setGenText] = useState("");
    const [score, setScore] = useState(0)
    const [inputClicked, setInputClicked] = useState(false);
    const [error, setError] = useState(null);

    const handleFileChange = async (event) => {
        const formData = new FormData();
        formData.append("image", event.target.files[0]);
        try {
            const response = await fetch(`http://localhost:8080/extract-text`, {
                method: 'POST',
                body: formData
            });
            if (!response.ok) {
                throw new Error('Failed to extract text');
            }
            const data = await response.json(); // assuming the response is plain text
            setGenText(data.text);
        } catch (error) {
            setError(error.message);
        }
    };

    const handleChange = (event) => {
        setGenText(event.target.value);
    };

    const submitExercise = async () => {
        setInputClicked(true);
        const requestData = {
            gen_text: genText
        };
        try {
            const response = await fetch(`http://localhost:8080/exercises/${id}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                credentials: "include", 
                body: JSON.stringify(requestData)
            });
            if (!response.ok) {
                throw new Error('Failed to submit exercise');
            }
            const data = await response.json();
            setScore(data)
        } catch (error) {
            setError(error.message);
        }
    }

    useEffect(() => {
        fetchExercise();
        return () => setExercise(null);
    }, [id]);

    const fetchExercise = async () => {
        try {
            const response = await fetch(`http://localhost:8080/exercises/${id}`, {
                method: "GET",
                credentials: "include", 
            });

            if (!response.ok) {
                if (response.status === 401) {
                    window.location.href = 'http://localhost:8080/login'
                } else {
                    throw new Error('Failed to fetch exercise');
                }
            }
            const data = await response.json();
            setExercise(data);
        } catch (error) {
            setError(error.message);
        }
    }
    
    // if (error) {
    //     return (
    //         <div>
    //             <Link to={`/`}>back</Link>
    //             <div>Error: {error}</div>
    //         </div>
    //     );
    // }

    if (!exercise) {
        return <div>Loading exercise...</div>
    }

    return (
        <div>
            <Link to={`/`}>back</Link>
            <p>{exercise.name}</p>
            <audio controls src={exercise.audio_path}>
                Your browser does not support the
                <code>audio</code> element.
            </audio>
            <input type="file" name="image" onChange={handleFileChange} accept="image/*" />
            <br />
            <textarea value={genText} onChange={handleChange} rows={10} cols={50} />
            <br />
            <input type="submit" onClick={submitExercise}/>
            {(() => {
                if (inputClicked && score === 0) {
                    return <div>Loading score...</div>
                }
                if (score !== 0) {
                    return <p>This is your score: {score}</p>
                }
            })()}
        </div>
    )    
}

export default SingleExercise;