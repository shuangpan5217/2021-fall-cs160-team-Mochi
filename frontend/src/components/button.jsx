import { Link } from "react-router-dom";

function Button({ title, type, url, clicked, small, dataCy }) {
    if (url)
        return (
            <Link
                className={`agenda btn ${type} ${small ? "small" : small}`}
                to={url}
                data-cy={dataCy}
            >
                {title}
            </Link>
        );
    if (clicked)
        return (
            <button
                className={`agenda btn ${type} ${small ? "small" : small}`}
                onClick={() => clicked()}
                data-cy={dataCy}
            >
                {title}
            </button>
        );
    return <></>;
}

export default Button;
