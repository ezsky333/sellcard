<template>
  <v-app-bar color="primary" dark flat elevation="1">
    <v-app-bar-nav-icon @click="drawer = !drawer" />
    <v-app-bar-title>
      <span class="font-weight-bold">SellCard 管理系统</span>
    </v-app-bar-title>

    <v-spacer />

    <!-- User Info & Logout -->
    <v-menu offset-y>
      <template v-slot:activator="{ props }">
        <v-btn icon v-bind="props" class="mr-2" :title="user?.username">
          <v-icon>mdi-account-circle</v-icon>
        </v-btn>
      </template>

      <v-list>
        <v-list-item>
          <template v-slot:prepend>
            <v-icon>mdi-account</v-icon>
          </template>
          <v-list-item-title>{{ user?.username }}</v-list-item-title>
          <v-list-item-subtitle>{{ user?.role }}</v-list-item-subtitle>
        </v-list-item>

        <v-divider class="my-2" />

        <v-list-item @click="handleLogout">
          <template v-slot:prepend>
            <v-icon color="error">mdi-logout</v-icon>
          </template>
          <v-list-item-title class="text-error">退出登录</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>
  </v-app-bar>

  <!-- Navigation Drawer -->
  <v-navigation-drawer v-model="drawer" mobile-breakpoint="md" elevation="0">
    <v-list>
      <v-list-subheader>导航菜单</v-list-subheader>

      <v-list-item
        v-for="item in menuItems"
        :key="item.id"
        :active="activeMenu === item.id"
        @click="handleMenuSelect(item.id)"
      >
        <template v-slot:prepend>
          <v-icon>{{ item.icon }}</v-icon>
        </template>
        <v-list-item-title>{{ item.label }}</v-list-item-title>
      </v-list-item>
    </v-list>
  </v-navigation-drawer>

  <!-- Main Content -->
  <v-main>
    <v-container fluid class="pa-6">
      <div v-if="activeMenu === 'dashboard'">
        <v-card class="elevation-1 rounded-lg">
          <v-card-title class="text-h6">仪表板</v-card-title>
          <v-divider />
          <v-card-text class="pa-6">
            <v-row>
              <v-col cols="12" md="4">
                <v-card color="blue-lighten-4" class="elevation-0 rounded-lg">
                  <v-card-text class="text-center">
                    <div class="text-h3 font-weight-bold text-blue">0</div>
                    <div class="text-grey">总用户数</div>
                  </v-card-text>
                </v-card>
              </v-col>
              <v-col cols="12" md="4">
                <v-card color="green-lighten-4" class="elevation-0 rounded-lg">
                  <v-card-text class="text-center">
                    <div class="text-h3 font-weight-bold text-green">0</div>
                    <div class="text-grey">活跃用户</div>
                  </v-card-text>
                </v-card>
              </v-col>
              <v-col cols="12" md="4">
                <v-card color="orange-lighten-4" class="elevation-0 rounded-lg">
                  <v-card-text class="text-center">
                    <div class="text-h3 font-weight-bold text-orange">0</div>
                    <div class="text-grey">总订单数</div>
                  </v-card-text>
                </v-card>
              </v-col>
            </v-row>

            <v-divider class="my-6" />

            <div class="text-h6 mb-4">欢迎, {{ user?.username }}!</div>
            <p class="text-grey">
              您正在使用 SellCard 管理系统。选择左侧菜单项来管理您的数据。
            </p>
          </v-card-text>
        </v-card>
      </div>

      <div v-if="activeMenu === 'users'">
        <v-card class="elevation-1 rounded-lg">
          <v-card-title class="text-h6">用户管理</v-card-title>
          <v-divider />
          <v-card-text class="pa-6">
            <v-row class="mb-4">
              <v-col>
                <v-btn color="primary" prepend-icon="mdi-plus">
                  新增用户
                </v-btn>
              </v-col>
              <v-col cols="12" md="4">
                <v-text-field
                  v-model="searchUser"
                  label="搜索用户"
                  prepend-inner-icon="mdi-magnify"
                  outlined
                  dense
                />
              </v-col>
            </v-row>

            <v-data-table
              :headers="[
                { title: '用户名', value: 'username' },
                { title: '邮箱', value: 'email' },
                { title: '角色', value: 'role' },
                { title: '状态', value: 'status' },
                { title: '操作', value: 'actions' },
              ]"
              :items="users"
              :search="searchUser"
            >
              <template v-slot:item.status="{ item }">
                <v-chip
                  :color="item.status === 'active' ? 'green' : 'grey'"
                  size="small"
                >
                  {{ item.status === 'active' ? '活跃' : '已禁用' }}
                </v-chip>
              </template>
              <template v-slot:item.actions>
                <v-btn icon size="small" variant="text" color="primary">
                  <v-icon>mdi-pencil</v-icon>
                </v-btn>
                <v-btn icon size="small" variant="text" color="error">
                  <v-icon>mdi-delete</v-icon>
                </v-btn>
              </template>
            </v-data-table>
          </v-card-text>
        </v-card>
      </div>

      <div v-if="activeMenu === 'settings'">
        <v-card class="elevation-1 rounded-lg">
          <v-card-title class="text-h6">系统设置</v-card-title>
          <v-divider />
          <v-card-text class="pa-6">
            <v-row>
              <v-col cols="12" md="6">
                <v-text-field
                  label="系统名称"
                  value="SellCard"
                  outlined
                  dense
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  label="系统版本"
                  value="1.0.0"
                  outlined
                  dense
                />
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="12" md="6">
                <v-select
                  label="语言"
                  :items="['中文', '英文']"
                  outlined
                  dense
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-select
                  label="主题"
                  :items="['亮色', '暗色']"
                  outlined
                  dense
                />
              </v-col>
            </v-row>
            <v-btn color="primary">
              保存设置
            </v-btn>
          </v-card-text>
        </v-card>
      </div>
    </v-container>
  </v-main>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAppStore } from '@/stores/app'

const router = useRouter()
const store = useAppStore()

const activeMenu = ref('dashboard')
const searchUser = ref('')
const drawer = ref(true)

const user = computed(() => store.user)

const menuItems = [
  { id: 'dashboard', label: '仪表板', icon: 'mdi-home' },
  { id: 'users', label: '用户管理', icon: 'mdi-account-multiple' },
  { id: 'settings', label: '系统设置', icon: 'mdi-cog' },
]

// Mock user data
const users = ref([
  { username: 'admin', email: 'admin@example.com', role: '管理员', status: 'active' },
  { username: 'user1', email: 'user1@example.com', role: '用户', status: 'active' },
  { username: 'user2', email: 'user2@example.com', role: '用户', status: 'active' },
])

const handleMenuSelect = (menuId: string) => {
  activeMenu.value = menuId
  // 在移动设备上关闭抽屉
  if (window.innerWidth < 960) {
    drawer.value = false
  }
}

const handleLogout = () => {
  store.logout()
  router.push('/login')
}
</script>

<style scoped>
:deep(.rounded-lg) {
  border-radius: 8px !important;
}

:deep(.v-app-bar) {
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

:deep(.v-navigation-drawer) {
  border-radius: 0 !important;
}
</style>
