import { useState } from "react";

function LoginPage(props) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();

    const response = await fetch("http://localhost:3000/v1/login?signup=true", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        username,
        password,
      }),
    });

    const responseJSON = await response.json();
    if (responseJSON.username) {
      alert("You have successfully signed up, " + responseJSON.username + "!");
    } else if (responseJSON.errMessage === "username already exists") {
      alert("That username already exists, please try again.");
    } else {
      alert("Something went wrong");
    }
  };
  return (
    <>
      <main className="form-signup">
        <form onSubmit={handleSubmit}>
          <h1 className="h3 mb-3 fw-normal">Please sign up</h1>
          <input
            type="username"
            className="form-control"
            placeholder="Username"
            required
            onChange={(e) => setUsername(e.target.value)}
          />
          <input
            type="password"
            className="form-control"
            placeholder="Password"
            required
            onChange={(e) => setPassword(e.target.value)}
          />
          <button className="w-100 btn btn-lg btn-primary" type="submit">
            Submit
          </button>
        </form>
      </main>
    </>
  );
}

export default LoginPage;
