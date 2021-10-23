import Button from "./button";
import InputBox from "./inputBox";
import ModalHeader from "./modalHeader.jsx";
import { useState } from "react";
import { useHistory } from "react-router-dom";
import "../css/uploadNotesWindow.css"
import RadioButton from "./radioButton";
import React, {useMemo} from 'react';
import {useDropzone} from 'react-dropzone';
import ModalWindow from "./modalWindow";

const dropzoneStyle = {
  flex: 1,
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'center',
  padding: '20px',
  borderWidth: 3,
  borderRadius: 2,
  fontSize: 18,
  borderColor: '#eeeeee',
  borderStyle: 'dashed',
  backgroundColor: '#fafafa',
  color: '#bdbdbd',
  outline: 'none',
  transition: 'border .24s ease-in-out'
};

const activeStyle = {
  borderColor: '#2196f3'
};

const acceptStyle = {
  borderColor: '#00e676'
};

const rejectStyle = {
  borderColor: '#ff1744'
};


function UploadNotesWindow(props) {
    let history = useHistory();

    const [title, setTitle] = useState("");
    const [desc, setDescription] = useState("");
    const [tag, setTag] = useState("");
    const [type, setType] = useState("");

    const setPublic = () => { 
      setType ("Public");
    };

    const setGroup = () => { 
      setType ("Group");
    };
    
    const setPrivate = () => { 
      setType ("Private");
    };

    const {
      acceptedFiles,
      getRootProps,
      getInputProps,
      isDragActive,
      isDragAccept,
      isDragReject
    } = useDropzone({accept: 'application/pdf'});
  
    const style = useMemo(() => ({
      ...dropzoneStyle,
      ...(isDragActive ? activeStyle : {}),
      ...(isDragAccept ? acceptStyle : {}),
      ...(isDragReject ? rejectStyle : {})
    }), [
      isDragActive,
      isDragReject,
      isDragAccept
    ]);

    const files = acceptedFiles.map(file => (
      <li key={file.path}>
        {file.path}
      </li>
    ));

    const attemptUpload = async () => {
      const response = await fetch("http://localhost:3000/v1/notes", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          title,
          desc,
          tag,
          type,
        }),
      });
      const responseJSON = await response.json();
      if (responseJSON.username) {
        alert("Upload successfully!");
        history.push("/home");
      } else {
        alert("Something went wrong");
      }
    };
    return (props.trigger) ?(
      <div className="popup">
      <ModalWindow body={
        <div className="d-flex flex-column align-items-center">
          <ModalHeader title="Upload Notes" />
          <div className="d-flex flex-column align-items-end">
            <InputBox label="Title" placeholder="title" onChange={setTitle} />
            <InputBox textArea label="Description" placeholder="description" onChange={setDescription} />
            <InputBox label="Tag" placeholder="tag" onChange={setTag} />
          </div>
          <p></p>
          <label className="agenda">
          Sharing 
          </label>
          <div className="d-flex flex-row ">
            <RadioButton group="sharing" label="Public" onChange={setPublic} />
            <RadioButton group="sharing" label="Group"  onChange={setGroup} />
            <RadioButton group="sharing" label="Private" onChange={setPrivate} />
          </div>
          <p></p>
          <label className="agenda">
              Attach File 
          </label>
          <div className="container">
            <div {...getRootProps({style})}>
              <input {...getInputProps()} />
              <p>Drag and drop here</p>
              <p>or</p>
              <p>click to select files (PDF only)</p>
            </div>
              <ul>{files}</ul>
          </div>
          <div className="d-flex flex-row ">
            <Button title="CANCEL" type="secondary" clicked={()=> props.setTrigger(false)} />
            <Button title="UPLOAD" type="primary" clicked={attemptUpload} />
          </div>
        </div>
      }/> 
      </div>
    ):"";
  }
  
  export default UploadNotesWindow;
  


