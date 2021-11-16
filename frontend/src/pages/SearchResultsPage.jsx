import Template from "../components/template";
import { useContext, useEffect } from "react";
import SectionTitle from "../components/sectionTitle";
import AppContext from "../components/AppContext";
import "../css/searchResultsPage.css";

function SearchResultsPage(props) {
    const myContext = useContext(AppContext);

    useEffect(() => {
        // return () => {
        //     myContext.setGlobalFilter("");
        //     myContext.setGlobalQuery("");
        // };
        //this code is breaking the search results page, need to find a different solution in a future PR
    }, []);

    return (
        <>
            <Template
                showSearch={true}
                showProfile={true}
                body={
                    <div className="d-flex flex-column align-items-start search-results-container">
                        <SectionTitle title="Search Results (20+)" />
                        <p>Filter: {myContext.filter}</p>
                        <p>Query: {myContext.query}</p>
                    </div>
                }
            />
        </>
    );
}

export default SearchResultsPage;
