import Button from "./button";
import InputBox from "./inputBox";
import ModalHeader from "./modalHeader.jsx";
import { useHistory } from "react-router-dom";
import { useState } from "react";
import ModalWindow from "./modalWindow";
import "../css/personalPage.css";

function CreateGroupWindow({ trigger, setTrigger, groups, setGroups }) {
    const history = useHistory();
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

    return trigger ? (
        <ModalWindow
            blur
            body={
                <div className="d-flex flex-column align-items-center">
                    <ModalHeader title="Create a Group" />
                    <div className="d-flex flex-column align-items-end">
                        <InputBox placeholder="name" onChange={setGroupName} />
                        <InputBox
                            placeholder="description"
                            onChange={setGroupDescription}
                        />
                    </div>
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
                </div>
            }
        />
    ) : (
        ""
    );
}

export default CreateGroupWindow;
