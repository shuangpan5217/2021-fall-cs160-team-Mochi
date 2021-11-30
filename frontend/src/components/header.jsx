import Logo from "./logo.jsx";
import "../css/header.css";
import SearchBar from "./searchBar.jsx";
import ProfileImage from "./profileImage.jsx";

function Header({ showSearch, showProfile, refreshProfileImage}) {
    return (
        <div className="d-flex flex-row justify-content-between">
            <div className="logo-spacing">
                <Logo canClick={showProfile}/>
            </div>
            {showSearch ? <SearchBar showFilterBtn={true}/> : <></>}
            {showProfile ? <ProfileImage refresh={refreshProfileImage}/> : <></>}
        </div>
    );
}

export default Header;
