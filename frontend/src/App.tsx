import './App.css';
import React, { SyntheticEvent, useState } from 'react';

function App() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const submit = async (e:SyntheticEvent) => {
      e.preventDefault();
        console.log({
          email,
          password
        })
  }

  return (
    <div className="App">
      <main className="form-signup">
        <form onSubmit={submit}>
            <h1 className="h3 mb-3 fw-normal">Please sign up</h1>
            <input type ="email" className="form-control" placeholder="name@example.com" required
                   onChange={e => setEmail(e.target.value)}
            />
            <input type="password" className="form-control" placeholder="Password" required
                   onChange={e => setPassword(e.target.value)}
            />
            <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
        </form>
      </main>
    </div>
  );
}

export default App;
