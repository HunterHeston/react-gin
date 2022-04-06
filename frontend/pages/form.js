import { useState } from "react";
export default function Form() {
  const [localURL, setLocalURL] = useState("");
  const [file, setFile] = useState(null);

  const submitFunction = (e) => {
    console.log(localURL);
    console.log(file);

    const form = new FormData();

    // 'file' matches the name the backend is expecting.
    form.append("file", file);
    form.append("test", "empty");

    fetch("http://localhost:8080/file-upload", {
      method: "POST",
      mode: "cors",
      body: form,
    })
      .then((res) => console.log(res))
      .catch((err) => console.error(err));

    e.preventDefault();
  };

  const imageChange = (e) => {
    setFile(e?.target?.files[0]);
    setLocalURL(URL.createObjectURL(e?.target?.files[0]));
  };

  return (
    <div>
      <h1>Upload a a photo</h1>
      <form action="submit" onSubmit={submitFunction}>
        <input type="file" onChange={imageChange} />
        <button type="submit">Upload!</button>
      </form>
      <div>
        <p>preview:</p>
        <img alt="preview of the image to be uploaded" src={localURL}></img>
      </div>
    </div>
  );
}
