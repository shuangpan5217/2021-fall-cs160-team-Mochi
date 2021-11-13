import Button from "./button";
import InputBox from "./inputBox";
import ModalHeader from "./modalHeader.jsx";
import { useHistory } from "react-router-dom";
import { useState } from "react";
import ModalWindow from "./modalWindow";
import "../css/personalPage.css";

function AddMemberWindow({ trigger, setTrigger }) {
    const history = useHistory();
    const [username, setUsername] = useState("");

    const attemptAddMember = async () => {
        if (username === "") {
            alert("Please enter a username.");
            return;
        }

        const response = await fetch("http://localhost:3000/v1/group/members", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization:
                    "bearer " + window.localStorage.getItem("authToken"),
            },
            body: JSON.stringify({
                username,
            }),
        });

        const responseJSON = await response.json();
        if (responseJSON.username) {
            setTrigger(false);
            history.push("/my_groups");
        } else if (responseJSON.status_code === 404) {
            alert("There is no such group.");
        } else {
            alert("Something went wrong with adding group member!");
            setTrigger(false);
        }
    };

    return trigger ? (
        <div className="popup-add">
            <ModalWindow
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
        </div>
    ) : (
        ""
    );
}

export default AddMemberWindow;
