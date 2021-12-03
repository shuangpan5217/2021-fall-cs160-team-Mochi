import { Document, Page } from "react-pdf/dist/esm/entry.webpack";
import { useState } from "react";
import "../css/pdfViewer.css";

function PDFViewer({ title, pdf, thumbnail, onClick }) {
    const [pages, setPages] = useState(null);

    function onDocumentLoadSuccess({ numPages }) {
        setPages(numPages);
    }

    return (
        <div className="pdf-wrapper" onClick={onClick}>
            <Document
                file={"data:application/pdf;base64," + pdf}
                onLoadSuccess={onDocumentLoadSuccess}
                className={thumbnail ? "" : "pdf-container"}
            >
                {thumbnail ? (
                    <div className="thumbnail-wrapper">
                        <Page pageNumber={1} className="pdf-page" scale={0.6} />
                        <div className="overlay agenda">{title}</div>
                    </div>
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
