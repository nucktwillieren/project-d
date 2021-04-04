import { BrowserRouter, Route, Switch } from 'react-router-dom'
import Login from "../Auth/Login"
import Logout from "../Auth/Logout"
import Registration from "../Auth/Registration"
import './App.css';
import TopNav from "../Basic/TopNav"
import { getToken, getUser } from '../../api/auth'
import { useState, useEffect } from 'react';

const App = () => {
  const [user, setUser] = useState(getUser());
  const [token, setToken] = useState(getToken());
  
  useEffect(() => {
    setToken(getToken());
    setUser(getUser());
  })

  return (
    <div className="App"  className="">
      <header className="App-header">
      </header>
      <BrowserRouter>
        <TopNav user={user}></TopNav>
        <Switch>
          <Route path="/login">
            <Login></Login>
          </Route>
          <Route path="/logout">
            <Logout></Logout>
          </Route>
          <Route path="/register">
            <Registration></Registration>
          </Route>
        </Switch>
      </BrowserRouter>
    </div>
  );
}

export default App;