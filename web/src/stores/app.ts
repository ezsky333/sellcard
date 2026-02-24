// Utilities
import { defineStore } from 'pinia'

export const useAppStore = defineStore('app', {
  state: () => ({
    token: (typeof localStorage !== 'undefined' && localStorage.getItem('token')) || null,
    user: (typeof localStorage !== 'undefined' && (() => { try { const u = localStorage.getItem('user'); return u ? JSON.parse(u) : null } catch { return null } })()) || null,
    redirectPath: '',
  }),
  actions: {
    login(token: string, user: any) {
      this.token = token
      this.user = user
      if (typeof localStorage !== 'undefined') {
        localStorage.setItem('token', token)
        localStorage.setItem('user', JSON.stringify(user))
      }
    },
    logout() {
      this.token = null
      this.user = null
      if (typeof localStorage !== 'undefined') {
        localStorage.removeItem('token')
        localStorage.removeItem('user')
      }
    },
    setRedirect(path: string) {
      this.redirectPath = path
    },
  },
})
