import React, { useEffect, useState } from "react";
import { useParams, Link } from "react-router-dom";

const SingleExercise = () => {
    const { id } = useParams();
    const [exercise, setExercise] = useState([]);
    const [genText, setGenText] = useState("");
    const [file, setFile] = useState(null);
    const [score, setScore] = useState(0)
    const [inputClicked, setInputClicked] = useState(false);

    const handleFileChange = async (event) => {
        setFile(event.target.files[0]);
        const formData = new FormData();
        formData.append("image", event.target.files[0]);
        try {
            const response = await fetch(`http://localhost:8080/extract-text`, {
                method: 'POST',
                body: formData
            });
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            const data = await response.json(); // assuming the response is plain text
            setGenText(data.text);
        } catch (error) {
            console.error('Error fetching data:', error);
        }
    };

    const handleChange = (event) => {
        setGenText(event.target.value);
    };

    const getResponse = async () => {
        setInputClicked(true);
        const requestData = {
            exercise_id: Number(id),
            user_id: 1,
            generate_text: genText
        };
        try {
            const response = await fetch(`http://localhost:8080/exercises/${id}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(requestData)
            });
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            const data = await response.json();
            setScore(data)
        } catch (error) {
            console.error('Error fetching data:', error);
        }
    }

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
            <textarea value={genText} onChange={handleChange} rows={10} cols={50} />
            <br />
            <input type="submit" onClick={getResponse}/>
            {(() => {
                if (inputClicked && score == 0) {
                    return <div>Loading...</div>
                }
                if (score != 0) {
                    return <p> This is your score: {score}</p>
                }
            })()}
        </div>
    )    
}

export default SingleExercise;