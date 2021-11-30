import { useEffect, useState } from "react";
import { useParams } from "react-router";
import NoteActions from "../components/noteActions";
import PDFViewer from "../components/pdfViewer";
import SectionTitle from "../components/sectionTitle";
import "../css/viewNotePage.css";
import Template from "../components/template";

function ViewNotesPage(props) {
    const { noteId } = useParams();
    const [pdf, setPDF] = useState({});
    const [title, setTitle] = useState("");
    const [descr, setDescr] = useState("");
    const [owner, setOwner] = useState({});
    const [type, setType] = useState("");
    const [tags, setTags] = useState([]);
    const [comments, setComments] = useState([]);
    const [hasAccess, setHasAccess] = useState(true);

    const getNoteData = async () => {
        let success = true;
        const noteResponse = await fetch(
            "http://localhost:3000/v1/notes/" + noteId,
            {
                method: "GET",
                headers: {
                    Authorization:
                        "bearer " + window.localStorage.getItem("authToken"),
                },
            }
        ).catch((err) => {
            console.error(err);
            success = false;
        });

        if (success) {
            const noteResponseJSON = await noteResponse.json();
            if (noteResponseJSON.note_reference) {
                setTitle(noteResponseJSON.title);
                setDescr(noteResponseJSON.description);
                setType(noteResponseJSON.type);
                setTags([noteResponseJSON.style]);
                if (noteResponseJSON.tag) {
                    setTags([
                        noteResponseJSON.style,
                        ...noteResponseJSON.tag.split(","),
                    ]);
                }
                await getPDF(noteResponseJSON.note_reference);
                await getOwnerImage(noteResponseJSON.note_owner);
                await getCommentData();
            } else if (
                noteResponseJSON.errMessage ===
                "Forbidden: No access to the note"
            ) {
                setHasAccess(false);
            } else {
                console.error("Could not load note.");
            }
        }
    };

    const getPDF = async (noteRef) => {
        let success = true;
        const pdfResponse = await fetch(
            "http://localhost:3000/v1/notes/file/" + noteRef,
            {
                method: "GET",
                headers: {
                    Authorization:
                        "bearer " + window.localStorage.getItem("authToken"),
                },
            }
        ).catch((err) => {
            console.error(err);
            success = false;
        });

        if (success) {
            const pdfResponseJSON = await pdfResponse.json();
            if (pdfResponseJSON.pdf_data) {
                setPDF(pdfResponseJSON.pdf_data);
            } else {
                console.error("Could not load note pdf.");
            }
        }
    };

    const getOwnerImage = async (noteOwner) => {
        let success = true;
        const imgResponse = await fetch("http://localhost:3000/v1/images", {
            method: "GET",
            headers: {
                Authorization:
                    "bearer " + window.localStorage.getItem("authToken"),
            },
        }).catch((err) => {
            console.error(err);
            success = false;
        });

        if (success) {
            const imgResponseJSON = await imgResponse.json();
            if (imgResponseJSON.user_image != null) {
                setOwner({ name: noteOwner, img: imgResponseJSON });
            } else {
                console.error("Could not load profile image.");
            }
        }
    };

    const getCommentData = async () => {
        let success = true;
        const commentResponse = await fetch(
            "http://localhost:3000/v1/notes/" + noteId + "/comments",
            {
                method: "GET",
                headers: {
                    Authorization:
                        "bearer " + window.localStorage.getItem("authToken"),
                },
            }
        ).catch((err) => {
            console.error(err);
            success = false;
        });

        if (success) {
            const commentResponseJSON = await commentResponse.json();
            if (commentResponseJSON.comments) {
                await getUserImages(commentResponseJSON.comments.reverse());
            } else {
                console.error("Could not load comments of this note.");
            }
        }
    };

    const getUserImages = async (commentArr) => {
        const users = [
            ...new Set(
                commentArr.map((commentElem) => ({
                    username: commentElem.username,
                }))
            ),
        ];

        let success = true;
        const imgResponse = await fetch(
            "http://localhost:3000/v1/images/multiple",
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization:
                        "bearer " + window.localStorage.getItem("authToken"),
                },
                body: JSON.stringify({
                    users,
                }),
            }
        ).catch((err) => {
            console.error(err);
            success = false;
        });

        if (success) {
            const imgResponseJSON = await imgResponse.json();
            if (imgResponseJSON.images) {
                let userImgs = {};
                for (let imgObj of imgResponseJSON.images) {
                    userImgs[imgObj.name] = {
                        user_image: imgObj.user_image,
                        type: imgObj.type,
                    };
                }
                for (let commentObj of commentArr) {
                    commentObj.img = userImgs[commentObj.username];
                }
                setComments(commentArr);
            } else {
                console.error(
                    "Could not load profile images for the comments of this note."
                );
            }
        }
    };

    useEffect(() => {
        getNoteData();
    }, []);

    return (
        <>
            <Template
                showSearch={true}
                showProfile={true}
                body={
                    <div className="d-flex flex-row justify-content-center">
                        {hasAccess ? (
                            <>
                                <div className="d-flex flex-column left-container">
                                    <SectionTitle
                                        title={title}
                                        subtitle={`by ${owner.name}`}
                                    />
                                    <PDFViewer pdf={pdf} />
                                </div>
                                <NoteActions
                                    title={title}
                                    descr={descr}
                                    type={type}
                                    tags={tags}
                                    comments={comments}
                                    noteId={noteId}
                                    owner={owner}
                                />
                            </>
                        ) : (
                            <p className="agenda">
                                You don't have access to this note.
                            </p>
                        )}
                    </div>
                }
            />
        </>
    );
}

export default ViewNotesPage;
