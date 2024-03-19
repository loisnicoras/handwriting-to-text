import React from "react";

function Exercise(props) {

    return(
        <div>
            <p>{props.data.name}</p>
            <audio controls src={props.data.audio_path}>
                Your browser does not support the
                <code>audio</code> element.
            </audio>
            <br />
            <input type="file" id="takePhoto" accept="image/*" capture="camera" style={{ display: 'none' }} />
            <button onClick={() => document.getElementById('takePhoto').click()}>
                Take a picture
            </button>
            <br />
            <input type="file" id="uploadPhoto" accept="image/*" style={{ display: 'none' }} />
            <button onClick={() => document.getElementById('uploadPhoto').click()}>
                Upload a picture
            </button>
            <br />
        </div>
    )
}

export default Exercise;