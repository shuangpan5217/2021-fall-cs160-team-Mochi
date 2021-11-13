function Tag({ title, onClick }) {
    if (onClick) {
        return (
            <div className="agenda tag" onClick={() => onClick(title)}>
                {title}
            </div>
        );
    }
    return <div className="agenda tag">{title}</div>;
}

export default Tag;
