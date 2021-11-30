import { useHistory } from "react-router-dom";
import Header from "./header.jsx";

function Template({
    noAuth,
    body,
    showSearch,
    showProfile,
    refreshProfileImage,
    blur,
}) {
    const history = useHistory();

    if (noAuth == null && window.localStorage.getItem("authToken") === "") {
        history.push("/login");
    }

    return (
        <div
            className={`d-flex flex-column full-width full-height ${
                blur ? "blur-bkgrd" : ""
            }`}
        >
            <Header
                showSearch={showSearch}
                showProfile={showProfile}
                refreshProfileImage={refreshProfileImage}
            />
            {body}
        </div>
    );
}

export default Template;
