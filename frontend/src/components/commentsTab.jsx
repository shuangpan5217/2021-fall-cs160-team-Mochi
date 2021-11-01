import { useState } from "react";
import CommentListItem from "./commentListItem";
import InputBox from "./inputBox";
import dummyProfile from "../media/mochi.jpeg";
import NoteActionButton from "./noteActionButton";

function CommentsTab({ comments, noteId, owner }) {
    const [newComment, setNewComment] = useState("");
    const [commentElems, setCommentElems] = useState(
        comments.map((comment) => (
            <CommentListItem
                img={dummyProfile}
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
                        img={dummyProfile}
                        name={owner}
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
                />
                <NoteActionButton title="Post" onClick={postComment} />
            </div>
            {commentElems}
        </div>
    );
}

export default CommentsTab;
