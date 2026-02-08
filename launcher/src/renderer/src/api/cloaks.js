import axios from 'axios'
import error from 'eslint-plugin-react/lib/util/error'

const BASE_URL = 'http://localhost:4000/api'
export const CLOAK_FORM_FILE_NAME = "cloak"
export const cloaks = axios.create({
    baseURL: BASE_URL + '/cloaks',
    withCredentials: true
})

export const getCloak = (name) => {
    return `http://localhost:4000/api/cloaks/uploads/${name}.png`
}

export const getCloakPath = async (id) => {
    try {
        const response = await cloaks.get(`/${id}`)
        return response.data
    }catch(error) {
        console.log(error)
    }
}

export const saveCloak = async (file) => {
    try {
        const response = await cloaks.post(`/`, file)
        return response.data.message
    } catch (error) {
        throw error
    }
}

export const deleteCloak = async (id) => {
    try {
        const response = await cloaks.delete(`/${id}`)
        return response.data
    }catch(error) {
        throw error
    }
}