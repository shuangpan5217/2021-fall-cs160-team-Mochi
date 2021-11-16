import { Link } from "react-router-dom";

function Button({ title, type, url, clicked, small }) {
  if (url)
    return (
      <Link className={`agenda btn ${type} ${small ? "small" : small}`} to={url}>
        {title}
      </Link>
    );
  if (clicked)
    return (
      <button className={`agenda btn ${type} ${small ? "small" : small}`} onClick={() => clicked()}>
        {title}
      </button>
    );
  return <></>;
}

export default Button;
