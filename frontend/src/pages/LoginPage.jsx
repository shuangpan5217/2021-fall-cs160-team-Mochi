import LoginWindow from "../components/loginWindow";
import ModalWindow from "../components/modalWindow";
import Template from "../components/template.jsx";

function LoginPage({setAuthToken}) {
  return (
    <>
      <Template showSearch={false} showProfile={false} body={<ModalWindow body={<LoginWindow setAuthToken={setAuthToken} />} />} />
    </>
  );
}

export default LoginPage;
