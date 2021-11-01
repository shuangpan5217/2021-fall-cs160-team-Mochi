import { useState } from "react";
import CommentsTab from "./commentsTab";
import InfoTab from "./infoTab";
import SharingTab from "./sharingTab";
import TabButton from "./tabButton";

function NoteActions({
    title,
    descr,
    type,
    tags,
    members,
    comments,
    noteId,
    owner,
}) {
    const [openTab, setOpenTab] = useState("Info");

    return (
        <div className="d-flex flex-column note-actions-container">
            <div className="d-flex flex-row justify-content-around">
                <TabButton
                    title="Info"
                    selected={openTab === "Info"}
                    clicked={() => setOpenTab("Info")}
                />
                {type === "Private" ? (
                    <></>
                ) : (
                    <TabButton
                        title="Sharing"
                        selected={openTab === "Sharing"}
                        clicked={() => setOpenTab("Sharing")}
                    />
                )}
                <TabButton
                    title="Comments"
                    selected={openTab === "Comments"}
                    clicked={() => setOpenTab("Comments")}
                />
            </div>
            {openTab === "Info" ? (
                <InfoTab title={title} descr={descr} tags={tags} />
            ) : openTab === "Sharing" ? (
                <SharingTab members={members} />
            ) : (
                <CommentsTab
                    comments={comments}
                    noteId={noteId}
                    owner={owner}
                />
            )}
        </div>
    );
}

export default NoteActions;
