import { useCallback } from "react";
import { useDropzone } from "react-dropzone";
import styled from "styled-components";

const getColor = (props) => {
    if (props.isDragAccept) {
        return "#00e676";
    }
    if (props.isDragReject) {
        return "#ff1744";
    }
    if (props.isDragActive) {
        return "#2196f3";
    }
    return "#eeeeee";
};

const Container = styled.div`
    flex: 1;
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 20px;
    border-width: 2px;
    border-radius: 2px;
    border-color: ${(props) => getColor(props)};
    border-style: dashed;
    background-color: #fafafa;
    color: #bdbdbd;
    outline: none;
    transition: border 0.24s ease-in-out;
`;

function UploadDropzone({ setFile }) {
    const onDrop = useCallback(
        (acceptFiles) => {
            setFile(acceptFiles[0]);
        },
        [setFile]
    );
    const {
        acceptedFiles,
        getRootProps,
        getInputProps,
        isDragActive,
        isDragAccept,
        isDragReject,
    } = useDropzone({ accept: "application/pdf", maxFile: 1, onDrop });

    const files = acceptedFiles.map((file) => (
        <li key={file.path}>{file.path}</li>
    ));

    return (
        <>
            <div className="container">
                <Container
                    {...getRootProps({
                        isDragActive,
                        isDragAccept,
                        isDragReject,
                    })}
                >
                    <input {...getInputProps()} />
                    <p>Drag and drop here</p>
                    <p>or</p>
                    <p>
                        click to select files (PDF only, one file per upload )
                    </p>
                </Container>
            </div>
            <ul>{files}</ul>
        </>
    );
}
export default UploadDropzone;