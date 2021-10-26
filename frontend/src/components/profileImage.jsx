import { useHistory } from "react-router-dom";

function ProfileImage({ url }) {
    const history = useHistory();

    return (
        <img
            src={url}
            alt="profile"
            className="profile-img"
            onClick={() => history.push("/my_notes")}
        />
    );
}

export default ProfileImage;
