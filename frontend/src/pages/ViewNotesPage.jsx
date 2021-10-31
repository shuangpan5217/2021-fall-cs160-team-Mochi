import { useEffect, useState } from "react";
import { useParams } from "react-router";
import PDFViewer from "../components/PDFViewer";
import Template from "../components/template";

function ViewNotesPage({ authToken }) {
    const { noteId } = useParams();
    const [pdf, setPDF] = useState({});

    const getNoteData = async () => {
        const pdfResponse = await fetch(
            "http://localhost:3000/v1/notes/file/" + noteId,
            {
                method: "GET",
                headers: {
                    Authorization: "bearer " + authToken,
                },
            }
        );

        const pdfResponseJSON = await pdfResponse.json();
        if (pdfResponseJSON.pdf_data) {
            setPDF(pdfResponseJSON.pdf_data);
        } else {
            alert("Could not load note pdf.");
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
                body={<PDFViewer pdf={pdf} />}
            />
        </>
    );
}

export default ViewNotesPage;
