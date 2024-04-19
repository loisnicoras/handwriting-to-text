import { BrowserRouter, Route, Routes } from "react-router-dom";
import ExerciseList from './ExerciseList';
import SingleExercise from "./SingleExercise";
import ExerciseTypes from "./ExerciseTypes";

function AppRouter() {
    return (
        <BrowserRouter>
          <Routes>
            <Route path="/" element={ <ExerciseTypes/> } />
            <Route path="/exercises/audio-exercises/" element={ <ExerciseList/> } />
            <Route path="/exercises/audio-exercises/:id" element={ <SingleExercise/> } />
          </Routes>
        </BrowserRouter>
    )
}
  
export default AppRouter;
