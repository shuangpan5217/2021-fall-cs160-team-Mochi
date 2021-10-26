import Button from "../components/button";
import Template from "../components/template";
import UploadNotesWindow from "../components/uploadNotesWindow";
import { useState } from 'react';

function HomePage({authToken}) {

  const [buttonUpload, setButtonUpload] = useState(false);

  return (
    <>
      <Template showSearch={false} showProfile={true} body={
        <div className="d-flex flex-row">
          <Button title="UPLOAD" type="primary" clicked={()=> setButtonUpload(true)} />
          <UploadNotesWindow authToken={authToken} trigger={buttonUpload} setTrigger={setButtonUpload}/>
        </div>
      }/>
    </>

  );
}

export default HomePage;
