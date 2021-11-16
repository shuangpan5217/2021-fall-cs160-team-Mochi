import defaultImg from "../media/default.jpeg";

function CommentListItem({ img, type, name, comment }) {
    return (
        <div className="d-flex flex-row">
            <img
                src={
                    img === ""
                        ? defaultImg
                        : `data:image/${type};base64, ${img}`
                }
                alt="user"
                className="user-img"
            />
            <div className="d-flex flex-column">
                <p className="agenda small user-name">{name}</p>
                <p className="agenda small comment">{comment}</p>
            </div>
        </div>
    );
}

export default CommentListItem;
