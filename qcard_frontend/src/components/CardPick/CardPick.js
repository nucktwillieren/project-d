import "./CardPick.css"
import { useSelector } from "react-redux"
import { useState, useEffect } from "react";
import axios from "axios"

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

const CardPick = () => {
  const initialToken = useSelector(state => state.token)
  const initialUser = useSelector(state => state.user)
  const [card, setCard] = useState({})
  const config = {
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${initialToken}`
    },
    mode: 'cors',
    cache: 'no-cache',
    credentials: 'same-origin',
    redirect: 'follow',
    referrerPolicy: 'no-referrer',
  }

  useEffect(() => {
    const initialFetch = async () => {
      const resp = await axios.get(
        `http://localhost:8080/api/v1/pair/${initialUser}`,
        config
      )
      console.log(resp.data)
      setCard({ ...card, ...resp.data })
      console.log(card)
    }
    const initialCreate = async () => {
      try {
        const resp = await axios.post(
          `http://localhost:8080/api/v1/pair/${initialUser}`,
          config
        )
      } catch (error) {
      }
    }
    initialCreate()
    initialFetch()
  }, [])

  return (
    <div className="container" >
      <div className="profile-container">
        <div><h2>{card.name}</h2></div>
        <ColoredLine color="rgba(0, 0, 0, 0.4)"></ColoredLine>
        <div className="avatar"><img className="embedded-logo" src={process.env.PUBLIC_URL + '/logo192.png'} alt="" /></div>
        <div className="form-box">
          <form>
            <input type="text" placeholder="Username" />
            <input type="password" placeholder="Password" />
            <button className="btn btn-info btn-block login" type="submit">Sign In</button>
          </form>
          <ColoredLine color="rgba(0, 0, 0, 0.4)"></ColoredLine>
          <div>No account yet?</div>
          <div><a href="/register">Sign Up Now!</a></div>
        </div>
      </div>
    </div >
  )
}

export { CardPick };