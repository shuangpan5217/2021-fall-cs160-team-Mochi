import "../css/forms.css";

function InputBox({ placeholder }) {
  return (
    <>
      <input type="text" className="agenda text-input-box" placeholder={placeholder} />
    </>
  );
}

export default InputBox;
