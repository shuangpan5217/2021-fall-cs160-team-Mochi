import Button from "./button";
import InputBox from "./inputBox";
import ModalHeader from "./modalHeader";

function LoginWindow(props) {
  return (
    <div className="d-flex flex-column align-items-center">
      <ModalHeader title="Welcome to MochiNote!" />
      <InputBox placeholder="username" />
      <InputBox placeholder="password" />
      <Button
        title="forgot your password?"
        type="link"
        clicked={() => alert("Username: test, Password: 1234")}
      />
      <div className="d-flex flex-row">
        <Button title="SIGN UP" type="secondary" url="/signup" />
        <Button title="LOG IN" type="primary" url="/home" />
      </div>
    </div>
  );
}

export default LoginWindow;
