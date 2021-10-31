import Template from "../components/template";
import { useContext, useEffect } from "react";
import AppContext from "../components/AppContext";

function SearchResultsPage(props) {
  const myContext = useContext(AppContext);

  useEffect(() => {
    return () => {
      myContext.setGlobalFilter("");
      myContext.setGlobalQuery("");
    };
  }, []);

  return (
    <>
      <Template
        showSearch={true}
        showProfile={true}
        body={
          <>
            <h1>Search Results Page</h1>
            <p>Filter: {myContext.filter}</p>
            <p>Query: {myContext.query}</p>
          </>
        }
      />
    </>
  );
}

export default SearchResultsPage;
