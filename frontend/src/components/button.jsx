import { Link } from "react-router-dom";

function Button({ title, type, url, clicked }) {
  if (url)
    return (
      <Link className={`agenda btn ${type}`} to={url}>
        {title}
      </Link>
    );
  if (clicked)
    return (
      <button className={`agenda btn ${type}`} onClick={() => clicked()}>
        {title}
      </button>
    );
  return <></>;
}

export default Button;
