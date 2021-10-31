function TabButton({ title, selected, clicked }) {
    return (
        <>
            <button
                className={`agenda btn ${selected ? "selected" : ""}`}
                onClick={clicked}
            >
                {title}
            </button>
        </>
    );
}

export default TabButton;
