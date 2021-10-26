import { useHistory } from "react-router-dom";
import "../css/logo.css";

function Logo(props) {
    const history = useHistory();

    return (
        <>
            <p className="continuo logo" onClick={() => history.push("/home")}>
                <span className="continuo-lg">M</span>ochi
                <span className="continuo-lg">N</span>ote
            </p>
        </>
    );
}

export default Logo;
