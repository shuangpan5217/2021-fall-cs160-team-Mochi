import Template from "../components/template";

function PersonalPage(props) {
    return (
        <>
            <Template
                showSearch={true}
                showProfile={true}
                body={<h1>Personal Page</h1>}
            />
        </>
    );
}

export default PersonalPage;
