import LoginWindow from "../components/loginWindow";
import ModalWindow from "../components/modalWindow";
import Template from "../components/template.jsx";

function LoginPage(props) {
  return (
    <>
      <Template body={<ModalWindow body={<LoginWindow />} />} />
    </>
  );
}

export default LoginPage;
