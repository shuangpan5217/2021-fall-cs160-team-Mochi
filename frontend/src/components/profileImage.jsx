import { useEffect, useState } from "react";
import { useHistory } from "react-router-dom";
import defaultImg from "../media/default.jpeg";

function ProfileImage(props) {
    const history = useHistory();
    const [img, setImg] = useState({});
    const [type, setType] = useState("");

    const getImage = async () => {
        let success = true;
        const imgResponse = await fetch("http://localhost:3000/v1/images", {
            method: "GET",
            headers: {
                Authorization:
                    "bearer " + window.localStorage.getItem("authToken"),
            },
        }).catch((err) => {
            console.error(err);
            success = false;
        });

        if (success) {
            const imgResponseJSON = await imgResponse.json();
            if (imgResponseJSON.user_image != null) {
                setImg(imgResponseJSON.user_image);
                setType(imgResponseJSON.type);
            } else {
                console.error("Could not load profile image.");
            }
        }
    };

    useEffect(() => {
        getImage();
    }, []);

    return (
        <img
            src={img === "" ? defaultImg : `data:image/${type};base64, ${img}`}
            alt="profile"
            className="profile-img"
            onClick={() => history.push("/my_notes")}
        />
    );
}

export default ProfileImage;
