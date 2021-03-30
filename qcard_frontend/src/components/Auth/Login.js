import { useState, useEffect } from 'react';
import "./Login.css"
import { userLogin } from '../../api/auth'
import { useHistory } from "react-router-dom";
import { getToken } from '../../api/auth'

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
  const [username, setUsername] = useState();
  const [password, setPassword] = useState();
  const [token, setToken] = useState(getToken())
  let history = useHistory();

  const handleSubmit = async e => {
    e.target.reset();
    await userLogin({
      username,
      password
    });
    setToken(getToken());
  }

  useEffect(() => {
    if (token) {
      history.push("/")
    }
  })
  
  return (
    <div className="container">
      <div className="login-container">
        <div id="output"></div>
        <div><h2>Sign In</h2></div>
        <ColoredLine color="rgba(0, 0, 0, 0.4)"></ColoredLine>
        <div className="avatar"><img className="embedded-logo" src={process.env.PUBLIC_URL + '/logo192.png'} alt=""/></div>
        <div className="form-box">
          <form onSubmit={handleSubmit}>
            <input name="username" type="text" placeholder="Username" onChange={e => setUsername(e.target.value)} />
            <input name="password" type="password" placeholder="Password" onChange={e => setPassword(e.target.value)} />
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