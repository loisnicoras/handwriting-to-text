import React, { useState, useEffect } from "react";
import Exercise from './Exercise'
// import logo from './logo.svg';
// import './App.css';

function App() {
  const [jsonData, setJsonData] = useState([]);

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
      console.error('Error fetching data:', error);
    }
  };


  return (
    <div>
      <h1>Data from Go Server</h1>
      <ul>
        {jsonData.map((item) => (
          <Exercise data={item} />
        ))}
      </ul>
    </div>
  );
}

export default App;
