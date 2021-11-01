function CommentListItem({ img, name, comment }) {
    return (
        <div className="d-flex flex-row">
            <img src={img} alt="user" className="user-img" />
            <div className="d-flex flex-column">
                <p className="agenda small user-name">{name}</p>
                <p className="agenda small comment">{comment}</p>
            </div>
        </div>
    );
}

export default CommentListItem;
