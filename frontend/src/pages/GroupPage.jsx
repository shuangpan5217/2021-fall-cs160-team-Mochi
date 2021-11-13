import Template from "../components/template";
import ModalHeader from "../components/modalHeader.jsx";
import ModalLeftWindow from "../components/modalLeftWindow";
import Button from "../components/button";
import { useState, React, useEffect } from "react";
import AddMemberWindow from "../components/addMemberWindow";
import "../css/personalPage.css";
import UploadNotesWindow from "../components/uploadNotesWindow";
import { useParams } from "react-router";

function GroupPage(props) {
    const [buttonAddMember, setButtonAddMember] = useState(false);
    const [groupProfile, setButtonGroupProfile] = useState(false);
    const [members, setMembers] = useState([]);
    const [allGroups, setGroups] = useState([]);
    const [notes, setNotes] = useState([]);
    const [buttonUpload, setButtonUpload] = useState(false);
    const [group, setGroup] = useState("");
    const [groupDescription, setGroupDescription] = useState("");
    const { groupId } = useParams();

    const getGroupInfo = async () => {
        let success = true;
        const groupInfoResponse = await fetch(
            "http://localhost:3000/v1/group/" + groupId,
            {
                method: "GET",
                headers: {
                    Authorization:
                        "bearer " + window.localStorage.getItem("authToken"),
                },
            }
        ).catch((err) => {
            console.error(err);
            success = false;
        });

        if (success) {
            const groupInfoResponseJSON = await groupInfoResponse.json();
            if (groupInfoResponseJSON.group_name) {
                setGroup(groupInfoResponseJSON.group_name);
                setGroupDescription(groupInfoResponseJSON.description);
            } else {
                console.error("Could not get group information.");
            }
        } else {
            return;
        }
    };

    const getGroupMembers = async () => {
        let success = true;
        const membersResponse = await fetch(
            "http://localhost:3000/v1/group/" + groupId + "/members",
            {
                method: "GET",
                headers: {
                    Authorization:
                        "bearer " + window.localStorage.getItem("authToken"),
                },
            }
        ).catch((err) => {
            console.error(err);
            success = false;
        });

        if (success) {
            const memberResponseJSON = await membersResponse.json();
            setMembers([]);
            if (memberResponseJSON.members) {
                for (const member of memberResponseJSON.members) {
                    setMembers((arr) => [...arr, member]);
                }
            }
        } else {
            return;
        }
    };

    useEffect(() => {
        getGroupInfo();
        getGroupMembers();
    }, []);

    return (
        <>
            <Template
                showSearch={true}
                showProfile={true}
                body={
                    <div className="d-flex flex-column left-side">
                        <ModalHeader title={`Hi Group ${group}`} />
                        <div className="d-flex flex-column align-items-left">
                            <ModalLeftWindow
                                body={
                                    <div className="d-flex flex row">
                                        <ModalHeader title="Members" />
                                        <div className="d-flex row agenda">
                                            {members.map((member) => (
                                                <h3>{member.username}</h3>
                                            ))}
                                        </div>
                                        <div className="d-flex">
                                            <Button
                                                title="ADD MEMBER"
                                                type="primary"
                                                clicked={() =>
                                                    setButtonAddMember(true)
                                                }
                                            />
                                            <AddMemberWindow
                                                trigger={buttonAddMember}
                                                setTrigger={setButtonAddMember}
                                            />
                                        </div>
                                        <div className="d-flex">
                                            <Button
                                                title="EDIT PROFILE"
                                                type="secondary"
                                                clicked={() =>
                                                    setButtonGroupProfile(true)
                                                }
                                            />
                                        </div>
                                    </div>
                                }
                            />
                        </div>
                        <div className="d-flex flex-column right-side-top">
                            <ModalHeader title="Our Group" />
                            <div className="agenda-big">{groupDescription}</div>
                        </div>

                        <div className="d-flex right-side-middle">
                            <ModalHeader title="Our Notes" />
                            <div className="d-flex flex-row right-side-down">
                                <Button
                                    title="BACK"
                                    type="secondary"
                                    url="/home"
                                />
                                <Button
                                    title="UPLOAD"
                                    type="primary"
                                    clicked={() => setButtonUpload(true)}
                                />
                                <UploadNotesWindow
                                    trigger={buttonUpload}
                                    setTrigger={setButtonUpload}
                                />
                            </div>
                        </div>
                    </div>
                }
            />
        </>
    );
}
export default GroupPage;
