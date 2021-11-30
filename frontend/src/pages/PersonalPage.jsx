import Template from "../components/template";
import ModalHeader from "../components/modalHeader.jsx";
import SectionTitle from "../components/sectionTitle.jsx";
import LeftPanel from "../components/leftPanel";
import Button from "../components/button";
import { useState, React, useEffect } from "react";
import AddFriendWindow from "../components/addFriendWindow";
import CreateGroupWindow from "../components/createGroupWindow";
import "../css/personalPage.css";
import UploadNotesWindow from "../components/uploadNotesWindow";
import { Link } from "react-router-dom";
import PDFViewer from "../components/pdfViewer";
import PersonalPrefWindow from "../components/personalPrefWindow";
import ModalWindow from "../components/modalWindow";

function PersonalPage(props) {
    const [buttonAddFriend, setButtonAddFriend] = useState(false);
    const [buttonGroup, setButtonGroup] = useState(false);
    const [buttonPersonalProfile, setButtonPersonalProfile] = useState(false);
    const [friends, setFriends] = useState([]);
    const [groups, setGroups] = useState([]);
    const [notes, setNotes] = useState([]);
    const [buttonUpload, setButtonUpload] = useState(false);
    const [user, setUser] = useState("");
    const [userDescription, setUserDescription] = useState("");

    const [pdfs, setPDFs] = useState([]);
    const [refreshProfileImage, setRefreshProfileImage] = useState(false);

    const getUserInfo = async () => {
        let success = true;
        const userInfoResponse = await fetch("http://localhost:3000/v1/user/", {
            method: "GET",
            headers: {
                Authorization:
                    "bearer " + window.localStorage.getItem("authToken"),
            },
        }).catch((err) => {
            console.error(err);
            success = false;
        });

        if (success) {
            const userInfoResponseJSON = await userInfoResponse.json();
            if (userInfoResponseJSON.username) {
                setUser(userInfoResponseJSON.username);
                setUserDescription(userInfoResponseJSON.description);
            } else {
                console.error("Could not get user information.");
            }
        }
    };

    const getMyFriends = async () => {
        let success = true;
        const friendsResponse = await fetch(
            "http://localhost:3000/v1/friends/",
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
            const friendResponseJSON = await friendsResponse.json();
            if (friendResponseJSON.friends) {
                setFriends(friendResponseJSON.friends);
            }
        }
    };

    const getMyGroups = async () => {
        let success = true;
        const groupsResponse = await fetch("http://localhost:3000/v1/groups/", {
            method: "GET",
            headers: {
                Authorization:
                    "bearer " + window.localStorage.getItem("authToken"),
            },
        }).catch((err) => {
            console.error(err);
            success = false;
        });

        if (success) {
            const groupResponseJSON = await groupsResponse.json();
            if (groupResponseJSON.allGroups) {
                setGroups(groupResponseJSON.allGroups);
            }
        }
    };

    const getUserNotesRef = async () => {
        let success = true;
        const userNotesResponse = await fetch(
            "http://localhost:3000/v1/notes/username/",
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
            const userNoteResponseJSON = await userNotesResponse.json();
            if (userNoteResponseJSON.notes) {
                setNotes(userNoteResponseJSON.notes);
                await getPDF(userNoteResponseJSON.notes);
            } else {
                console.error("Could not load note.");
            }
        }
    };

    const getPDF = async (notes) => {
        let success = true;
        for (const note of notes) {
            const pdfResponse = await fetch(
                "http://localhost:3000/v1/notes/file/" + note.note_reference,
                {
                    method: "GET",
                    headers: {
                        Authorization:
                            "bearer " +
                            window.localStorage.getItem("authToken"),
                    },
                }
            ).catch((err) => {
                console.error(err);
                success = false;
            });

            if (success) {
                const pdfResponseJSON = await pdfResponse.json();
                if (pdfResponseJSON.pdf_data) {
                    const pdfOjbect = {
                        note_id: note.note_id,
                        pdf_data: pdfResponseJSON.pdf_data,
                    };
                    setPDFs((arr) => [...arr, pdfOjbect]);
                } else {
                    console.error("Could not load note pdf.");
                }
            }
        }
    };
    useEffect(() => {
        getUserInfo();
        getMyFriends();
        getMyGroups();
        getUserNotesRef();
    }, []);

    return (
        <>
            <Template
                showSearch={true}
                showProfile={true}
                blur
                refreshProfileImage={refreshProfileImage}
                body={
                    <div className="d-flex flex-column page-container">
                        <SectionTitle title={`Hi ${user}!`} />
                        <div className="d-flex flex-row">
                            <div className="d-flex flex-column left-panel justify-content-between agenda align-items-center">
                                <div className="d-flex flex-column">
                                    <ModalHeader title="My Friends" />
                                    {friends.map((friend) => (
                                        <h3>{friend.username}</h3>
                                    ))}
                                    <Button
                                        small
                                        title="ADD FRIEND"
                                        type="primary"
                                        clicked={() => setButtonAddFriend(true)}
                                    />
                                    <ModalHeader title="My Groups" />
                                    {groups.map((group) => (
                                        <Link
                                            to={"/group/" + group.group_id}
                                            style={{
                                                color: "inherit",
                                                textDecoration: "inherit",
                                            }}
                                        >
                                            <h3>{group.group_name}</h3>{" "}
                                        </Link>
                                    ))}
                                    <Button
                                        small
                                        title="CREATE GROUP"
                                        type="primary"
                                        clicked={() => setButtonGroup(true)}
                                    />
                                </div>
                                <Button
                                    small
                                    title="EDIT PROFILE"
                                    type="secondary"
                                    clicked={() =>
                                        setButtonPersonalProfile(true)
                                    }
                                />
                            </div>
                            <div className="d-flex flex-column right-panel">
                                <SectionTitle title="Biography" />
                                <div className="agenda big">
                                    {userDescription}
                                </div>
                                <SectionTitle title="My Notes" />
                                <div className="d-flex flex-row flex-wrap mynote-results-container">
                                    {pdfs.map((eachPDF) => (
                                        <Link
                                            to={"/note/" + eachPDF.note_id}
                                            style={{
                                                color: "inherit",
                                                textDecoration: "inherit",
                                            }}
                                        >
                                            <PDFViewer
                                                thumbnail
                                                pdf={eachPDF.pdf_data}
                                            />
                                        </Link>
                                    ))}
                                </div>
                                <div className="d-flex flex-row justify-content-center">
                                    <Button
                                        title="UPLOAD"
                                        type="primary"
                                        clicked={() => setButtonUpload(true)}
                                    />
                                </div>
                            </div>
                        </div>

                        <AddFriendWindow
                            friends={friends}
                            setFriends={setFriends}
                            trigger={buttonAddFriend}
                            setTrigger={setButtonAddFriend}
                        />
                        {buttonGroup ? (
                            <ModalWindow
                                blur
                                body={
                                    <CreateGroupWindow
                                        groups={groups}
                                        setGroups={setGroups}
                                        setTrigger={setButtonGroup}
                                    />
                                }
                            />
                        ) : (
                            <></>
                        )}
                        <UploadNotesWindow
                            trigger={buttonUpload}
                            setTrigger={setButtonUpload}
                        />
                        <PersonalPrefWindow
                            trigger={buttonPersonalProfile}
                            setTrigger={setButtonPersonalProfile}
                            setBio={setUserDescription}
                            setRefreshProfileImage={setRefreshProfileImage}
                        />
                    </div>
                }
            />
        </>
    );
}

export default PersonalPage;
