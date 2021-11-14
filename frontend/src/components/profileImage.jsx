import { useHistory } from "react-router-dom";

function ProfileImage(props) {
    const history = useHistory();

    const type = "waiting";
    const img = "waiting";

    return (
        <img
            src={`data:image/${type};base64, ${img}`}
            alt="profile"
            className="profile-img"
            onClick={() => history.push("/my_notes")}
        />
    );
}

export default ProfileImage;
