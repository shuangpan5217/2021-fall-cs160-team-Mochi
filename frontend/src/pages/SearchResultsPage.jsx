import Template from "../components/template";
import { useContext, useEffect, useState } from "react";
import SectionTitle from "../components/sectionTitle";
import AppContext from "../components/AppContext";
import PDFViewer from "../components/PDFViewer";
import mockPDF from "../media/mockPDF.pdf";
import "../css/searchResultsPage.css";

function SearchResultsPage(props) {
    const myContext = useContext(AppContext);
    const [thumbnails, setThumbnails] = useState([]);

    useEffect(() => {
        // return () => {
        //     myContext.setGlobalFilter("");
        //     myContext.setGlobalQuery("");
        // };
        //this code is breaking the search results page, need to find a different solution in a future PR
        const numThumbnails = 20;
        let tempArr = [];
        for (let i = 0; i < numThumbnails; i++) {
            tempArr.push(mockPDF);
        }
        setThumbnails(tempArr);
    }, []);

    return (
        <>
            <Template
                showSearch={true}
                showProfile={true}
                body={
                    <div className="d-flex flex-column align-items-start search-results-container">
                        <SectionTitle title="Search Results (20+)" />
                        {/* <p>Filter: {myContext.filter}</p>
                        <p>Query: {myContext.query}</p> */}
                        <div className="d-flex flex-row flex-wrap">
                            {thumbnails.map((pdf) => (
                                <PDFViewer thumbnail pdf={pdf} />
                            ))}
                        </div>
                    </div>
                }
            />
        </>
    );
}

export default SearchResultsPage;
