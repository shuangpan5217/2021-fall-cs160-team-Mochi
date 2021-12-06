import Template from "../components/template";
import Button from "../components/button";
import { useContext, useEffect, useState } from "react";
import AppContext from "../components/appContext";
import { useHistory } from "react-router-dom";
import SectionTitle from "../components/sectionTitle";
import PDFViewer from "../components/pdfViewer";
import "../css/searchResultsPage.css";

function SearchResultsPage(props) {
    const history = useHistory();
    const myContext = useContext(AppContext);
    const [count, setCount] = useState(0);
    const [thumbnails, setThumbnails] = useState([]);
    const [filteredThumbnails, setFilteredThumbnails] = useState([]);
    const [offset, setOffset] = useState(0);
    const [showLoadMore, setShowLoadMore] = useState(true);
    const [styles, setStyles] = useState({});

    const getSearchResults = async () => {
        if (!myContext.query) {
            return;
        }

        let success = true;
        const searchResponse = await fetch(
            `http://localhost:3001/v1/notes/search/${myContext.query}?offset=${offset}`, //hardcoded tag
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

                const newStyles = {};
                searchResponseJSON.notes.forEach(
                    (note) => (newStyles[note.note_id] = note.style)
                );
                setStyles({ ...styles, ...newStyles });

                const titles = {};
                searchResponseJSON.notes.forEach(
                    (note) => (titles[note.note_id] = note.title)
                );

                let noteRefs = searchResponseJSON.notes.map((note) => ({
                    path: note.note_reference,
                }));
                await getPDFs(noteRefs, titles);
            } else {
                console.error("Could not load search results.");
            }
        }
    };

    const getPDFs = async (filePaths, titles) => {
        let success = true;
        const pdfResponse = await fetch(
            "http://localhost:3001/v1/notes/files",
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
                    ...pdfResponseJSON.filesData.map((fileData) => ({
                        ...fileData,
                        title: titles[fileData.note_id],
                    })),
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

    useEffect(() => {
        const newThumbnails = thumbnails.filter(
            (pdf) =>
                myContext.filter === "" ||
                styles[pdf.note_id] === myContext.filter
        );

        newThumbnails.forEach((pdf) => console.log(pdf));

        setFilteredThumbnails(newThumbnails);
        setCount(newThumbnails.length);
    }, [thumbnails]);

    return (
        <>
            <Template
                showSearch={true}
                showProfile={true}
                showFilterBtn={true}
                blur
                body={
                    <div className="d-flex flex-column align-items-start search-results-container">
                        <SectionTitle title={`Search Results (${count})`} />
                        {count > 0 ? (
                            <>
                                <div className="d-flex flex-row justify-content-center flex-wrap full-width">
                                    {filteredThumbnails.map((pdf) => (
                                        <PDFViewer
                                            thumbnail
                                            pdf={pdf.pdf_data}
                                            title={pdf.title}
                                            onClick={() =>
                                                history.push(
                                                    "/note/" + pdf.note_id
                                                )
                                            }
                                        />
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
                            </>
                        ) : (
                            <p className="agenda">
                                No results found for the search:&nbsp;
                                {myContext.query}
                                {myContext.filter === ""
                                    ? ""
                                    : ` and the filter ${myContext.filter}`}
                            </p>
                        )}
                    </div>
                }
            />
        </>
    );
}

export default SearchResultsPage;
