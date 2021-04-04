import "./Login.css"
import { Redirect } from "react-router-dom";
import { useHistory } from "react-router-dom";
import { getUser, userLogout } from '../../api/auth'
import { LOGOUT } from "../../store"
import { useDispatch } from "react-redux"


const Logout = () => {
  const user = getUser();
  let history = useHistory();
  const dispatch = useDispatch();

  if (user) {
    userLogout()
    dispatch({ type: LOGOUT })
    history.push("/")
  }

  return (
    <Redirect to="/"></Redirect>
  )
}

export default Logout;