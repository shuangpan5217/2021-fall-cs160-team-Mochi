import UserListItem from "./userListItem";
import InputBox from "./inputBox";
import NoteActionHeader from "./noteActionHeader";
import { useEffect, useState } from "react";
import NoteActionButton from "./noteActionButton";

function SharingTab({ noteId, type }) {
    const [username, setUsername] = useState("");
    const [members, setMembers] = useState([]);
    const [isOwner, setIsOwner] = useState(true);

    const getMemberData = async () => {
        let success = true;
        const memberResponse = await fetch(
            "http://localhost:3000/v1/notes/" + noteId + "/members",
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
            const memberResponseJSON = await memberResponse.json();
            if (memberResponseJSON.users) {
                await getUserImages(memberResponseJSON.users);
                setUsername("");
            } else if (
                memberResponseJSON.errMessage ===
                "no access to members of the note"
            ) {
                setIsOwner(false);
            } else {
                console.error("Could not load shared members of this note.");
            }
        }
    };

    const getUserImages = async (memberUsers) => {
        const users = memberUsers.map((member) => ({
            username: member.username,
        }));

        let success = true;
        const imgResponse = await fetch(
            "http://localhost:3000/v1/images/multiple",
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization:
                        "bearer " + window.localStorage.getItem("authToken"),
                },
                body: JSON.stringify({
                    users,
                }),
            }
        ).catch((err) => {
            console.error(err);
            success = false;
        });

        if (success) {
            const imgResponseJSON = await imgResponse.json();
            if (imgResponseJSON.images) {
                let userImgs = {};
                for (let imgObj of imgResponseJSON.images) {
                    userImgs[imgObj.name] = {
                        user_image: imgObj.user_image,
                        type: imgObj.type,
                    };
                }
                for (let member of memberUsers) {
                    member.img = userImgs[member.username];
                }
                setMembers(memberUsers);
            } else {
                console.error(
                    "Could not load profile images for the shared members of this note."
                );
            }
        }
    };

    const addMember = async () => {
        let success = true;
        const memberResponse = await fetch(
            `http://localhost:3000/v1/notes/${noteId}/members`,
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization:
                        "bearer " + window.localStorage.getItem("authToken"),
                },
                body: JSON.stringify({
                    users: [{ username }],
                }),
            }
        ).catch((err) => {
            console.error(err);
            success = false;
        });

        if (success) {
            const memberResponseJSON = await memberResponse.json();
            if (memberResponseJSON.note_id) {
                await getMemberData();
            } else {
                alert("Could not add this member");
            }
        }
    };

    useEffect(() => {
        getMemberData();
    }, []);

    return (
        <div className="d-flex flex-column full-width">
            <NoteActionHeader title="Shared with" />
            {type === "shared" ? (
                <>
                    {isOwner ? (
                        <div className="d-flex flex-column align-items-center full-width">
                            <InputBox
                                placeholder="member"
                                onChange={setUsername}
                                fullWidth
                                clear={username === ""}
                            />
                            <NoteActionButton
                                title="Share"
                                onClick={addMember}
                            />
                        </div>
                    ) : (
                        <p className="agenda">
                            You don't have access to see shared members.
                        </p>
                    )}
                    <div className="scrollable-container">
                        {members.map((member) => (
                            <UserListItem
                                img={member.img.user_image}
                                type={member.img.type}
                                name={member.username}
                            />
                        ))}
                    </div>
                </>
            ) : (
                <p className="agenda">This note is public!</p>
            )}
        </div>
    );
}

export default SharingTab;
