import "./Login.css"
import { Redirect } from "react-router-dom";
import { useHistory } from "react-router-dom";
import { getUser, userLogout } from '../../api/auth'


const Logout = () => {
  const user = getUser();
  let history = useHistory();

  if (user) {
    userLogout()
    history.push("/")
  }

  return (
    <Redirect to="/"></Redirect>
  )
}

export default Logout;