import { Routes, Route } from "react-router-dom"
import ExerciseList from './ExerciseList';
import { BrowserRouter } from "react-router-dom";

function AppRouter() {
    return (
        <BrowserRouter>
          <Routes>
            <Route path="/" element={ <ExerciseList/> } />
            {/* <Route path="about" element={ <Users/> } /> */}
          </Routes>
        </BrowserRouter>
    )
}
  

// export default AppRouter;


// import React from 'react';
// import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
// import ExerciseList from './ExerciseList';
// import SingleExercise from './SingleExercise';

// const App = () => {
//   return (
//     <Router>
//       <Switch>
//         <Route exact path="/exercises" component={ExerciseList} />
//         <Route path="/exercises/:id" component={SingleExercise} />
//       </Switch>
//     </Router>
//   );
// };

// export default App;

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