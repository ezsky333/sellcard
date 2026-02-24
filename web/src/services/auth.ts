// API base URL
const API_URL = import.meta.env.VITE_API_URL || ''

// Turnstile configuration
export const TURNSTILE_SITE_KEY = import.meta.env.VITE_TURNSTILE_SITE_KEY || '0x4AAAAAAAv8X5OMVd5YPvYA'

export interface LoginRequest {
  username: string
  password: string
  turnstile_token?: string
}

export interface LoginResponse {
  token: string
  user: {
    username: string
    role: string
  }
}

/**
 * Login with username and password and Turnstile token
 */
export async function login(credentials: LoginRequest): Promise<LoginResponse> {
  const response = await fetch(`${API_URL}/api/v1/auth/login`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(credentials),
  })

  if (!response.ok) {
    const error = await response.json().catch(() => ({ message: 'Login failed' }))
    throw new Error(error.message || 'Login failed')
  }

  return response.json()
}
