import { Document, Page } from "react-pdf/dist/esm/entry.webpack";
import { useState } from "react";

function PDFViewer({ pdf }) {
    const [numPages, setNumPages] = useState(null);
    const [pageNumber, setPageNumber] = useState(1);

    function onDocumentLoadSuccess({ newNumPages }) {
        setNumPages(newNumPages);
    }

    return (
        <>
            <Document
                file={"data:application/pdf;base64," + pdf}
                onLoadSuccess={onDocumentLoadSuccess}
            >
                <Page pageNumber={pageNumber} />
            </Document>
            <p>
                Page {pageNumber} of {numPages}
            </p>
        </>
    );
}

export default PDFViewer;
