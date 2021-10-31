function CommentListItem({ img, name, comment }) {
    return (
        <div className="d-flex flex-row">
            <img
                src={img}
                alt="user"
                style={{ width: "50px", height: "50px" }}
            />
            <div className="d-flex flex-column">
                <p>{name}</p>
                <p>{comment}</p>
            </div>
        </div>
    );
}

export default CommentListItem;
