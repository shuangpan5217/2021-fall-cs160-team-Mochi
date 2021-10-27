import Template from "../components/template";

function PersonalNotePage(props) {
    return (
        <>
            <Template
                showSearch={true}
                showProfile={true}
                body={<h1>Dummy Personal Note Page</h1>}
            />
        </>
    );
}

export default PersonalNotePage;