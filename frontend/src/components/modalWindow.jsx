import "../css/modal.css"

function ModalWindow({body}) {
  return (
    <div className="modal-container centered">
    {body}
    </div>
  );
}

export default ModalWindow;
