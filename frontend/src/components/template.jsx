import { useHistory } from "react-router-dom";
import Header from "./header.jsx";

function Template({ noAuth, body, showSearch, showProfile }) {
    const history = useHistory();

    if (noAuth == null && window.localStorage.getItem("authToken") === "") {
        history.push("/login");
    }

    return (
        <div className="d-flex flex-column full-width full-height justify-content-between">
            <Header showSearch={showSearch} showProfile={showProfile} />
            <div>{body}</div>
        </div>
    );
}

export default Template;
