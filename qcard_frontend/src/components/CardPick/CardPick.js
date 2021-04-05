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
  const [username, setUsername] = useState()
  const [email, setEmail] = useState()
  const [name, setName] = useState()
  const [gender, setGender] = useState()
  const [birthday, setBirthday] = useState()
  const [relationship, setRelationship] = useState()
  const [Interest, setInterest] = useState()
  const [club, setClub] = useState()
  const [favoriteCourse, setFavoriteCourse] = useState()
  const [favoriteCountry, setFavoriteCountry] = useState()
  const [trouble, setTrouble] = useState()
  const [exchange, setExchange] = useState()
  const [trying, setTrying] = useState()

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
      setUsername(resp.data.username)
      setEmail(resp.data.email)
      setName(resp.data.name)
      setGender(resp.data.gender)
      setBirthday(resp.data.birthday)
      setRelationship(resp.data.relationship)
      setInterest(resp.data.interest)
      setClub(resp.data.interest)
      setFavoriteCourse(resp.data.favoriteCourse)
      setFavoriteCountry(resp.data.favoriteCountry)
      setTrouble(resp.data.trouble)
      setExchange(resp.data.exchange)
      setTrying(resp.data.trying)
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
        <div><h2>{username}</h2></div>
        <ColoredLine color="rgba(0, 0, 0, 0.4)"></ColoredLine>
        <div className="avatar"><img className="embedded-logo" src={process.env.PUBLIC_URL + '/logo192.png'} alt="" /></div>
        <div className="form-box">
          <div>Username: {username} </div>
          <div>Email: {email} </div>
          <div>Name: {name}</div>
          <div>Gender: {gender}</div>
          <div>Birthday: {birthday}</div>
          <div>Relationship: {relationship}</div>
          <div>Interest:{Interest} </div>
          <div>Club: {club}</div>
          <div>Favorite Courses: {favoriteCourse}</div>
          <div>Favorite Country: {favoriteCountry}</div>
          <div>Trouble: {trouble}</div>
          <div>Exchange: {exchange}</div>
          <div>Trying: {trying}</div>
          <ColoredLine color="rgba(0, 0, 0, 0.4)"></ColoredLine>
          <button className="btn btn-info btn-block login" type="submit">Send a Friend Request</button>
        </div>
      </div>
    </div >
  )
}

export { CardPick };