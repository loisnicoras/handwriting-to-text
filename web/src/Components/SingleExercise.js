import React, { useEffect, useState } from "react";
import { useParams, Link, Form } from "react-router-dom";

const SingleExercise = () => {
    const { id } = useParams();
    const [exercise, setExercise] = useState([]);
    const [genText, setGenText] = useState("");
    const [file, setFile] = useState(null);

    const handleFileChange = async (event) => {
        setFile(event.target.files[0]);
        const formData = new FormData();
        formData.append('file', event.target.files[0]);
        try {
            const response = await fetch(`http://localhost:8080/extract-text`, {
                method: 'POST',
                body: formData
            });
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            const data = await response.text(); // assuming the response is plain text
            console.log(response)
            setGenText(data);
        } catch (error) {
            console.error('Error fetching data:', error);
        }
    };

    const handleChange = (event) => {
        setGenText(event.target.value);
    };

    useEffect(() => {
        const getExercice = async () => {
            try {
                const response = await fetch(`http://localhost:8080/exercises/${id}`);
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                const data = await response.json();
                setExercise(data);
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
            <Link to={`/`}>back</Link>
            <p>{exercise.name}</p>
            <audio controls src={exercise.audio_path}>
                Your browser does not support the
                <code>audio</code> element.
            </audio>
            <input type="file" name="image" onChange={handleFileChange} accept="image/*" />
            <br />
            <textarea value={genText} onChange={handleChange} rows={4} cols={50} />
        </div>
    )    
}

export default SingleExercise;