import { BrowserRouter as Router, Route, Switch, Redirect } from 'react-router-dom';
import LoginPage from './pages/LoginPage.jsx';
import SignUpPage from './pages/SignUpPage.jsx';
import HomePage from './pages/HomePage.jsx';
import './App.css';
import { useState } from "react";

function App() {
  const [authToken, setAuthToken] = useState("");
  return (
    <>
    <Router>
      <Switch>
        <Redirect from="/" to="/login" exact />
        <Route path="/login" component={(props) => <LoginPage setAuthToken={setAuthToken}/>} />
        <Route path="/signup" component={(props) => <SignUpPage setAuthToken={setAuthToken}/>} />
        <Route path="/home" component={(props) => <HomePage authToken={authToken}/>} />
      </Switch>
    </Router>
  </>
  );
}

export default App;
