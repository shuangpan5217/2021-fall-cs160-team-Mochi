import NoteActionHeader from "./noteActionHeader";
import Tag from "./tag";

function InfoTab({ tags }) {
    const tagElems = tags.map((tag) => <Tag title={tag} />);
    return (
        <div className="d-flex flex-column">
            <NoteActionHeader title="Math Notes" />
            <p>description</p>
            <NoteActionHeader title="Tags" />
            <div className="d-flex flex-row">
                {tagElems}
            </div>
        </div>
    );
}

export default InfoTab;
