import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap/dist/js/bootstrap.js';
import './TopNav.css'
import { useState } from "react";
import { CreateCategoryCenteredModal, CreatePostCenteredModal } from "./LeftNav"
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import {
  faSignInAlt, faSdCard,
  faSignOutAlt, faUser,
  faUserPlus, faPlus,
  faIdCard, faPenAlt
} from '@fortawesome/free-solid-svg-icons'

const LogoutState = () => {
  return (
    <li className="nav-item">
      <a className="nav-link" href="/logout">Sign Out <FontAwesomeIcon icon={faSignOutAlt} /></a>
    </li>
  )
}

const LoginState = () => {
  return (
    <li className="nav-item">
      <a className="nav-link" href="/login">Sign In <FontAwesomeIcon icon={faSignInAlt} /></a>
    </li>
  )
}

const RegistrationState = () => {
  return (
    <li className="nav-item">
      <a className="nav-link" href="/register">Sign Up <FontAwesomeIcon icon={faUserPlus} /></a>
    </li>
  )
}

const UserState = (props) => {
  const user = props.user
  return (
    <li className="nav-item">
      <a className="nav-link" href="/user">{user} <FontAwesomeIcon icon={faUser} /></a>
    </li>
  )
}

const TopNav = (props) => {
  const [modalShow, setModalShow] = useState(false);
  const [postModalShow, setPostModalShow] = useState(false);
  const user = props.user

  return (
    <nav
      className="navbar navbar-expand-lg navbar-light sticky-top"
    >
      <div className="container">
        <a className="navbar-brand" href="/"><FontAwesomeIcon icon={faSdCard} style={{ color: "#0168B7" }} /> QCard</a>
        <button
          className="navbar-toggler"
          type="button"
          data-toggle="collapse"
          data-target="#navbarSupportedContent"
          aria-controls="navbarSupportedContent"
          aria-expanded="false"
          aria-label="Toggle navigation"
        >
          <span className="navbar-toggler-icon"></span>
        </button>
        <div className="collapse navbar-collapse" id="navbarSupportedContent">
          <ul className="navbar-nav ml-auto mb-2 mb-lg-0">
            <li className="nav-item">
              <a className="nav-link" onClick={() => setModalShow(true)}>New Category <FontAwesomeIcon icon={faPlus} /></a>
            </li>
            <li className="nav-item">
              <a className="nav-link" onClick={() => setPostModalShow(true)}>New Post<FontAwesomeIcon icon={faPenAlt} /></a>
            </li>
            <li className="nav-item">
              <a className="nav-link" href="/pick">Pick A Card <FontAwesomeIcon icon={faIdCard} /></a>
            </li>
            {
              user
                ? <><UserState user={user}></UserState><LogoutState></LogoutState></>
                : <><LoginState></LoginState><RegistrationState></RegistrationState></>
            }
          </ul>
        </div>
      </div>
      <CreateCategoryCenteredModal show={modalShow} onHide={() => setModalShow(false)} />
      <CreatePostCenteredModal show={postModalShow} onHide={() => setPostModalShow(false)} />
    </nav>
  )
}

export default TopNav;