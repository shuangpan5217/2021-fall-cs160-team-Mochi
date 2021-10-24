import { useMemo, useCallback } from "react";
import {useDropzone} from 'react-dropzone';

function UploadDropzone({setFile}){
  const dropzoneStyle = {
    flex: 1,
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    padding: '20px',
    borderWidth: 3,
    borderRadius: 2,
    fontSize: 18,
    borderColor: '#eeeeee',
    borderStyle: 'dashed',
    backgroundColor: '#fafafa',
    color: '#bdbdbd',
    outline: 'none',
    transition: 'border .24s ease-in-out'
  };
    
  const activeStyle = {
    borderColor: '#2196f3'
  };
  
  const acceptStyle = {
    borderColor: '#00e676'
  };
  
  const rejectStyle = {
    borderColor: '#ff1744'
  };
  const onDrop = useCallback(acceptedFiles => {
    setFile(acceptedFiles[0]);
  }, [setFile])
  const {
    acceptedFiles,
    getRootProps,
    getInputProps,
    isDragActive,
    isDragAccept,
    isDragReject
  } = useDropzone({accept: 'application/pdf', maxFile: 1, onDrop});

  const style = useMemo(() => ({
    ...dropzoneStyle,
    ...(isDragActive ? activeStyle : {}),
    ...(isDragAccept ? acceptStyle : {}),
    ...(isDragReject ? rejectStyle : {})
  }), [
    isDragActive,
    isDragReject,
    isDragAccept
  ]);

  const files = acceptedFiles.map(file => (
    <li key={file.path}>
      {file.path}
    </li>
  ));

  return(          
    <>
      <div {...getRootProps({style})}>
      <input {...getInputProps()} />
      <p>Drag and drop here</p>
      <p>or</p>
      <p>click to select files (PDF only, one file per upload )</p>
      </div>
      <ul>{files}</ul>
    </>) 
}
export default UploadDropzone;