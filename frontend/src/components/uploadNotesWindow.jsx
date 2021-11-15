import Button from "./button";
import Tag from "./tag";
import InputBox from "./inputBox";
import ModalHeader from "./modalHeader.jsx";
import { useHistory } from "react-router-dom";
import { useEffect, useState } from "react";
import "../css/uploadNotesWindow.css";
import RadioButton from "./radioButton";
import ModalWindow from "./modalWindow";
import UploadDropzone from "./uploadDropzone";

function UploadNotesWindow({ trigger, setTrigger }) {
    const history = useHistory();
    const [title, setTitle] = useState("");
    const [description, setDescription] = useState("");
    const [tag, setTag] = useState("");
    const [tags, setTags] = useState([]);
    const [tagElems, setTagElems] = useState([]);
    const [type, setType] = useState("");
    const [style, setStyle] = useState("");
    const [file, setFile] = useState(null);

    useEffect(() => {
        const removeTag = (tagTitle) => {
            const tempTags = tags.filter((t) => t !== tagTitle);
            setTags(tempTags);
        };

        setTagElems(
            tags.map((tagTitle) => (
                <Tag title={tagTitle} key={tagTitle} onClick={removeTag} />
            ))
        );
    }, [tags]);

    const addTag = () => {
        if (tags.includes(tag)) {
            alert("Please enter a unique tag.");
        } else {
            setTags([...tags, tag]);
        }
        setTag("");
    };

    const attemptUpload = async () => {
        let formData = new FormData();
        formData.append("noteFile", file);

        const pdfResponse = await fetch("http://localhost:3000/v1/notes/file", {
            method: "POST",
            headers: {
                Authorization:
                    "bearer " + window.localStorage.getItem("authToken"),
            },
            body: formData,
        });

        const pdfResponseJSON = await pdfResponse.json();
        if (!pdfResponseJSON.note_reference) {
            alert("Something went wrong with PDF upload!");
            return;
        }

        const note_reference = pdfResponseJSON.note_reference;

        const response = await fetch("http://localhost:3000/v1/notes/", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization:
                    "bearer " + window.localStorage.getItem("authToken"),
            },
            body: JSON.stringify({
                title,
                description,
                tag: tags.join(","),
                type,
                note_reference,
                style,
            }),
        });

        const responseJSON = await response.json();
        if (responseJSON.note_id) {
            history.push("/note/" + responseJSON.note_id);
        } else {
            alert("Something went wrong with note upload!");
            setTrigger(false);
        }
    };

    return trigger ? (
        <div className="popup">
            <ModalWindow
                body={
                    <div className="d-flex flex-column align-items-center">
                        <ModalHeader title="Upload Notes" />
                        <div className="d-flex flex-column align-items-start">
                            <InputBox
                                label="Title"
                                placeholder="title"
                                onChange={setTitle}
                            />
                            <InputBox
                                textArea
                                label="Description"
                                placeholder="description"
                                onChange={setDescription}
                                size="large"
                            />
                            <div className="d-flex flex-row">
                                <InputBox
                                    label="Tag"
                                    placeholder="tag"
                                    onChange={setTag}
                                    clear={tag === ""}
                                    onEnter={addTag}
                                    size="small"
                                />
                                <div className="d-flex flex-row tag-wrapper flex-wrap">
                                    {tagElems}
                                </div>
                            </div>
                            <div className="d-flex flex-row">
                                <label className="agenda small label-spacing radio-label-spacing">
                                    Sharing
                                </label>
                                <RadioButton
                                    group="sharing"
                                    label="Public"
                                    onChange={() => setType("public")}
                                />
                                <RadioButton
                                    group="sharing"
                                    label="Group"
                                    onChange={() => setType("Group")}
                                />
                                <RadioButton
                                    group="sharing"
                                    label="Private"
                                    onChange={() => setType("Private")}
                                />
                            </div>
                            <div className="d-flex flex-row">
                                <label className="agenda small label-spacing radio-label-spacing">
                                    Style
                                </label>
                                <RadioButton
                                    group="style"
                                    label="Outline"
                                    onChange={() => setStyle("Outline")}
                                />
                                <RadioButton
                                    group="style"
                                    label="Cornell"
                                    onChange={() => setStyle("Cornell")}
                                />
                                <RadioButton
                                    group="style"
                                    label="Boxing"
                                    onChange={() => setStyle("Boxing")}
                                />
                                <RadioButton
                                    group="style"
                                    label="Charting"
                                    onChange={() => setStyle("Charting")}
                                />
                                <RadioButton
                                    group="style"
                                    label="Mapping"
                                    onChange={() => setStyle("Mapping")}
                                />
                                <RadioButton
                                    group="style"
                                    label="Sentence"
                                    onChange={() => setStyle("Sentence")}
                                />
                            </div>
                            <div className="d-flex flex-row">
                                <label className="agenda small label-spacing">
                                    Attach File
                                </label>
                                <UploadDropzone setFile={setFile} />
                            </div>
                        </div>
                        <div className="d-flex flex-row ">
                            <Button
                                title="CANCEL"
                                type="secondary"
                                clicked={() => setTrigger(false)}
                            />
                            <Button
                                title="UPLOAD"
                                type="primary"
                                clicked={attemptUpload}
                            />
                        </div>
                    </div>
                }
            />
        </div>
    ) : (
        ""
    );
}

export default UploadNotesWindow;
