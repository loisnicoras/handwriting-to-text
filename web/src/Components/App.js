import React, { useState, useEffect } from "react";
import ExerciseName from './ExerciseName'
// import logo from './logo.svg';
// import './App.css';

function App() {
  const [jsonData, setJsonData] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  
  useEffect(() => {
    fetchData();
  }, []);

  const fetchData = async () => {
    try {
      const response = await fetch("http://localhost:8080/exercises/");
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
      <h1>Data from Go Server</h1>
      <ul>
        {jsonData.map((item, index) => (
          <ExerciseName key={index} data={item} />
        ))}
      </ul>
    </div>
  );
}

export default App;
