import Button from "../components/button";
import SearchBar from "../components/searchBar";
import Template from "../components/template";
import UploadNotesWindow from "../components/uploadNotesWindow";
import ModalHeader from "../components/modalHeader";
import { useState } from 'react';
import "../css/homepage.css"

function HomePage() {

  const [buttonUpload, setButtonUpload] = useState(false);

  return (
    <>
      <Template showSearch={false} showProfile={true} body={
        <div className="d-flex flex-column style">
          <ModalHeader title="Study Your Way" />
          <div className="d-flex flex-column align-items-center">
            <SearchBar />
            <div className="d-flex flex-row">
              <Button title="UPLOAD" type="primary" clicked={()=> setButtonUpload(true)} />
              <Button title="SERACH" type="secondary" url="/search" />
              <UploadNotesWindow trigger={buttonUpload} setTrigger={setButtonUpload}/>
            </div>
          </div>
        </div>
      }/>
    </>
  );
}

export default HomePage;
