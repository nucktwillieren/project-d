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
import { LeftNav } from "./components/Basic/LeftNav"
import { CardPick } from "./components/CardPick/CardPick"

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
        <TopNav user={user}>
        </TopNav>

        <LeftNav></LeftNav>
        <div id="emptyBlock" style={{ height: "8vh" }}></div>
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
          <VerifyAuth>
            <Route path="/logout" component={Logout} />
            <Route path="/user">
              <UserPage />
            </Route>
            <Route path="/pick" component={CardPick}></Route>
            <Route path="/" component={HomePage} />
          </VerifyAuth>
        </Switch>
      </BrowserRouter>
    </div >
  );
}

export default App;