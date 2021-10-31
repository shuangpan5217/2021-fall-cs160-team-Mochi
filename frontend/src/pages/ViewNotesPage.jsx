import { useEffect, useState } from "react";
import { useParams } from "react-router";
import NoteActions from "../components/noteActions";
import PDFViewer from "../components/PDFViewer";
import SectionTitle from "../components/sectionTitle";
import Template from "../components/template";

function ViewNotesPage() {
    const { noteId } = useParams();
    const [pdf, setPDF] = useState({});
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

                    // setTags(noteResponseJSON.tags.split(","));
                    setTags("test1,test2,test3".split(","));
                    setMembers("test1,test2,test3".split(","));
                    setComments("test1,test2,test3".split(","));
                }
            } else {
                console.error("Could not load note.");
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
                    <div className="d-flex flex-row">
                        <div className="d-flex flex-column">
                            <SectionTitle
                                title="Math Notes"
                                subtitle="by Mochi"
                            />
                            <PDFViewer pdf={pdf} />
                        </div>
                        <NoteActions
                            tags={tags}
                            members={members}
                            comments={comments}
                        />
                    </div>
                }
            />
        </>
    );
}

export default ViewNotesPage;
