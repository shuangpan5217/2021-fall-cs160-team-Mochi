import {
  BrowserRouter as Router,
  Route,
  Switch,
  Redirect,
} from "react-router-dom";
import LoginPage from "./pages/LoginPage.jsx";
import SignUpPage from "./pages/SignUpPage.jsx";
import HomePage from "./pages/HomePage.jsx";
import "./App.css";
import { useState } from "react";
import SearchResultsPage from "./pages/SearchResultsPage.jsx";
import ViewNotesPage from "./pages/ViewNotesPage.jsx";
import AppContext from "./components/AppContext";
import PersonalPage from "./pages/PersonalPage.jsx";
import GroupPage from "./pages/GroupPage.jsx";

function App() {
  const [filter, setFilter] = useState("");
  const [query, setQuery] = useState("");

  const setGlobalFilter = (newFilter) => {
    setFilter(newFilter);
  };

  const setGlobalQuery = (newQuery) => {
    setQuery(newQuery);
  };

  const globalVars = {
    filter,
    query,
    setGlobalFilter,
    setGlobalQuery,
  };

  return (
    <AppContext.Provider value={globalVars}>
      <Router>
        <Switch>
          <Redirect from="/" to="/login" exact />
          <Route path="/login" component={(props) => <LoginPage />} />
          <Route path="/signup" component={(props) => <SignUpPage />} />
          <Route path="/home" component={(props) => <HomePage />} />
          <Route path="/search" component={(props) => <SearchResultsPage />} />
          <Route path="/note/:noteId" component={(props) => <ViewNotesPage/>}/>
          <Route path="/my_notes" component={(props) => <PersonalPage />} />
          <Route path="/my_groups" component={(props) => <GroupPage />} />
        </Switch>
      </Router>
    </AppContext.Provider>
  );
}

export default App;
