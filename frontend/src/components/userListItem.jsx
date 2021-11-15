import defaultImg from "../media/default.jpeg";

function UserListItem({ img, type, name }) {
    return (
        <div className="d-flex flex-row align-items-center">
            <img
                src={
                    img === ""
                        ? defaultImg
                        : `data:image/${type};base64, ${img}`
                }
                alt="user"
                className="user-img"
            />
            <p className="agenda small user-name">{name}</p>
        </div>
    );
}

export default UserListItem;
