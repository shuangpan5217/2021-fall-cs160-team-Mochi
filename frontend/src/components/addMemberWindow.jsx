import Button from "./button";
import InputBox from "./inputBox";
import ModalHeader from "./modalHeader.jsx";
import { useHistory } from "react-router-dom";
import { useState } from "react";
import ModalWindow from "./modalWindow";
import "../css/personalPage.css";
import { useParams } from "react-router";

function AddMemberWindow({ trigger, setTrigger, members, setMembers }) {
    const history = useHistory();
    const [username, setUsername] = useState("");
    const { groupId } = useParams();
    const [users, setUsers] = useState("");

    const attemptAddMember = async () => {
        if (username === "") {
            alert("Please enter a username.");
            return;
        }

        const response = await fetch(
            "http://localhost:3000/v1/groups/" + groupId + "/members",
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
        );

        const responseJSON = await response.json();
        if (responseJSON.group_id) {
            setMembers([...members, { username }]);
            setTrigger(false);
        } else {
            alert("Something went wrong with adding group member!");
        }
    };

    return trigger ? (
        <ModalWindow
            blur
            body={
                <div className="d-flex flex-column align-items-center">
                    <ModalHeader title="Add member" />
                    <div className="d-flex flex-column align-items-end">
                        <InputBox
                            placeholder="Enter a Username"
                            onChange={setUsername}
                        />
                    </div>
                    <br></br>
                    <div className="d-flex flex-row ">
                        <Button
                            title="BACK"
                            type="secondary"
                            clicked={() => setTrigger(false)}
                        />
                        <Button
                            title="ADD"
                            type="primary"
                            clicked={attemptAddMember}
                        />
                    </div>
                </div>
            }
        />
    ) : (
        ""
    );
}

export default AddMemberWindow;
