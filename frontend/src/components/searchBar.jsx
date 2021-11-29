import { useContext, useState, useEffect } from "react";
import { useHistory } from "react-router-dom";
import AppContext from "./appContext";
import RadioButton from "./radioButton";
import "../css/searchbar.css";

function SearchBar({ showFilterBtn, updateSearch }) {
    const history = useHistory();
    const myContext = useContext(AppContext);
    const [showFilters, toggleFilters] = useState(
        (myContext.filter !== "") & showFilterBtn
    );
    const [query, setQuery] = useState(myContext.query);

    useEffect(() => {
        if (updateSearch) {
            myContext.setGlobalQuery(query);
            myContext.setGlobalFilter("");
            history.push("/search");
        }
    }, [updateSearch, history, myContext, query]);

    const updateFilter = (filter) => {
        myContext.setGlobalQuery(query);
        myContext.setGlobalFilter(filter);
    };

    const handleKeyDown = (event) => {
        if (event.key === "Enter" && query !== "") {
            myContext.setGlobalQuery(query);
            history.push("/search");
        }
    };

    return (
        <div className="d-flex flex-column">
            <div
                className={`d-flex flex-row ${
                    showFilterBtn ? "search-bar" : "search-bar-big"
                }`}
            >
                <input
                    type="text"
                    className={`agenda search-input ${
                        showFilters ? "filter" : ""
                    } ${showFilterBtn ? "" : "no-filter-btn"}`}
                    placeholder={
                        showFilterBtn
                            ? "Search by title, tag, or description"
                            : "Search Notes"
                    }
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
                        label="All Styles"
                        onChange={() => updateFilter("")}
                        checked={myContext.filter === ""}
                    />
                    <RadioButton
                        group="style"
                        label="Outline"
                        onChange={() => updateFilter("Outline")}
                        checked={myContext.filter === "Outline"}
                    />
                    <RadioButton
                        group="style"
                        label="Cornell"
                        onChange={() => updateFilter("Cornell")}
                        checked={myContext.filter === "Cornell"}
                    />
                    <RadioButton
                        group="style"
                        label="Boxing"
                        onChange={() => updateFilter("Boxing")}
                        checked={myContext.filter === "Boxing"}
                    />
                    <RadioButton
                        group="style"
                        label="Charting"
                        onChange={() => updateFilter("Charting")}
                        checked={myContext.filter === "Charting"}
                    />
                    <RadioButton
                        group="style"
                        label="Mapping"
                        onChange={() => updateFilter("Mapping")}
                        checked={myContext.filter === "Mapping"}
                    />
                    <RadioButton
                        group="style"
                        label="Sentence"
                        onChange={() => updateFilter("Sentence")}
                        checked={myContext.filter === "Sentence"}
                    />
                </div>
            ) : (
                <></>
            )}
        </div>
    );
}

export default SearchBar;
