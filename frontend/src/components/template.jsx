import Header from "./header.jsx";

function Template({ body }) {
  return (
    <div className="d-flex flex-column full-width full-height justify-content-between">
      <Header />
      <div>{body}</div>
    </div>
  );
}

export default Template;
