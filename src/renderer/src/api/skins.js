import axios from 'axios'

const BASE_URL = 'http://localhost:4000/api'

export const skins = axios.create({
  baseURL: BASE_URL + '/skins',
  withCredentials: true
})

export const getSkin = async (id) => {
  try {
    const response = await skins.get(`/${id}`)
    return response.data
  }catch(error) {
    console.log(error)
  }
}
