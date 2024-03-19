import React, { useState } from "react";
import Exercise from "./Exercise";

function ExerciseName(props){
    const [fetchedData, setFetchedData] = useState([]);
    const [isClicked, setOnClick] = useState(false)

    const getExercice = async (buttonId) => {
      setOnClick(true) 
        try {
          // Perform the fetch request with the buttonId concatenated to the URL
          const response = await fetch(`http://localhost:8080/exercises/${buttonId}`);
          if (!response.ok) {
            throw new Error('Network response was not ok');
          }
          const data = await response.json();
          setFetchedData(data);
          console.log(fetchedData)
        } catch (error) {
          console.error('Error fetching data:', error);
        }
      };

    return (
        <div>
            <button onClick={() => getExercice(props.data.id)}>{props.data.name}</button>
            {isClicked && (
              <Exercise data={fetchedData}/>
            )}
        </div>
    );
}


export default ExerciseName; 