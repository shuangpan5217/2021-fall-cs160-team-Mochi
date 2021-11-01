function TabButton({ title, selected, clicked }) {
    return (
        <>
            <button
                className={`agenda btn tab-btn ${selected ? "selected" : ""}`}
                onClick={clicked}
            >
                {title}
            </button>
        </>
    );
}

export default TabButton;
