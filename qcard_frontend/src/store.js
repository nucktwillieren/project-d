import { createStore, compose } from 'redux'

const initialState = {
  token: null,
  user: null
}

export const SET_TOKEN = "SET_TOKEN"
export const SET_USER = "SET_USER"

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