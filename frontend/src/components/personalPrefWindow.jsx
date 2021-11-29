import ModalWindow from "./modalWindow";
import SignUpWindow from "./signUpWindow";

function PersonalPrefWindow({
    trigger,
    setTrigger,
    setBio,
    setRefreshProfileImage,
}) {
    return trigger ? (
        <ModalWindow
            blur
            body={
                <SignUpWindow
                    edit
                    setTrigger={setTrigger}
                    setBio={setBio}
                    setRefreshProfileImage={setRefreshProfileImage}
                />
            }
        />
    ) : (
        <></>
    );
}

export default PersonalPrefWindow;
