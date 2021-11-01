import { Document, Page } from "react-pdf/dist/esm/entry.webpack";
import { useState } from "react";

function PDFViewer({ pdf }) {
    const [numPages, setNumPages] = useState(null);

    function onDocumentLoadSuccess({ newNumPages }) {
        setNumPages(newNumPages);
    }

    return (
        <>
            <Document
                file={"data:application/pdf;base64," + pdf}
                onLoadSuccess={onDocumentLoadSuccess}
                className="pdf-container"
            >
                <Page pageNumber={1} className="pdf-page" scale={1.5}/>
            </Document>
        </>
    );
}

export default PDFViewer;
