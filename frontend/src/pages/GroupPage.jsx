import Template from "../components/template";

function GroupPage(props) {
    return (
        <>
            <Template
                showSearch={true}
                showProfile={true}
                body={<h1>Group Page</h1>}
            />
        </>
    );
}

export default GroupPage;
