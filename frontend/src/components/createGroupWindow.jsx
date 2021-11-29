import Button from "./button";
import InputBox from "./inputBox";
import ModalHeader from "./modalHeader.jsx";
import { useState, useEffect } from "react";
import ModalWindow from "./modalWindow";
import "../css/personalPage.css";

function CreateGroupWindow({
    setTrigger,
    groups,
    setGroups,
    edit,
    groupId,
    setBio,
}) {
    const [group_name, setGroupName] = useState("");
    const [description, setGroupDescription] = useState("");

    const attemptCreatGroup = async () => {
        if (group_name === "" || description === "") {
            alert("Please enter group name and group description.");
            return;
        }

        const response = await fetch("http://localhost:3000/v1/groups/", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization:
                    "bearer " + window.localStorage.getItem("authToken"),
            },
            body: JSON.stringify({
                group_name,
                description,
            }),
        });

        const responseJSON = await response.json();
        if (responseJSON.group_id) {
            setGroups([...groups, { group_name }]);
            setTrigger(false);
        } else {
            alert("Something went wrong with creating group!");
        }
    };

    const getGroupInfo = async () => {
        let success = true;
        const groupInfoResponse = await fetch(
            "http://localhost:3000/v1/groups/" + groupId,
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
            if (groupInfoResponseJSON.group_id) {
                setGroupName(groupInfoResponseJSON.group_name);
                setGroupDescription(groupInfoResponseJSON.description);
            } else {
                console.error("Could not load group info.");
            }
        }
    };

    const updateInfo = async () => {
        if (group_name === "" || description === "") {
            alert("Please fill out all fields.");
            return;
        }

        let success = true;
        const updateInfoResponse = await fetch(
            "http://localhost:3000/v1/groups/" + groupId,
            {
                method: "PATCH",
                headers: {
                    "Content-Type": "application/json",
                    Authorization:
                        "bearer " + window.localStorage.getItem("authToken"),
                },
                body: JSON.stringify({
                    description,
                    group_name,
                }),
            }
        ).catch((err) => {
            console.error(err);
            success = false;
        });

        if (success) {
            const updateInfoResponseJSON = await updateInfoResponse.json();
            if (updateInfoResponseJSON.group_id) {
                setBio(description);
                setTrigger(false);
            } else {
                alert("Could not update the group information.");
            }
        }
    };

    useEffect(() => {
        if (edit) {
            getGroupInfo();
        }
    }, [edit]);

    return (
        <div className="d-flex flex-column align-items-center">
            <ModalHeader
                title={edit ? "Update Group Info" : "Create New Group"}
            />
            <div className="d-flex flex-column align-items-end">
                <InputBox
                    placeholder="name"
                    onChange={setGroupName}
                    initVal={group_name}
                />
                <InputBox
                    placeholder="description"
                    onChange={setGroupDescription}
                    initVal={description}
                    textArea
                />
            </div>
            {edit ? (
                <div className="d-flex flex-row">
                    <Button
                        title="DISCARD"
                        type="secondary"
                        clicked={() => setTrigger(false)}
                    />
                    <Button title="SAVE" type="primary" clicked={updateInfo} />
                </div>
            ) : (
                <div className="d-flex flex-row ">
                    <Button
                        title="BACK"
                        type="secondary"
                        clicked={() => setTrigger(false)}
                    />
                    <Button
                        title="CREATE"
                        type="primary"
                        clicked={attemptCreatGroup}
                    />
                </div>
            )}
        </div>
    );
}

export default CreateGroupWindow;
