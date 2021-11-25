import Button from "./button";
import InputBox from "./inputBox";
import ModalHeader from "./modalHeader.jsx";
import { useState, useEffect } from "react";
import { useHistory } from "react-router-dom";
import UploadDropzone from "./uploadDropzone";

function SignUpWindow({ edit, setTrigger }) {
    let history = useHistory();
    const [first_name, setFirstname] = useState("");
    const [last_name, setLastname] = useState("");
    const [username, setUsername] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");
    const [description, setDescription] = useState("");
    const [file, setFile] = useState(null);
    const [fileUpdated, setFileUpdated] = useState(false);

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

        let success = true;

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
        ).catch((err) => {
            console.error(err);
            success = false;
        });

        if (success) {
            const responseJSON = await response.json();
            if (responseJSON.username) {
                if (file != null) {
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
                        await uploadProfileImg(true);
                    } else if (loginResponseJSON.status_code === 401) {
                        alert("Incorrect username or password.");
                    } else {
                        alert("Something went wrong");
                    }
                } else {
                    history.push("/login");
                }
            } else if (responseJSON.errMessage === "username already exists") {
                alert("That username already exists, please try again.");
            } else {
                alert("Something went wrong");
            }
        }
    };

    const uploadProfileImg = async (gotoLogin) => {
        let success = true;

        let formData = new FormData();
        formData.append("userImage", file);

        const imgResponse = await fetch("http://localhost:3000/v1/images", {
            method: "POST",
            headers: {
                Authorization:
                    "bearer " + window.localStorage.getItem("authToken"),
            },
            body: formData,
        }).catch((err) => {
            console.error(err);
            success = false;
        });

        if (success) {
            const imgResponseJSON = await imgResponse.json();
            if (!imgResponseJSON.message) {
                alert("Something went wrong with Image upload!");
            } else if (gotoLogin) {
                history.push("/login");
            }
        }
    };

    const getUserInfo = async () => {
        let success = true;
        const userInfoResponse = await fetch("http://localhost:3000/v1/user", {
            method: "GET",
            headers: {
                Authorization:
                    "bearer " + window.localStorage.getItem("authToken"),
            },
        }).catch((err) => {
            console.error(err);
            success = false;
        });

        if (success) {
            const userInfoResponseJSON = await userInfoResponse.json();
            if (userInfoResponseJSON.email) {
                setFirstname(userInfoResponseJSON.first_name);
                setLastname(userInfoResponseJSON.last_name);
                setEmail(userInfoResponseJSON.email);
                setDescription(userInfoResponseJSON.description);

                await getImage();
            } else {
                console.error("Could not load user info.");
            }
        }
    };

    const getImage = async () => {
        let success = true;
        const imgResponse = await fetch("http://localhost:3000/v1/images", {
            method: "GET",
            headers: {
                Authorization:
                    "bearer " + window.localStorage.getItem("authToken"),
            },
        }).catch((err) => {
            console.error(err);
            success = false;
        });

        if (success) {
            const imgResponseJSON = await imgResponse.json();
            if (imgResponseJSON.name) {
                setFile({ path: imgResponseJSON.name });
            } else {
                console.error("Could not load profile image.");
            }
        }
    };

    const updateInfo = async () => {
        let success = true;
        const updateInfoResponse = await fetch(
            "http://localhost:3000/v1/user",
            {
                method: "PATCH",
                headers: {
                    "Content-Type": "application/json",
                    Authorization:
                        "bearer " + window.localStorage.getItem("authToken"),
                },
                body: JSON.stringify({
                    description,
                    email,
                    first_name,
                    last_name,
                }),
            }
        ).catch((err) => {
            console.error(err);
            success = false;
        });

        if (success) {
            const updateInfoResponseJSON = await updateInfoResponse.json();
            if (updateInfoResponseJSON.username) {
                if (password) {
                    if (password !== confirmPassword) {
                        alert("Passwords don't match.");
                    } else {
                        await updatePassword();
                    }
                }
                setTrigger(false);
            } else {
                alert("Could not update your profile information.");
            }
        }
    };

    const updatePassword = async () => {
        let success = true;
        const updatePasswordResponse = await fetch(
            "http://localhost:3000/v1/password/" + password,
            {
                method: "PATCH",
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
            const updatePasswordResponseJSON =
                await updatePasswordResponse.json();
            if (updatePasswordResponseJSON.username) {
                if (fileUpdated) {
                    await uploadProfileImg(false);
                }
            } else {
                alert("Could not update your password.");
            }
        }
    };

    const updateFile = (newFile) => {
        setFileUpdated(true);
        setFile(newFile);
    };

    useEffect(() => {
        if (edit) {
            getUserInfo();
        }
    }, [edit]);

    return (
        <div className="d-flex flex-column align-items-center">
            <ModalHeader
                title={edit ? "Update Your Info" : "Create New Account"}
            />
            <div className="d-flex flex-row">
                <div className="d-flex flex-column">
                    <InputBox
                        placeholder="first name"
                        onChange={setFirstname}
                        initVal={first_name}
                        dataCy="first-name-input"
                    />
                    <InputBox
                        placeholder="last name"
                        onChange={setLastname}
                        initVal={last_name}
                        dataCy="last-name-input"
                    />
                </div>
                <UploadDropzone file={file} setFile={updateFile} profile />
            </div>
            <InputBox
                placeholder="email"
                onChange={setEmail}
                size="large"
                initVal={email}
                dataCy="email-input"
            />
            {edit ? (
                <></>
            ) : (
                <InputBox
                    placeholder="username"
                    onChange={setUsername}
                    size="large"
                    dataCy="username-input"
                />
            )}
            <InputBox
                placeholder={`${edit ? "new" : ""} password`}
                onChange={setPassword}
                size="large"
                dataCy="pwd-input"
                mask
            />
            <InputBox
                placeholder={`confirm ${edit ? "new" : ""} password`}
                onChange={setConfirmPassword}
                size="large"
                dataCy="confirm-pwd-input"
                mask
            />
            <InputBox
                textArea
                placeholder="biography"
                onChange={setDescription}
                size="large"
                initVal={description}
                dataCy="bio-input"
            />
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
                <div className="d-flex flex-row">
                    <Button
                        title="BACK"
                        type="secondary"
                        url="/login"
                        dataCy="back-btn"
                    />
                    <Button
                        title="SIGN UP"
                        type="primary"
                        clicked={attemptSignup}
                        dataCy="signup-btn"
                    />
                </div>
            )}
        </div>
    );
}

export default SignUpWindow;
