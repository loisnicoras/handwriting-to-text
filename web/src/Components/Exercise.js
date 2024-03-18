import React, { useState } from "react";

function Exercise(props){
    const [fetchedData, setFetchedData] = useState(null);

    const getExercice = async (buttonId) => {
        try {
          // Perform the fetch request with the buttonId concatenated to the URL
          const response = await fetch(`http://localhost:8080/exercises/${buttonId}`);
          if (!response.ok) {
            throw new Error('Network response was not ok');
          }
          const data = await response.json();
          setFetchedData(data);
        } catch (error) {
          console.error('Error fetching data:', error);
        }
      };

    return (
        <div>
            <button onClick={() => getExercice(props.data.id)}>{props.data.name}</button>
        </div>
    );
}


export default Exercise; 