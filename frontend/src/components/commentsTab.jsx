import { useState } from "react";
import CommentListItem from "./commentListItem";
import InputBox from "./inputBox";
import dummyProfile from "../media/mochi.jpeg";
import NoteActionButton from "./noteActionButton";

function CommentsTab({ comments }) {
    const [addCommment, setAddComment] = useState("");
    const commentElems = comments.map((comment) => (
        <CommentListItem img={dummyProfile} name="dummy" comment={comment} />
    ));

    return (
        <div className="d-flex flex-column">
            <div className="d-flex flex-column align-items-center">
                <InputBox
                    placeholder="comment"
                    onChange={setAddComment}
                    textArea
                />
                <NoteActionButton
                    title="Post"
                    onClick={() => console.log("post")}
                />
            </div>
            {commentElems}
        </div>
    );
}

export default CommentsTab;
