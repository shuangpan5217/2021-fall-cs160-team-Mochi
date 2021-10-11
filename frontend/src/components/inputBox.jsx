import "../css/forms.css";

function InputBox({ placeholder, onChange }) {
  return (
    <>
      <input
        type="text"
        className="agenda text-input-box"
        placeholder={placeholder}
        onChange={(e) => onChange(e.target.value)}
      />
    </>
  );
}

export default InputBox;
