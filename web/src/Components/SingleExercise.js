import React, { useEffect, useState } from "react";
import { useParams, Link } from "react-router-dom";

const SingleExercise = () => {
    const { id } = useParams();
    const [exercise, setExercise] = useState([]);
    const [getText, setGetText] = useState({})
    const [file, setFile] = useState(null);

    const handleFileChange = (event) => {
        setFile(event.target.files[0]);
    };

    const handleChange = (event) => {
        setGetText(event.target.value);
    };

    const obtainText = async () =>{
        const formData = new FormData();
        formData.append('file', file);
        try {
            // Perform the fetch request with the buttonId concatenated to the URL
            const response = await fetch(`http://localhost:8080/extract-text`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: formData
            });
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            const data = await response.json();
            setGetText(data);
            } catch (error) {
            console.error('Error fetching data:', error);
        }
    }

    useEffect(() => {
        const getExercice = async () => {
            try {
                // Perform the fetch request with the buttonId concatenated to the URL
                const response = await fetch(`http://localhost:8080/exercises/${id}`);
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                const data = await response.json();
                setExercise(data);
                console.log(exercise)
                } catch (error) {
                console.error('Error fetching data:', error);
            }
            };
        getExercice();
        return () => setExercise(null);
    }, [id]);

    if (!exercise) {
        return <div>Loading...</div>
    }

    return (
        <div>
            {<Link to={`/`}>back</Link>}
            <p>{exercise.name}</p>
            <audio controls src={exercise.audio_path}>
                Your browser does not support the
                <code>audio</code> element.
            </audio>
            <form onSubmit={obtainText}>
                <br />
                <input type="file" onChange={handleFileChange} accept="image/*" capture="camera" />
                <br />
                <input type="file" onChange={handleFileChange} accept="image/*" />
                <br />
                <input type="submit" value="Upload" />
            </form>
            {(getText != {}) && (
                <textarea value={getText} onChange={handleChange} rows={4} cols={50} />
            )}
        </div>
    )    
}

export default SingleExercise;