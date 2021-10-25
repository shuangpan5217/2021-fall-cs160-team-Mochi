import "../css/forms.css";

function InputBox({ placeholder, onChange, label, textArea }) {
  var LabelElem = (<></>);
  if (label != null) {
    LabelElem = (       
      <label className="agenda small">
        {label }&nbsp; 
      </label>
    )
  }
  if(textArea != null) {
    return (
      <div className="d-flex flex-row align-items-center">
      {LabelElem}
        <textarea
          type="text"
          className="agenda text-input-box"
          placeholder={placeholder}
          onChange={(e) => onChange(e.target.value)}
          rows = "3"
        />
      </div>
    );
  } else {
  return (
    <div className="d-flex flex-row align-items-center">
    {LabelElem}
      <input
        type="text"
        className="agenda text-input-box"
        placeholder={placeholder}
        onChange={(e) => onChange(e.target.value)}
      />
    </div>
  );
  }
}

export default InputBox;
