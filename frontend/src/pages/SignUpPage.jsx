import SignUpWindow from "../components/signUpWindow.jsx";
import ModalWindow from "../components/modalWindow";
import Template from "../components/template.jsx";

function SignUpPage(props) {
  return (
    <>
      <Template body={<ModalWindow body={<SignUpWindow />} />} />
    </>
  );
}

export default SignUpPage;
