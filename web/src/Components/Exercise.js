import React from "react";

function Exercise(props){
    return (
        <div>
            <button><span className="label">{props.data.name}</span></button>
        </div>
    );
}


export default Exercise; 