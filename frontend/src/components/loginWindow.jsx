import Button from "./button";
import InputBox from "./inputBox";
import ModalHeader from "./modalHeader";
import { useHistory } from "react-router-dom";
import { useState, useEffect } from "react";

function LoginWindow(props) {
    const history = useHistory();
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    useEffect(() => {
        window.localStorage.setItem("authToken", "");
    }, []);

    const attemptLogin = async () => {
        const response = await fetch("http://localhost:3000/v1/login", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
                username,
                password,
            }),
        });

        const responseJSON = await response.json();
        if (!responseJSON.status_code) {
            window.localStorage.setItem("authToken", responseJSON.token);
            history.push("/home");
        } else if (responseJSON.status_code === 401) {
            alert("Incorrect username or password.");
        } else {
            alert("Something went wrong");
        }
    };

    return (
        <div className="d-flex flex-column align-items-center">
            <ModalHeader title="Welcome to MochiNote!" />
            <InputBox placeholder="Username" onChange={setUsername} />
            <InputBox placeholder="Password" onChange={setPassword} mask />
            <div className="d-flex flex-row">
                <Button title="SIGN UP" type="secondary" url="/signup" />
                <Button title="LOG IN" type="primary" clicked={attemptLogin} />
            </div>
        </div>
    );
}

export default LoginWindow;
