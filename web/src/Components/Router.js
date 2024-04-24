import { BrowserRouter, Route, Routes } from "react-router-dom";
import ExerciseList from './ExerciseList';
import SingleExercise from "./SingleExercise";
import ExerciseTypes from "./ExerciseTypes";
import VowelsExerciseList from "./VowelExerciseList";

function AppRouter() {
    return (
        <BrowserRouter>
          <Routes>
            <Route path="/" element={ <ExerciseTypes/> } />
            <Route path="/exercises/audio-exercises/" element={ <ExerciseList/> } />
            <Route path="/exercises/audio-exercises/:id" element={ <SingleExercise/> } />
            <Route path="/exercises/vowels-exercises/" element={<VowelsExerciseList/>} />
          </Routes>
        </BrowserRouter>
    )
}
  
export default AppRouter;
