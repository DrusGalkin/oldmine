import axios from 'axios'
import Cookies from 'js-cookie'

const BASE_URL = 'http://localhost:4000/api'
export const SESS_NAME = 'session_id'

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
    Cookies.set(SESS_NAME, response.data.session_id)

    return true
  } catch (error) {
    console.log(error)
    return false
  }
}

export const checkAuth = async () => {
    const response = await profile()
    if (response.status === 401) {
      return false
    } else if (response.status === 200) {
      return true
    }


}

export const profile = async () => {
  try{
    const response = await auth.get('/profile')
    return response
  }catch(error){
    console.log(error)
  }
}
