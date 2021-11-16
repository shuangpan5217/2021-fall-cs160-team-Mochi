import { useState } from "react";
import CommentListItem from "./commentListItem";
import InputBox from "./inputBox";
import NoteActionButton from "./noteActionButton";

function CommentsTab({ comments, noteId, owner }) {
    const [newComment, setNewComment] = useState("");
    const [commentElems, setCommentElems] = useState(
        comments.map((comment) => (
            <CommentListItem
                img={comment.img.user_image}
                type={comment.img.type}
                name={comment.username}
                comment={comment.content}
            />
        ))
    );

    const postComment = async () => {
        let success = true;
        const commentResponse = await fetch(
            "http://localhost:3000/v1/comments/",
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization:
                        "bearer " + window.localStorage.getItem("authToken"),
                },
                body: JSON.stringify({
                    content: newComment,
                    note_id: noteId,
                }),
            }
        ).catch((err) => {
            console.error(err);
            success = false;
        });

        if (success) {
            const commentResponseJSON = await commentResponse.json();
            if (commentResponseJSON.comment_id) {
                const newCommentElem = (
                    <CommentListItem
                        img={owner.img.user_image}
                        type={owner.img.type}
                        name={owner.name}
                        comment={newComment}
                    />
                );
                setCommentElems([newCommentElem, ...commentElems]);
                setNewComment("");
            } else {
                console.error("Could not post comment.");
            }
        }
    };

    return (
        <div className="d-flex flex-column full-width">
            <div className="d-flex flex-column align-items-center full-width">
                <InputBox
                    placeholder="comment"
                    onChange={setNewComment}
                    textArea
                    clear={newComment === ""}
                    fullWidth
                />
                <NoteActionButton title="Post" onClick={postComment} />
            </div>
            <div className="scrollable-container">{commentElems}</div>
        </div>
    );
}

export default CommentsTab;
