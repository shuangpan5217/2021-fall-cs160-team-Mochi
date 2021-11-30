import Button from "../components/button";
import SearchBar from "../components/searchBar";
import Template from "../components/template";
import UploadNotesWindow from "../components/uploadNotesWindow";
import ModalHeader from "../components/modalHeader";
import { useState } from "react";
import "../css/homepage.css";

function HomePage() {
    const [buttonUpload, setButtonUpload] = useState(false);
    const [updateSearch, setUpdateSearch] = useState(false);

    return (
        <>
            <Template
                showSearch={false}
                showProfile={true}
                blur
                body={
                    <div className="d-flex flex-column center-home-content">
                        <ModalHeader title="Study Your Way" />
                        <div className="d-flex flex-column align-items-center">
                            <SearchBar
                                showFilterBtn={false}
                                updateSearch={updateSearch}
                            />
                            <div className="d-flex flex-row">
                                <Button
                                    title="UPLOAD"
                                    type="secondary"
                                    clicked={() => setButtonUpload(true)}
                                />
                                <Button
                                    title="SEARCH"
                                    type="primary"
                                    clicked={() => setUpdateSearch(true)}
                                />
                            </div>
                        </div>
                    </div>
                }
            />
            <div>
                <UploadNotesWindow
                    trigger={buttonUpload}
                    setTrigger={setButtonUpload}
                />
            </div>
        </>
    );
}

export default HomePage;
