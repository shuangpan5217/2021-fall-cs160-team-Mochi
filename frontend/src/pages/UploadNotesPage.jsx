import UploadNotesWindow from "../components/uploadNotesWindow.jsx";
import ModalWindow from "../components/modalWindow";
import Template from "../components/template.jsx";

function UploadNotesPage(props) {
  return (
    <>
      <Template body={<ModalWindow body={<UploadNotesWindow />} />} />
    </>
  );
}

export default UploadNotesPage;