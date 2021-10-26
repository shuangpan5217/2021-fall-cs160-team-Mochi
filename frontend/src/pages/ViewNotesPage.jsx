import Template from "../components/template";

function ViewNotesPage(props) {
    return (
        <>
            <Template
                showSearch={true}
                showProfile={true}
                body={<h1>View Notes Page</h1>}
            />
        </>
    );
}

export default ViewNotesPage;
