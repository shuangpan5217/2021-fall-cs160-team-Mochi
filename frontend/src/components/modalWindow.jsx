import "../css/modal.css";

function ModalWindow({ body, blur }) {
    return (
        <div className={blur ? "blur-background" : ""}>
            <div className="modal-container centered">{body}</div>
        </div>
    );
}

export default ModalWindow;
