import Button from "../components/button";
import Template from "../components/template";
import UploadNotesWindow from "../components/uploadNotesWindow";
import { useState } from 'react';

function HomePage(props) {

  const [buttonUpload, setButtonUpload] = useState(false);

  return (
    <>
      <Template body={
        <div className="d-flex flex-row">
          <Button title="UPLOAD" type="primary" clicked={()=> setButtonUpload(true)} />
          <UploadNotesWindow trigger={buttonUpload} setTrigger={setButtonUpload}>
          </UploadNotesWindow>
        </div>
      }/>
    </>

  );
}

export default HomePage;
