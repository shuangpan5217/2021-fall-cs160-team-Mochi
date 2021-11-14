import Button from "./button";
import InputBox from "./inputBox";
import ModalHeader from "./modalHeader.jsx";
import { useState, useEffect } from "react";
import { useHistory } from "react-router-dom";
import UploadDropzone from "./uploadDropzone";

function SignUpWindow(props) {
    let history = useHistory();
    const [first_name, setFirstname] = useState("");
    const [last_name, setLastname] = useState("");
    const [username, setUsername] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");
    const [description, setDescription] = useState("");
    const [file, setFile] = useState(null);

    useEffect(() => {
        window.localStorage.setItem("authToken", "");
    }, []);

    const attemptSignup = async () => {
        if (
            first_name === "" ||
            last_name === "" ||
            username === "" ||
            email === "" ||
            password === ""
        ) {
            alert("Please fill out all fields.");
            return;
        } else if (password !== confirmPassword) {
            alert("Passwords don't match.");
            return;
        }

        const response = await fetch(
            "http://localhost:3000/v1/login?signup=true",
            {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({
                    first_name,
                    last_name,
                    username,
                    email,
                    password,
                    description,
                }),
            }
        );

        const responseJSON = await response.json();
        if (responseJSON.username) {
            const loginResponse = await fetch(
                "http://localhost:3000/v1/login",
                {
                    method: "POST",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify({
                        username,
                        password,
                    }),
                }
            );

            const loginResponseJSON = await loginResponse.json();
            if (!loginResponseJSON.status_code) {
                window.localStorage.setItem(
                    "authToken",
                    loginResponseJSON.token
                );
                let formData = new FormData();
                formData.append("userImage", file);

                const imgResponse = await fetch(
                    "http://localhost:3000/v1/images",
                    {
                        method: "POST",
                        headers: {
                            Authorization:
                                "bearer " +
                                window.localStorage.getItem("authToken"),
                        },
                        body: formData,
                    }
                );

                const imgResponseJSON = await imgResponse.json();
                if (!imgResponseJSON.message) {
                    alert("Something went wrong with Image upload!");
                } else {
                    history.push("/login");
                }
            } else if (loginResponseJSON.status_code === 401) {
                alert("Incorrect username or password.");
            } else {
                alert("Something went wrong");
            }
        } else if (responseJSON.errMessage === "username already exists") {
            alert("That username already exists, please try again.");
        } else {
            alert("Something went wrong");
        }
    };

    return (
        <div className="d-flex flex-column align-items-center">
            <ModalHeader title="Create New Account" />
            <div className="d-flex flex-row">
                <div className="d-flex flex-column">
                    <InputBox
                        placeholder="first name"
                        onChange={setFirstname}
                    />
                    <InputBox placeholder="last name" onChange={setLastname} />
                </div>
                <UploadDropzone setFile={setFile} profile />
            </div>
            <InputBox placeholder="email" onChange={setEmail} size="large" />
            <InputBox
                placeholder="username"
                onChange={setUsername}
                size="large"
            />
            <InputBox
                placeholder="password"
                onChange={setPassword}
                size="large"
                mask
            />
            <InputBox
                placeholder="confirm password"
                onChange={setConfirmPassword}
                size="large"
                mask
            />
            <InputBox
                textArea
                placeholder="biography"
                onChange={setDescription}
                size="large"
            />
            <div className="d-flex flex-row">
                <Button title="BACK" type="secondary" url="/login" />
                <Button
                    title="SIGN UP"
                    type="primary"
                    clicked={attemptSignup}
                />
            </div>
        </div>
    );
}

export default SignUpWindow;
