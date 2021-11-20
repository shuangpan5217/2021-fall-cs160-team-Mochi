import { Document, Page } from "react-pdf/dist/esm/entry.webpack";
import { useState } from "react";

function PersonalPagePDFViewer({ pdf }) {
    const [numPages, setNumPages] = useState(null);

    function onDocumentLoadSuccess({ newNumPages }) {
        setNumPages(newNumPages);
    }

    return (
        <div className="thumbnail-wrapper">
            <Document
                file={"data:application/pdf;base64," + pdf}
                onLoadSuccess={onDocumentLoadSuccess}
            >
                <Page pageNumber={1} className="pdf-page" scale={0.75} />
            </Document>
        </div>
    );
}

export default PersonalPagePDFViewer;
