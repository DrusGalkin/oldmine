import axios from 'axios'
import error from 'eslint-plugin-react/lib/util/error'

const BASE_URL = 'http://localhost:4000/api'
export const SKIN_FORM_FILE_NAME = "skin"
export const skins = axios.create({
  baseURL: BASE_URL + '/skins',
  withCredentials: true
})

export const getSkin = (name) => {
  return `http://localhost:4000/api/skins/uploads/${name}.png`
}

export const getSkinPath = async (id) => {
  try {
    const response = await skins.get(`/${id}`)
    return response.data
  }catch(error) {
    console.log(error)
  }
}

export const saveSkin = async (file) => {
  try {
    const response = await skins.post(`/`, file, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })

    return response.data.message
  } catch (error) {
    throw error
  }
}


export const putSkin = async (file) => {
  try {
    const response = await skins.put(`/`, file, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })

    return response.data.message
  } catch (e) {
    console.error(e)
  }
}


export const deleteSkin = async (id) => {
  try {
    const response = await skins.delete(`/${id}`)
    return response.data
  }catch(error) {
    throw error
  }
}