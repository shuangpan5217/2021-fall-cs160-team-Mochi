import { useContext, useState } from "react";
import { useHistory } from "react-router-dom";
import AppContext from "./AppContext";
import RadioButton from "./radioButton";
import "../css/searchbar.css";

function SearchBar({ showFilterBtn }) {
    const history = useHistory();
    const myContext = useContext(AppContext);
    const [showFilters, toggleFilters] = useState(myContext.filter !== "");
    const [query, setQuery] = useState(myContext.query);

    const updateFilter = (filter) => {
        myContext.setGlobalQuery(query);
        myContext.setGlobalFilter(filter);
    };

    const setOutlineStyle = () => {
        updateFilter("Outline");
    };

    const setCornellStyle = () => {
        updateFilter("Cornell");
    };

    const setBoxingStyle = () => {
        updateFilter("Boxing");
    };

    const setChartingStyle = () => {
        updateFilter("Charting");
    };

    const setMappingStyle = () => {
        updateFilter("Mapping");
    };

    const setSentenceStyle = () => {
        updateFilter("Sentence");
    };

    const handleKeyDown = (event) => {
        if (event.key === "Enter" && query !== "") {
            myContext.setGlobalQuery(query);
            history.push("/search");
        }
    };

    if (showFilterBtn == null) {
        return (
            <div className="d-flex flex-row search-bar-big">
                <input
                    type="text"
                    className={`agenda search-input ${
                        showFilters ? "filter" : ""
                    } ${showFilterBtn ? "" : "no-filter-btn"}`}
                    placeholder="Search Notes"
                    value={query}
                    onChange={(e) => setQuery(e.target.value)}
                    onKeyDown={handleKeyDown}
                />
            </div>
        );
    } else { 
        return (
        <div className="d-flex flex-column">
            <div className="d-flex flex-row search-bar">
                <input
                    type="text"
                    className={`agenda search-input ${
                        showFilters ? "filter" : ""
                    } ${showFilterBtn ? "" : "no-filter-btn"}`}
                    placeholder="Search by title, tag, or description"
                    value={query}
                    onChange={(e) => setQuery(e.target.value)}
                    onKeyDown={handleKeyDown}
                />
                {showFilterBtn ? (
                    <button
                        className={`agenda small filter-btn ${
                            showFilters ? "filter" : ""
                        }`}
                        onClick={() => toggleFilters(!showFilters)}
                    >
                        Filter
                    </button>
                ) : (
                    <></>
                )}
            </div>
            {showFilters ? (
                <div className="d-flex flex-row filters-container">
                    <RadioButton
                        group="style"
                        label="Outline"
                        onChange={setOutlineStyle}
                        checked={myContext.filter === "Outline"}
                    />
                    <RadioButton
                        group="style"
                        label="Cornell"
                        onChange={setCornellStyle}
                        checked={myContext.filter === "Cornell"}
                    />
                    <RadioButton
                        group="style"
                        label="Boxing"
                        onChange={setBoxingStyle}
                        checked={myContext.filter === "Boxing"}
                    />
                    <RadioButton
                        group="style"
                        label="Charting"
                        onChange={setChartingStyle}
                        checked={myContext.filter === "Charting"}
                    />
                    <RadioButton
                        group="style"
                        label="Mapping"
                        onChange={setMappingStyle}
                        checked={myContext.filter === "Mapping"}
                    />
                    <RadioButton
                        group="style"
                        label="Sentence"
                        onChange={setSentenceStyle}
                        checked={myContext.filter === "Sentence"}
                    />
                </div>
            ) : (
                <></>
            )}
        </div>
    );
    }
}

export default SearchBar;
