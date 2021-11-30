import { Document, Page } from "react-pdf/dist/esm/entry.webpack";
import { useState } from "react";
import "../css/pdfViewer.css";

function PDFViewer({ pdf, thumbnail, onClick }) {
    const [pages, setPages] = useState(null);

    function onDocumentLoadSuccess({ numPages }) {
        setPages(numPages);
    }

    return (
        <div
            className={`${thumbnail ? "thumbnail-wrapper" : ""} pdf-wrapper`}
            onClick={onClick}
        >
            <Document
                file={"data:application/pdf;base64," + pdf}
                onLoadSuccess={onDocumentLoadSuccess}
                className={thumbnail ? "" : "pdf-container"}
            >
                {thumbnail ? (
                    <Page pageNumber={1} className="pdf-page" scale={0.6} />
                ) : (
                    Array(pages)
                        .fill()
                        .map((_, i) => (
                            <Page
                                pageNumber={i + 1}
                                className="pdf-page"
                                scale={1.5}
                            />
                        ))
                )}
            </Document>
        </div>
    );
}

export default PDFViewer;
