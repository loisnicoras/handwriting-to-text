import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";

function ExerciseList() {
  const [jsonData, setJsonData] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [fetchedData, setFetchedData] = useState([]);
  
  useEffect(() => {
    fetchData();
  }, []);

  const fetchData = async () => {
    try {
      const response = await fetch("http://localhost:8080/exercises/audio-exercises/");
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
  
  if (loading) {
    return (
      <div>
        <h1>Loading...</h1>
      </div>
    );
  }
  
  if (error) {
    return (
    <div>
      <h1>Error: {error}</h1>
    </div>
    );
  }

  if (jsonData === null) {
    return (
    <div>
      <h1>No data available</h1>
    </div>
    );
  }

  return (
    <div>
      <Link to={`/`}>back</Link>
      <h1>Data from Go Server</h1>
      <ul>
        {jsonData.map((item, index) => (
          <li key={item.id}>
            <Link 
              to={`/exercises/audio-exercises/${item.id}`}> 
              {item.name}
            </Link>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default ExerciseList;
