import Button from "./button";
import InputBox from "./inputBox";
import ModalHeader from "./modalHeader.jsx";
import { useState } from "react";
import ModalWindow from "./modalWindow";
import "../css/personalPage.css";

function AddFriendWindow({ trigger, setTrigger, friends, setFriends }) {
    const [username2, setUsername2] = useState("");

    const attemptAddFriend = async () => {
        if (username2 === "") {
            alert("Please enter your friend's username.");
            return;
        }

        const response = await fetch("http://localhost:3001/v1/friends/", {
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
            setFriends([...friends, { username: username2 }]);
            setTrigger(false);
        } else if (responseJSON.status_code === 404) {
            alert("There is no such username.");
        } else {
            alert("Something went wrong with adding friend!");
        }
    };

    return trigger ? (
        <ModalWindow
            blur
            body={
                <div className="d-flex flex-column align-items-center">
                    <ModalHeader title="Add friend" />
                    <div className="d-flex flex-column align-items-end">
                        <InputBox
                            placeholder="Enter a Username"
                            onChange={setUsername2}
                        />
                    </div>
                    <div className="d-flex flex-row no-top-spacing">
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
    ) : (
        ""
    );
}

export default AddFriendWindow;
