import './App.css';
import { useState } from 'react';

function App() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = e => {
    e.preventDefault();
    console.log(username, password);
  };

  return (
    <div className="App">
      <main className="form-signup">
        <form onSubmit={handleSubmit}>
            <h1 className="h3 mb-3 fw-normal">Please sign up</h1>
            <input type ="username" className="form-control" placeholder="name@example.com" required
                   onChange={e => setUsername(e.target.value)}
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
