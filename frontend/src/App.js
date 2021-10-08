import { BrowserRouter as Router, Route, Switch, Redirect } from 'react-router-dom';
import LoginPage from './pages/LoginPage.jsx';
import SignUpPage from './pages/SignUpPage.jsx';
import HomePage from './pages/HomePage.jsx';
import './App.css';

function App() {
  return (
    <>
    <Router>
      <Switch>
        <Redirect from="/" to="/login" exact />
        <Route path="/login" component={(props) => <LoginPage />} />
        <Route path="/signup" component={(props) => <SignUpPage />} />
        <Route path="/home" component={(props) => <HomePage />} />
      </Switch>
    </Router>
  </>
  );
}

export default App;
