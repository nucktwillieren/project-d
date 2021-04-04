import { createStore, compose } from 'redux'

const initialState = {
  token: null,
  user: null
}

export const SET_TOKEN = "SET_TOKEN"
export const SET_USER = "SET_USER"
export const LOGOUT = "LOGOUT"

const rootReducer = (state = initialState, action) => {
  switch (action.type) {
    case SET_TOKEN:
      localStorage.setItem("token", action.payload)
      return {
        ...state,
        token: action.payload
      }
    case SET_USER:
      localStorage.setItem("user", action.payload)
      return {
        ...state,
        user: action.payload
      }
    case LOGOUT:
      localStorage.removeItem("user")
      localStorage.removeItem("token")
      return {
        initialState
      }
    default:
      return state
  }
}

let preloadedState
const persistedToken = localStorage.getItem("token")
const persistedUser = localStorage.getItem("user")
if (persistedToken && persistedUser) {
  preloadedState = {
    token: persistedToken,
    user: persistedUser
  }
}

const store = createStore(
  rootReducer,
  preloadedState,
  compose(
    window.__REDUX_DEVTOOLS_EXTENSION__
      ? window.__REDUX_DEVTOOLS_EXTENSION__()
      : (f) => f
  )
)

export default store