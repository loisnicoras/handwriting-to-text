import { BrowserRouter, Route, Routes } from "react-router-dom";
import ExerciseList from './ExerciseList';
import SingleExercise from "./SingleExercise";

function AppRouter() {
    return (
        <BrowserRouter>
          <Routes>
            <Route path="/exercises/audio-exercises/" element={ <ExerciseList/> } />
            <Route path="/exercises/audio-exercises/:id" element={ <SingleExercise/> } />
          </Routes>
        </BrowserRouter>
    )
}
  
export default AppRouter;



// import React from 'react';
// import { useParams } from 'react-router-dom';

// const SingleExercise = () => {
//   // Extracting the id parameter from the route
//   const { id } = useParams();

//   return (
//     <div>
//       <h2>Exercise Details</h2>
//       <p>Exercise ID: {id}</p>
//       {/* Other details of the exercise can be displayed here */}
//     </div>
//   );
// };

// export default SingleExercise;