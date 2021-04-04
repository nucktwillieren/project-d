import { useState, useEffect } from 'react';
import "./Registration.css"
import { userRegister } from '../../api/auth'
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

const Registration = () => {
  const [username, setUsername] = useState();
  const [password, setPassword] = useState();
  const [password_second, setPasswordSecond] = useState();
  const [email, setEmail] = useState()
  const [token, setToken] = useState(getToken());
  const [warning, setWarning] = useState("");
  const [warningClass, setWarningClass] = useState("text-muted");
  let history = useHistory();

  const handleSubmit = async e => {
    e.preventDefault();
    await userRegister({
      username,
      password,
      password_second,
      email,
    });
    setToken(getToken());
  }

  useEffect(() => {
    if (token) {
      history.push("/")
    }
  })

  useEffect(() => {
    if (password !== password_second) {
      setWarning("Passwords do not match");
      setWarningClass("text-danger")
    } else if (password === password_second && password && password_second) {
      setWarning("OK!");      
      setWarningClass("text-success")
    }
  }, [password, password_second])
  
  return (
    <div className="container">
      <div className="login-container">
        <div id="output"></div>
        <div><h2>Sign Up</h2></div>
        <ColoredLine color="rgba(0, 0, 0, 0.4)"></ColoredLine>
        <div className="avatar"><img className="embedded-logo" src={process.env.PUBLIC_URL + '/logo192.png'} alt=""/></div>
        <div className="form-box">
          <form onSubmit={handleSubmit}>
            <input name="username" type="text" placeholder="Username" onChange={e => setUsername(e.target.value)} />
            <input name="email" type="email" placeholder="Email" onChange={e => setEmail(e.target.value)} />
            <input name="password" type="password" placeholder="Password" onChange={e => setPassword(e.target.value)} />
            <input name="password_second" type="password" placeholder="Password Again" onChange={e => setPasswordSecond(e.target.value)} />
            <span className={warningClass}>{warning}</span>
            <button className="btn btn-info btn-block login" type="submit">Sign Up</button>
          </form>
          <ColoredLine color="rgba(0, 0, 0, 0.4)"></ColoredLine>
          <div>Already have an account?</div>
          <div><a href="/login">Sign In Now!</a></div>
        </div>        
      </div>            
    </div>
  )
}

export default Registration;