import Button from "./button";
import InputBox from "./inputBox";
import ModalHeader from "./modalHeader.jsx";
import { useHistory } from "react-router-dom";
import { useState } from "react";
import ModalWindow from "./modalWindow";
import "../css/personalPage.css";

function CreateGroupWindow({ trigger, setTrigger }) {
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
            setTrigger(false);
            history.push("/my_notes");
        } else {
            alert("Something went wrong with creating group!");
            setTrigger(false);
        }
    };

    return trigger ? (
        <div className="popup-add">
            <ModalWindow
                body={
                    <div className="d-flex flex-column align-items-center">
                        <ModalHeader title="Create a Group" />
                        <div className="d-flex flex-column align-items-end">
                            <InputBox
                                placeholder="name"
                                onChange={setGroupName}
                            />
                        </div>
                        <div className="d-flex flex-column align-items-end">
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
        </div>
    ) : (
        ""
    );
}

export default CreateGroupWindow;
