function NoteActionButton({ title, onClick }) {
    return <button className="btn primary note-action-btn" onClick={onClick}>{title}</button>;
}

export default NoteActionButton;
