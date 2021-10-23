import "../css/uploadNotesWindow.css";

function RadioButton({ onChange, label, group }) {
  return (
    <label className="agenda small radio-button" >
        <input type="radio" name={group}
        onChange={(e) => onChange(e.target.value)}/> {label} 
    </label>
  );
}

export default RadioButton;
