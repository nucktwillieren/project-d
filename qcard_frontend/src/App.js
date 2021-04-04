import { BrowserRouter, Route, Switch } from 'react-router-dom'
import Login from "./components/Auth/Login"
import Logout from "./components/Auth/Logout"
import Registration from "./components/Auth/Registration"
import { HomePage } from "./components/Home/Home"
import './App.css';
import TopNav from "./components/Basic/TopNav"
import { VerifyAuth, RedirectIfAuth } from "./VerifyAuth"
import { UserPage } from "./components/Dashboard/User"
import { useSelector } from "react-redux"

const App = () => {
  const initialToken = useSelector(state => state.token)
  const initialUser = useSelector(state => state.user)
  let user
  if (initialToken && initialUser) {
    user = initialUser
  }
  return (
    <div className="App">
      <header className="App-header">
      </header>
      <BrowserRouter>
        <TopNav user={user}></TopNav>
        <Switch>
          <Route path="/login">
            <RedirectIfAuth>
              <Login></Login>
            </RedirectIfAuth>
          </Route>
          <Route path="/register">
            <RedirectIfAuth>
              <Registration></Registration>
            </RedirectIfAuth>
          </Route>
          <Route path="/logout" component={Logout} />
          <Route path="/" component={HomePage} />
          <Route path="/user">
            <VerifyAuth>
              <UserPage />
            </VerifyAuth>
          </Route>
        </Switch>
      </BrowserRouter>
    </div>
  );
}

export default App;