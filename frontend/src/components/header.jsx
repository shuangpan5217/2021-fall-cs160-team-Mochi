import Logo from "./logo.jsx";
import "../css/header.css";

function Header(props) {
  return (
    <>
      <div className="logo-spacing">
        <Logo />
      </div>
    </>
  );
}

export default Header;
