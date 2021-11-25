import ModalWindow from "./modalWindow";
import SignUpWindow from "./signUpWindow";

function PersonalPrefWindow({ trigger, setTrigger }) {
    return trigger ? <ModalWindow blur body={<SignUpWindow edit setTrigger={setTrigger}/>} /> : <></>;
}

export default PersonalPrefWindow;
