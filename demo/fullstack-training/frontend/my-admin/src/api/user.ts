import axios from 'axios'
import { useUserStore } from '../stores/user'

const api = axios.create({
  baseURL: 'http://localhost:8080/api',
  timeout: 10000
})

// Request interceptor
api.interceptors.request.use(
  config => {
    const userStore = useUserStore()
    if (userStore.token) {
      config.headers.Authorization = `Bearer ${userStore.token}`
    }
    return config
  },
  error => {
    return Promise.reject(error)
  }
)

// Response interceptor
api.interceptors.response.use(
  response => {
    return response
  },
  error => {
    if (error.response?.status === 401) {
      const userStore = useUserStore()
      userStore.logout()
      window.location.href = '/login'
    }
    return Promise.reject(error)
  }
)

// User API
export const login = (username: string, password: string) => {
  return api.post('/user/login', { username, password })
}

export const getUserList = (params: { username?: string; page?: number; page_size?: number }) => {
  return api.get('/user/list', { params })
}

export const createUser = (data: { username: string; password: string; role?: string }) => {
  return api.post('/user/add', data)
}

export const updateUser = (data: { id: number; username: string; password?: string; role?: string }) => {
  return api.put('/user/edit', data)
}

export const deleteUser = (id: number) => {
  return api.delete(`/user/delete/${id}`)
}

// Menu API
export const getMenuList = () => {
  return api.get('/menu/list')
}

export const createMenu = (data: { title: string; path: string; icon?: string; parent_id?: number; sort?: number; hidden?: boolean; role?: string }) => {
  return api.post('/menu/add', data)
}

export const updateMenu = (data: { id: number; title: string; path: string; icon?: string; parent_id?: number; sort?: number; hidden?: boolean; role?: string }) => {
  return api.put('/menu/edit', data)
}

export const deleteMenu = (id: number) => {
  return api.delete(`/menu/delete/${id}`)
}