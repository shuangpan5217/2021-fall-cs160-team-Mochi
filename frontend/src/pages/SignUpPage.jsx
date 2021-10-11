import '../App.css';
import { useState } from "react";
import { useHistory } from "react-router-dom";
import Background from '../images/background.jpg';
import MochiNote from '../images/mochinote.png';

function SignUpPage(props) {
  let history = useHistory();

  const [firstname, setFirstname] = useState("");
  const [lastname, setLastname] = useState("");
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();

    const response = await fetch("http://localhost:3000/v1/login?signup=true", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        firstname,
        lastname,
        username,
        email,
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
    <div>
      <div className="bglayout" style={{background : `url(${Background})`}}>
        <img src={MochiNote} 
          width= "20%" alt= "MochiNote"/>
        <div className="form-signup">
          <form onSubmit={handleSubmit}>
            <h1 className="h6 mb-6 fw-normal">
              Get started with MochiNote today! Create your account by filling out the information below.
            </h1>
            <div className='form-inputs'>
              <label className='form-label'>First name</label>
              <input
                type="text"
                className="form-control"
                placeholder="Enter your first name"
                required
                onChange={(e) => setFirstname(e.target.value)}
              />
            </div>
            <div className='form-inputs'>
              <label className='form-label'>Last name</label>
              <input
                type="text"
                className="form-control"
                placeholder="Enter your last name"
                required
                onChange={(e) => setLastname(e.target.value)}
              />
            </div>
            <div className='form-inputs'>
              <label className='form-label'>Username</label>
              <input
                type="text"
                className="form-control"
                placeholder="Enter your username"
                required
                onChange={(e) => setUsername(e.target.value)}
              />
            </div>
            <div className='form-inputs'>
                <label className='form-label'>Email</label>
              <input
                type="email"
                className="form-control"
                placeholder="Enter your Email"
                required
                onChange={(e) => setEmail(e.target.value)}
              />
            </div>
            <div className='form-inputs'>
                <label className='form-label'>Password</label>
              <input
                type="password"
                className="form-control"
                placeholder="Enter your password"
                required
                onChange={(e) => setPassword(e.target.value)}
              />
            </div>
            <button className="w-50 btn btn-lg btn-primary" type="submit"
              onClick={() => {
                history.push("/home");
              }}
            >
              Submit
            </button>
            <span className='form-input-login'>
            Already have an account? Login <a href='./login'>here</a>
            </span>
          </form>
        </div>
      </div> 
    </div>
  );
}

export default SignUpPage;
