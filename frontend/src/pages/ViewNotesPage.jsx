import { useEffect, useState } from "react";
import { useParams } from "react-router";
import NoteActions from "../components/noteActions";
import PDFViewer from "../components/PDFViewer";
import SectionTitle from "../components/sectionTitle";
import "../css/viewNotePage.css";
import Template from "../components/template";

function ViewNotesPage(props) {
    const { noteId } = useParams();
    const [pdf, setPDF] = useState({});
    const [title, setTitle] = useState("");
    const [descr, setDescr] = useState("");
    const [owner, setOwner] = useState("");
    const [type, setType] = useState("");
    const [tags, setTags] = useState([]);
    const [members, setMembers] = useState([]);
    const [comments, setComments] = useState([]);

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
                setOwner(noteResponseJSON.note_owner);
                setType(noteResponseJSON.type);
                setTags([noteResponseJSON.style]);
                if (noteResponseJSON.tag) {
                    setTags([
                        noteResponseJSON.style,
                        ...noteResponseJSON.tag.split(","),
                    ]);
                }

                const pdfResponse = await fetch(
                    "http://localhost:3000/v1/notes/file/" +
                        noteResponseJSON.note_reference,
                    {
                        method: "GET",
                        headers: {
                            Authorization:
                                "bearer " +
                                window.localStorage.getItem("authToken"),
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
                } else {
                    return;
                }
            } else {
                console.error("Could not load note.");
            }
        }

        getCommentData();
        // getMemberData();
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
                setComments(commentResponseJSON.comments.reverse());
            } else {
                console.error("Could not load shared members of this note.");
            }
        }
    };

    const getMemberData = async () => {
        let success = true;
        const memberResponse = await fetch(
            "http://localhost:3000/v1/notes/" + noteId + "/members",
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
            const memberResponseJSON = await memberResponse.json();
            if (memberResponseJSON.users) {
                setMembers(memberResponseJSON.users);
            } else {
                console.error("Could not load shared members of this note.");
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
                        <div className="d-flex flex-column left-container">
                            <SectionTitle
                                title={title}
                                subtitle={`by ${owner}`}
                            />
                            <PDFViewer pdf={pdf} />
                        </div>
                        <NoteActions
                            title={title}
                            descr={descr}
                            type={type}
                            tags={tags}
                            members={members}
                            comments={comments}
                            noteId={noteId}
                            owner={owner}
                        />
                    </div>
                }
            />
        </>
    );
}

export default ViewNotesPage;
