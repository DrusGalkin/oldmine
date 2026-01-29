import axios from 'axios'

const BASE_URL = 'http://localhost:4000/api'
const SESS_NAME = 'sess_id'

export const auth = axios.create({
  baseURL: BASE_URL + '/auth',
  withCredentials: true
})

export const login = async (data) => {
  const { email, password } = data

  try {
    const response = await auth.post('/login', {
      email: email,
      password: password
    })

    console.log(response)

    localStorage.setItem(SESS_NAME, response.data.session_id)
    return 'Успешная авторизация'
  } catch (error) {
    console.log(error)
    return 'Невалидные данные'
  }
}
