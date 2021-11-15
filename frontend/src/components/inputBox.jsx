import { useEffect, useState } from "react";
import "../css/forms.css";

function InputBox({
    placeholder,
    onChange,
    label,
    textArea,
    mask,
    clear,
    fullWidth,
    size,
    onEnter,
}) {
    const [value, setValue] = useState("");
    useEffect(() => {
        if (clear) {
            setValue("");
        }
    }, [clear]);

    useEffect(() => {
        onChange(value);
    }, [value, onChange]);

    const handleKeyDown = (event) => {
        if (onEnter && event.key === "Enter" && event.target.value !== "") {
            onEnter();
        }
    };

    var LabelElem = <></>;

    if (label != null) {
        LabelElem = (
            <label className="agenda small label-spacing">{label}</label>
        );
    }

    return (
        <div
            className={`d-flex flex-row align-items-start ${
                fullWidth ? "full-width" : ""
            }`}
        >
            {LabelElem}
            {textArea ? (
                <textarea
                    type="text"
                    className={`agenda text-input-box ${
                        size ? (size === "small" ? "sm-box" : "lg-box") : ""
                    }`}
                    placeholder={placeholder}
                    onChange={(e) => setValue(e.target.value)}
                    rows="3"
                    value={value}
                    onKeyDown={handleKeyDown}
                />
            ) : (
                <input
                    type={mask ? "password" : "text"}
                    className={`agenda text-input-box ${
                        size ? (size === "small" ? "sm-box" : "lg-box") : ""
                    }`}
                    placeholder={placeholder}
                    onChange={(e) => setValue(e.target.value)}
                    value={value}
                    onKeyDown={handleKeyDown}
                />
            )}
        </div>
    );
}

export default InputBox;
