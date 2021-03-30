import { useEffect } from 'react';
import "./Login.css"
import { useHistory } from "react-router-dom";
import { getUser, userLogout } from '../../api/auth'


const Logout = () => {
  const user = getUser();
  let history = useHistory();

  useEffect(() => {
    if (user) {
      userLogout()
      history.push("/")
    }
  })
  return (
    <></>
  )
}

export default Logout;