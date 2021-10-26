import Header from "./header.jsx";

function Template({ body, showSearch, showProfile}) {
  return (
    <div className="d-flex flex-column full-width full-height justify-content-between">
      <Header showSearch={showSearch} showProfile={showProfile}/>
      <div>{body}</div>
    </div>
  );
}

export default Template;
