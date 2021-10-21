import Button from "./button";
import InputBox from "./inputBox";
import ModalHeader from "./modalHeader.jsx";
import { useState } from "react";
import { useHistory } from "react-router-dom";
import "../css/uploadNotes.css"

function UploadNotesWindow(props) {
    let history = useHistory();

    const [title, setTitle] = useState("");
    const [desc, setDescription] = useState("");
    const [tag, setTag] = useState("");

    const attemptUpload = async () => {
      const response = await fetch("http://localhost:3000/v1/notes", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          title,
          desc,
          tag,
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
    return (
      <div className="d-flex flex-column align-items-center">
        <ModalHeader title="Upload Notes" />
        <label className="form-upload">
          Title 
          <InputBox placeholder="title" onChange={setTitle} />
        </label>
        <label className="form-upload">
          Description 
          <InputBox placeholder="description" onChange={setDescription} />
        </label>
        <label className="form-upload">
          Tags 
          <InputBox placeholder="tag" onChange={setTag} />
        </label>
        <label>
          Sharing 
          Public
          Group
          Private
        </label>
        <label>
            Attach File 
        </label>
        <div className="d-flex flex-row ">
          <Button title="CANCEL" type="secondary" url="/home" />
          <Button title="UPLOAD" type="primary" clicked={attemptUpload} />
        </div>
      </div>
    );
  }
  
  export default UploadNotesWindow;
  