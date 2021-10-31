function UserListItem({ img, name }) {
    return (
        <div className="d-flex flex-row">
            <img
                src={img}
                alt="user"
                style={{ width: "50px", height: "50px" }}
            />
            <p>{name}</p>
        </div>
    );
}

export default UserListItem;
