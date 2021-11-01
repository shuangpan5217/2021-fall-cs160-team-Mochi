import { useEffect, useState } from "react";
import "../css/forms.css";

function InputBox({ placeholder, onChange, label, textArea, mask, clear }) {
    const [value, setValue] = useState("");
    useEffect(() => {
        if (clear) {
            setValue("");
        }
    }, [clear]);

    useEffect(() => {
        onChange(value);
    }, [value, onChange]);

    var LabelElem = <></>;

    if (label != null) {
        LabelElem = <label className="agenda small">{label}&nbsp;</label>;
    }

    if (textArea != null) {
        return (
            <div className="d-flex flex-row align-items-center full-width">
                {LabelElem}
                <textarea
                    type="text"
                    className="agenda text-input-box"
                    placeholder={placeholder}
                    onChange={(e) => setValue(e.target.value)}
                    rows="3"
                    value={value}
                />
            </div>
        );
    } else {
        return (
            <div className="d-flex flex-row align-items-center full-width">
                {LabelElem}
                <input
                    type={mask ? "password" : "text"}
                    className="agenda text-input-box"
                    placeholder={placeholder}
                    onChange={(e) => setValue(e.target.value)}
                    value={value}
                />
            </div>
        );
    }
}

export default InputBox;
