import { defineStore } from 'pinia'
import { ref } from 'vue'

interface UserInfo {
  id: number
  username: string
  role: string
}

export const useUserStore = defineStore('user', () => {
  const token = ref<string>(localStorage.getItem('token') || '')
  const userInfo = ref<UserInfo>(JSON.parse(localStorage.getItem('userInfo') || '{}'))

  function setToken(newToken: string) {
    token.value = newToken
    localStorage.setItem('token', newToken)
  }

  function setUserInfo(user: UserInfo) {
    userInfo.value = user
    localStorage.setItem('userInfo', JSON.stringify(user))
  }

  function logout() {
    token.value = ''
    userInfo.value = { id: 0, username: '', role: '' }
    localStorage.removeItem('token')
    localStorage.removeItem('userInfo')
  }

  return {
    token,
    userInfo,
    setToken,
    setUserInfo,
    logout
  }
})