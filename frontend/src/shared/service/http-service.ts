import axios from 'axios'

export const apiService = axios.create({
  baseURL: import.meta.env.HTTP_API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
})
