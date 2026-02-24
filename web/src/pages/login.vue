<template>
  <div class="login-wrapper">
    <v-container class="login-container" fluid>
      <v-row class="h-100" align="center" justify="center">
        <v-col cols="12" sm="10" md="8" lg="5" xl="4" class="d-flex justify-center login-col">
          <v-card elevation="8" class="login-card pa-8">
            <!-- Logo/Title -->
            <div class="text-center mb-6">
              <h1 class="text-h4 font-weight-bold mb-2 text-blue-darken-2">
                SellCard
              </h1>
              <p class="text-grey">管理系统</p>
            </div>

            <!-- Login Form -->
            <v-form @submit.prevent="handleLogin" v-model="valid">
              <!-- Username Field -->
              <v-text-field
                v-model="form.username"
                label="用户名"
                type="text"
                outlined
                dense
                class="mb-4"
                :rules="usernameRules"
                prepend-inner-icon="mdi-account"
                @keyup.enter="handleLogin"
              />

              <!-- Password Field -->
              <v-text-field
                v-model="form.password"
                label="密码"
                type="password"
                outlined
                dense
                class="mb-4"
                :rules="passwordRules"
                prepend-inner-icon="mdi-lock"
                @keyup.enter="handleLogin"
              />

              <!-- Turnstile Verification -->
              <div v-if="turnstileEnabled" class="mb-4">
                <div ref="turnstileContainer" class="turnstile-container" />
              </div>

              <!-- Turnstile Loading Message -->
              <v-alert
                v-if="!turnstileEnabled"
                type="info"
                density="compact"
                class="mb-4"
              >
                Turnstile 验证服务暂不可用，请稍后重试
              </v-alert>

              <!-- Error Message -->
              <v-alert
                v-if="errorMessage"
                type="error"
                dismissible
                class="mb-4"
                @click:close="errorMessage = ''"
              >
                {{ errorMessage }}
              </v-alert>

              <!-- Login Button -->
              <v-btn
                type="submit"
                block
                size="large"
                color="primary"
                :loading="loading"
                :disabled="loading || !valid || (turnstileEnabled && !turnstileToken)"
              >
                登录
              </v-btn>
            </v-form>

            <!-- Footer -->
            <v-divider class="my-4" />
            <p class="text-center text-caption text-grey">
              演示账号: admin / 123456
            </p>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAppStore } from '@/stores/app'
import { login, TURNSTILE_SITE_KEY } from '@/services/auth'

const router = useRouter()
const store = useAppStore()

const form = ref({
  username: '',
  password: '',
})

const valid = ref(false)
const loading = ref(false)
const errorMessage = ref('')
const turnstileEnabled = ref(true)
const turnstileToken = ref('')
const turnstileContainer = ref<HTMLElement | null>(null)

const usernameRules = [
  (v: string) => !!v || '用户名不能为空',
]

const passwordRules = [
  (v: string) => !!v || '密码不能为空',
]

// 初始化 Turnstile
const initTurnstile = async () => {
  // 等待容器挂载
  await new Promise(resolve => setTimeout(resolve, 100))

  if (!window.turnstile) {
    console.warn('Turnstile script not loaded')
    turnstileEnabled.value = false
    return
  }

  if (!turnstileContainer.value) {
    console.warn('Turnstile container not found')
    turnstileEnabled.value = false
    return
  }

  try {
    window.turnstile.render(turnstileContainer.value, {
      sitekey: TURNSTILE_SITE_KEY,
      theme: 'light',
      callback: (token: string) => {
        turnstileToken.value = token
      },
      'error-callback': () => {
        console.error('Turnstile error')
        errorMessage.value = 'Turnstile 验证加载失败'
      },
    })
  } catch (error) {
    console.error('Failed to render Turnstile:', error)
    turnstileEnabled.value = false
  }
}

const handleLogin = async () => {
  if (!valid.value) return

  // 如果启用了 Turnstile，检查 token
  if (turnstileEnabled.value && !turnstileToken.value) {
    errorMessage.value = '请完成人机验证'
    return
  }

  loading.value = true
  errorMessage.value = ''

  try {
    const response = await login({
      username: form.value.username,
      password: form.value.password,
      turnstile_token: turnstileToken.value || undefined,
    })

    // Store token and user info
    store.login(response.token, response.user)

    // Redirect to backend or home
    const redirectPath = store.redirectPath || '/backend'
    store.setRedirect('')
    router.push(redirectPath)
  } catch (error) {
    errorMessage.value = error instanceof Error ? error.message : '登录失败，请检查账号和密码'
    // Reset Turnstile on error
    if (turnstileEnabled.value && window.turnstile) {
      window.turnstile.reset()
      turnstileToken.value = ''
    }
  } finally {
    loading.value = false
  }
}

// 生命周期：加载 Turnstile 脚本
onMounted(() => {
  const loadTurnstile = () => {
    if (window.turnstile) {
      initTurnstile()
    } else {
      const script = document.createElement('script')
      script.src = 'https://challenges.cloudflare.com/turnstile/v0/api.js'
      script.async = true
      script.defer = true
      script.onload = () => {
        initTurnstile()
      }
      script.onerror = () => {
        console.error('Failed to load Turnstile script')
        turnstileEnabled.value = false
      }
      document.head.appendChild(script)
    }
  }

  // 延迟加载以确保 DOM 已挂载
  setTimeout(() => {
    loadTurnstile()
  }, 50)
})
</script>

<script lang="ts">
declare global {
  interface Window {
    turnstile: {
      render: (
        container: HTMLElement,
        options: {
          sitekey: string
          theme: 'light' | 'dark'
          callback?: (token: string) => void
          'error-callback'?: () => void
          'expired-callback'?: () => void
          'timeout-callback'?: () => void
        },
      ) => string
      reset: (widgetId?: string) => void
      remove: (widgetId?: string) => void
      getResponse: (widgetId?: string) => string | undefined
    }
  }
}
export {}
</script>

<style scoped>
.login-wrapper {
  width: 100%;
  height: 100vh;
  overflow: hidden;
}

.login-container {
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 20px;
}

.login-card {
  width: 100%;
  padding: 20px;
  border-radius: 8px !important;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3) !important;
}
.login-col{
  max-width: 450px;
}
:deep(.v-text-field) {
  font-size: 14px;
}

:deep(.v-card-title) {
  padding: 20px;
}

.turnstile-container {
  display: flex;
  justify-content: center;
  margin: 8px 0;
}

:deep(.cf-turnstile) {
  display: flex;
  justify-content: center;
  margin: 0 auto;
}
</style>
