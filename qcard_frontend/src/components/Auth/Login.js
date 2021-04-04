import { useState } from 'react';
import "./Login.css"
import { useHistory } from "react-router-dom";
import { useDispatch } from "react-redux"
import { SET_TOKEN, SET_USER } from "../../store"
import axios from 'axios'

const ColoredLine = ({ color }) => (
  <hr
    style={{
      color: color,
      backgroundColor: color,
      height: 0.5,
      width: 200,
    }}
  />
);

const Login = () => {
  let history = useHistory();
  const dispatch = useDispatch()

  const [username, setUsername] = useState();
  const [password, setPassword] = useState();

  const handleSubmit = async e => {
    e.preventDefault();
    try {
      const config = {
        headers: {
          'Content-Type': 'application/json'
        },
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        redirect: 'follow',
        referrerPolicy: 'no-referrer',
      }
      const data = { username, password }
      const resp = await axios.post(
        'http://localhost:8080/api/v1/auth/login',
        data, config
      )

      dispatch({
        type: SET_TOKEN,
        payload: resp.data.access_token
      })

      dispatch({
        type: SET_USER,
        payload: resp.data.user
      })

      history.push("/")
    } catch (error) {
      console.log(error)
    }
  }

  return (
    <div className="container">
      <div className="login-container">
        <div id="output"></div>
        <div><h2>Sign In</h2></div>
        <ColoredLine color="rgba(0, 0, 0, 0.4)"></ColoredLine>
        <div className="avatar"><img className="embedded-logo" src={process.env.PUBLIC_URL + '/logo192.png'} alt="" /></div>
        <div className="form-box">
          <form onSubmit={handleSubmit}>
            <input type="text" placeholder="Username" onChange={e => setUsername(e.target.value)} />
            <input type="password" placeholder="Password" onChange={e => setPassword(e.target.value)} />
            <button className="btn btn-info btn-block login" type="submit">Sign In</button>
          </form>
          <ColoredLine color="rgba(0, 0, 0, 0.4)"></ColoredLine>
          <div>No account yet?</div>
          <div><a href="/register">Sign Up Now!</a></div>
        </div>
      </div>
    </div>
  )
}

export default Login;