import { BrowserRouter as Router, Route, Switch, Redirect } from 'react-router-dom';
import LoginPage from './pages/LoginPage.jsx';
import SignUpPage from './pages/SignUpPage.jsx';
import HomePage from './pages/HomePage.jsx';
import './App.css';
import { useState } from "react";
import SearchResultsPage from './pages/SearchResultsPage.jsx';
import ViewNotesPage from './pages/ViewNotesPage.jsx';
import AppContext from './components/AppContext';
import PersonalPage from './pages/PersonalPage.jsx';

function App() {
  const [authToken, setAuthToken] = useState("");
  const [filter, setFilter] = useState("");
  const [query, setQuery] = useState("");

  const setGlobalFilter = (newFilter) => {
    setFilter(newFilter);
  };

  const setGlobalQuery = (newQuery) => {
    setQuery(newQuery);
  };

  const searchParams = {
    filter,
    query,
    setGlobalFilter,
    setGlobalQuery
  }

  return (
    <AppContext.Provider value={searchParams}>
    <Router>
      <Switch>
        <Redirect from="/" to="/login" exact />
        <Route path="/login" component={(props) => <LoginPage setAuthToken={setAuthToken}/>} />
        <Route path="/signup" component={(props) => <SignUpPage setAuthToken={setAuthToken}/>} />
        <Route path="/home" component={(props) => <HomePage authToken={authToken}/>} />
        <Route path="/search" component={(props) => <SearchResultsPage/>}/>
        <Route path="/view" component={(props) => <ViewNotesPage/>}/>
        <Route path="/profile" component={(props) => <PersonalPage/>}/>
      </Switch>
    </Router>
  </AppContext.Provider>
  );
}

export default App;
