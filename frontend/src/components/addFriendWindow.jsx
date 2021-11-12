import Button from "./button";
import InputBox from "./inputBox";
import ModalHeader from "./modalHeader.jsx";
import { useHistory } from "react-router-dom";
import { useState } from "react";
import ModalWindow from "./modalWindow";
import "../css/personalPage.css";

function AddFriendWindow({ trigger, setTrigger }) {
    const history = useHistory();
    const [username2, setUsername2] = useState("");

    const attemptAddFriend = async () => {
        if (username2 === "") {
            alert("Please enter your friend's username.");
            return;
        }

        const response = await fetch("http://localhost:3000/v1/friends/", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization:
                    "bearer " + window.localStorage.getItem("authToken"),
            },
            body: JSON.stringify({
                username2,
            }),
        });

        const responseJSON = await response.json();
        if (responseJSON.username) {
            alert("Succefully adding friend");
            history.push("/my_notes");
        } else {
            alert("Something went wrong with adding friend!");
            setTrigger(false);
        }
    };

    return trigger ? (
        <div className="popup-addfriend">
            <ModalWindow
                body={
                    <div className="d-flex flex-column align-items-center">
                        <ModalHeader title="Add friend" />
                        <div className="d-flex flex-column align-items-end">
                            <InputBox
                                placeholder="Enter a Username"
                                onChange={setUsername2}
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
                                clicked={attemptAddFriend}
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

export default AddFriendWindow;
