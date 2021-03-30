export const parseLoginPayload = (data) => {
  localStorage.setItem("user", data.user)
  localStorage.setItem("token", data.access_token)
}

export const userLogin = async (credentials) => {
  return fetch('http://localhost:8080/api/v1/auth/login', {
    method: 'POST',
    mode: 'cors',
    cache: 'no-cache',
    credentials: 'same-origin',
    redirect: 'follow',
    referrerPolicy: 'no-referrer',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(credentials)
  })
    .then(data => data.json())
    .then(jsonData => parseLoginPayload(jsonData))
    .catch("err, please retry")
}

export const userLogout = () => {
  localStorage.removeItem("user")
  localStorage.removeItem("token")
}

export const getToken = () => {
  const token = localStorage.getItem('token');

  return token
}

export const getUser = () => {
  return localStorage.getItem('user')
}

export const userRegister = (payload) => {
  return fetch('http://localhost:8080/api/v1/auth/registration', {
    method: 'POST',
    mode: 'cors',
    cache: 'no-cache',
    credentials: 'same-origin',
    redirect: 'follow',
    referrerPolicy: 'no-referrer',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(payload)
  })
    .then(data => data.json())
    .then(jsonData => console.log(jsonData))
    .catch("err, please retry")
}