import UserListItem from "./userListItem";
import InputBox from "./inputBox";
import dummyProfile from "../media/mochi.jpeg";
import NoteActionHeader from "./noteActionHeader";
import { useState } from "react";
import NoteActionButton from "./noteActionButton";

function SharingTab({ members }) {
    const [addMember, setAddMember] = useState("");
    const memberElems = members.map((member) => (
        <UserListItem img={dummyProfile} name={member} />
    ));

    return (
        <div className="d-flex flex-column">
            <NoteActionHeader title="Shared with" />
            {memberElems}
            <div className="d-flex flex-column align-items-center">
                <InputBox placeholder="member" onChange={setAddMember} />
                <NoteActionButton
                    title="Share"
                    onClick={() => console.log("share")}
                />
            </div>
        </div>
    );
}

export default SharingTab;
