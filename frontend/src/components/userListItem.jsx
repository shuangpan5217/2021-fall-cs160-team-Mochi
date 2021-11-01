function UserListItem({ img, name }) {
    return (
        <div className="d-flex flex-row align-items-center">
            <img src={img} alt="user" className="user-img" />
            <p className="agenda small user-name">{name}</p>
        </div>
    );
}

export default UserListItem;
