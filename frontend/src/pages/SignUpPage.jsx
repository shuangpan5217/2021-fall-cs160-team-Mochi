import SignUpWindow from "../components/signUpWindow.jsx";
import ModalWindow from "../components/modalWindow";
import { useEffect } from "react";
import Template from "../components/template.jsx";

function SignUpPage(props) {
    useEffect(() => {
        window.localStorage.setItem("authToken", "");
    }, []);

    return (
        <>
            <Template
                noAuth
                showSearch={false}
                showProfile={false}
                body={<ModalWindow body={<SignUpWindow />} />}
            />
        </>
    );
}

export default SignUpPage;
