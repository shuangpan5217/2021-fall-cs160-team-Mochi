import { useState } from "react";
import CommentsTab from "./commentsTab";
import InfoTab from "./infoTab";
import SharingTab from "./sharingTab";
import TabButton from "./tabButton";

function NoteActions({ tags, members, comments }) {
    const [openTab, setOpenTab] = useState("Info");

    return (
        <div className="d-flex flex-column">
            <div className="d-flex flex-row">
                <TabButton
                    title="Info"
                    selected={openTab === "Info"}
                    clicked={() => setOpenTab("Info")}
                />
                <TabButton
                    title="Sharing"
                    selected={openTab === "Sharing"}
                    clicked={() => setOpenTab("Sharing")}
                />
                <TabButton
                    title="Comments"
                    selected={openTab === "Comments"}
                    clicked={() => setOpenTab("Comments")}
                />
            </div>
            {openTab === "Info" ? (
                <InfoTab tags={tags} />
            ) : openTab === "Sharing" ? (
                <SharingTab members={members} />
            ) : (
                <CommentsTab comments={comments}/>
            )}
        </div>
    );
}

export default NoteActions;
