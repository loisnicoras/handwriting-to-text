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
  
  // const isUserLogged = () => {
  //   fetch('http://localhost:8080/exercises/1', {
  //     method: 'GET',
  //     headers: {
  //       'Content-Type': 'application/json',
  //       // Add any other headers as needed
  //     },
  //     credentials: 'include' // Include cookies in the request
  //   })
  //   .then(response => {
  //     if (response.status === 401) {
  //       // Unauthorized, redirect to login page
  //       window.location.href = 'http://localhost:8080/login';
  //     }
  //   })
  //   .catch(error => {
  //     // Handle any errors that occurred during the fetch
  //     console.error('Error:', error);
  //   });
  // }

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
          <li key={item.id}>
            <Link to={`/exercises/${item.id}`}>{item.name}</Link>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default ExerciseList;
