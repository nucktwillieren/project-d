import "./Login.css"
import { Redirect } from "react-router-dom";
import { LOGOUT } from "../../store"
import { useDispatch } from "react-redux"


const Logout = () => {
  const dispatch = useDispatch();

  dispatch({ type: LOGOUT })

  return (
    <Redirect to="/"></Redirect>
  )
}

export default Logout;