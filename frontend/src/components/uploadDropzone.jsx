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

const Container2 = styled.div`
    flex: 1;
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 20px;
    border-width: 2px;
    border-radius: 50%;
    width: 100px;
    height: 100px;
    border-color: ${(props) => getColor(props)};
    border-style: dashed;
    background-color: #fafafa;
    color: #bdbdbd;
    outline: none;
    transition: border 0.24s ease-in-out;
`;

function UploadDropzone({ setFile, profile }) {
    const onDrop = useCallback(
        (acceptFiles) => {
            setFile(acceptFiles[0]);
        },
        [setFile]
    );
    const acceptType = profile ? "image/jpeg, image/png" : "application/pdf";
    const {
        acceptedFiles,
        getRootProps,
        getInputProps,
        isDragActive,
        isDragAccept,
        isDragReject,
    } = useDropzone({ accept: acceptType, maxFile: 1, onDrop });

    return (
        <div className="d-flex flex-column align-items-center">
            <div className="container agenda">
                {profile ? (
                    <Container2
                        {...getRootProps({
                            isDragActive,
                            isDragAccept,
                            isDragReject,
                        })}
                    >
                        <input {...getInputProps()} />
                        {acceptedFiles.length > 0 ? (
                            <p className="smaller text-center">{acceptedFiles[0].path}</p>
                        ) : (
                            <p className="smaller text-center">profile photo</p>
                        )}
                    </Container2>
                ) : (
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
                            click to select files (PDF only, one file per upload
                            )
                        </p>
                    </Container>
                )}
            </div>
            {acceptedFiles.length > 0 && !profile ? (
                <p className="agenda small">{acceptedFiles[0].path}</p>
            ) : (
                <></>
            )}
        </div>
    );
}
export default UploadDropzone;
