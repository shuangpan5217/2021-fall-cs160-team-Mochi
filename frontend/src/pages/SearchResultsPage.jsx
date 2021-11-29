import Template from "../components/template";
import Button from "../components/button";
import { useContext, useEffect, useState } from "react";
import SectionTitle from "../components/sectionTitle";
import PDFViewer from "../components/pdfViewer";
import "../css/searchResultsPage.css";

function SearchResultsPage(props) {
    const [count, setCount] = useState(0);
    const [thumbnails, setThumbnails] = useState([]);
    const [offset, setOffset] = useState(0);
    const [showLoadMore, setShowLoadMore] = useState(true);

    const getSearchResults = async () => {
        let success = true;
        const searchResponse = await fetch(
            `http://localhost:3000/v1/notes/search/valid?offset=${offset}`, //hardcoded tag
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
            const searchResponseJSON = await searchResponse.json();
            if (searchResponseJSON.notes) {
                setOffset(offset + 20);
                let noteRefs = searchResponseJSON.notes.map((note) => ({
                    path: note.note_reference,
                }));
                await getPDFs(noteRefs);
            } else {
                console.error("Could not load search results.");
            }
        }
    };

    const getPDFs = async (filePaths) => {
        let success = true;
        const pdfResponse = await fetch(
            "http://localhost:3000/v1/notes/files",
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization:
                        "bearer " + window.localStorage.getItem("authToken"),
                },
                body: JSON.stringify({
                    filePaths,
                }),
            }
        ).catch((err) => {
            console.error(err);
            success = false;
        });

        if (success) {
            const pdfResponseJSON = await pdfResponse.json();
            if (pdfResponseJSON.count) {
                if (pdfResponseJSON.count < 20) {
                    setShowLoadMore(false);
                }
                setCount(count + pdfResponseJSON.count);
                setThumbnails([
                    ...thumbnails,
                    ...pdfResponseJSON.filesData.map((file) => file.pdf_data),
                ]);
            } else if (pdfResponseJSON.filesData) {
                setShowLoadMore(false);
            } else {
                console.error("Could not load note pdf.");
            }
        }
    };

    useEffect(() => {
        getSearchResults();
    }, []);

    return (
        <>
            <Template
                showSearch={true}
                showProfile={true}
                body={
                    <div className="d-flex flex-column align-items-start search-results-container">
                        <SectionTitle title={`Search Results (${count})`} />
                        <div className="d-flex flex-row justify-content-center flex-wrap full-width">
                            {thumbnails.map((pdf) => (
                                <PDFViewer thumbnail pdf={pdf} />
                            ))}
                        </div>
                        {showLoadMore ? (
                            <div className="d-flex flex-row justify-content-center full-width">
                                <Button
                                    title="LOAD MORE"
                                    type="primary"
                                    clicked={() => getSearchResults()}
                                />
                            </div>
                        ) : (
                            <></>
                        )}
                    </div>
                }
            />
        </>
    );
}

export default SearchResultsPage;
