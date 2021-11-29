import ModalWindow from "./modalWindow";
import CreateGroupWindow from "./createGroupWindow";

function GroupPrefWindow({ trigger, setTrigger, setBio, setName, groupId }) {
    return trigger ? (
        <ModalWindow
            blur
            body={
                <CreateGroupWindow
                    edit
                    setTrigger={setTrigger}
                    setBio={setBio}
                    setName={setName}
                    groupId={groupId}
                />
            }
        />
    ) : (
        <></>
    );
}

export default GroupPrefWindow;
