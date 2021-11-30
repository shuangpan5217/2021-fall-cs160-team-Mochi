function ModalHeader({ title, small }) {
    return (
        <>
            <p
                className={`agenda section-header ${
                    small ? "modal-section-header" : ""
                }`}
            >
                {title}
            </p>
        </>
    );
}

export default ModalHeader;
