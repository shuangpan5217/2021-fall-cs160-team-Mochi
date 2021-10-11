import Button from "./button";
import InputBox from "./inputBox";
import ModalHeader from "./modalHeader.jsx";
import { useState } from "react";
import { useHistory } from "react-router-dom";

function SignUpWindow(props) {
  let history = useHistory();

  const [first_name, setFirstname] = useState("");
  const [last_name, setLastname] = useState("");
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const attemptSignup = async () => {
    const response = await fetch("http://localhost:3000/v1/login?signup=true", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        first_name,
        last_name,
        username,
        email,
        password,
      }),
    });

    const responseJSON = await response.json();
    if (responseJSON.username) {
      history.push("/home");
    } else if (responseJSON.errMessage === "username already exists") {
      alert("That username already exists, please try again.");
    } else {
      alert("Something went wrong");
    }
  };
  return (
    <div className="d-flex flex-column align-items-center">
      <ModalHeader title="Create New Account" />
      <InputBox placeholder="First name" onChange={setFirstname} />
      <InputBox placeholder="Last name" onChange={setLastname} />
      <InputBox placeholder="Email" onChange={setEmail} />
      <InputBox placeholder="Username" onChange={setUsername} />
      <InputBox placeholder="Password" onChange={setPassword} />
      <div className="d-flex flex-row">
        <Button title="BACK" type="secondary" url="/login" />
        <Button title="SIGN UP" type="primary" clicked={attemptSignup} />
      </div>
    </div>
  );
}

export default SignUpWindow;
