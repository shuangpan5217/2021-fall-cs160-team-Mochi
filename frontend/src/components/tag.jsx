function Tag({ title, onClick }) {
    return (
        <div className="agenda tag" onClick={() => onClick(title)}>
            {title}
        </div>
    );
}

export default Tag;
