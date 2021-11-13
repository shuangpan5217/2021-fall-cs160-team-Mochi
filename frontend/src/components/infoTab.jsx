import NoteActionHeader from "./noteActionHeader";
import Tag from "./tag";

function InfoTab({ title, descr, tags }) {
    const tagElems = tags.map((tag) => <Tag title={tag} key={tag} />);
    return (
        <div className="d-flex flex-column full-width">
            <NoteActionHeader title={title} />
            <p className="agenda">{descr}</p>
            <NoteActionHeader title="Tags" />
            <div className="d-flex flex-row">{tagElems}</div>
        </div>
    );
}

export default InfoTab;
