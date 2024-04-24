import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";

function VowelsExerciseList() {
    const [jsonData, setJsonData] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [fetchedData, setFetchedData] = useState([]);
    
    useEffect(() => {
      fetchData();
    }, []);
  
    const fetchData = async () => {
      try {
        const response = await fetch("http://localhost:8080/exercises/vowels-exercises/");
        console.log(response)
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        const data = await response.json();
        setJsonData(data);
      } catch (error) {
        setError(error.message);
      } finally {
        setLoading(false);
      }
    };
    
    return (
        <div>
            <Link to={`/`}>back</Link>
            <h1>Data from Go Server</h1>
            <ul>
                {jsonData.map((item, index) => (
                <li key={item.id}>{item.name}</li>
                ))}
            </ul>
        </div>
    )
}

export default VowelsExerciseList