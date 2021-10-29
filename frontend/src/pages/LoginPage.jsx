import LoginWindow from "../components/loginWindow";
import ModalWindow from "../components/modalWindow";
import Template from "../components/template.jsx";

function LoginPage(props) {
    return (
        <>
            <Template
                noAuth
                showSearch={false}
                showProfile={false}
                body={<ModalWindow body={<LoginWindow />} />}
            />
        </>
    );
}

export default LoginPage;
