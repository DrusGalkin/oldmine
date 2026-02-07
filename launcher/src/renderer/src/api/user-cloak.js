import axios from 'axios'

const BASE_URL = 'http://localhost:4000/api'
export const CLOAK_FORM_FILE_NAME = "cloak"
export const cap = axios.create({
  baseURL: BASE_URL + '/cloaks',
  withCredentials: true
})

export const getCloak = (name) => {
  return `http://localhost:4000/api/cloaks/uploads/${name}.png`
}

export const getCloakPath = async (id) => {
  try {
    const response = await cap.get(`/${id}`)
    return response.data
  }catch(error) {
    console.log(error)
  }
}

export const saveCloak = async (file) => {
  try {
    const response = await cap.post(`/`, file)
    return response.data.message
  } catch (error) {
    throw error
  }
}



export const deleteCloak = async (id) => {
  try {
    const response = await cap.delete(`/${id}`)
    return response.data
  }catch(error) {
    throw error
  }
}