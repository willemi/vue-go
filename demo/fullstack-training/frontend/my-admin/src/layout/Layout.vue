<template>
  <div class="layout">
    <el-container>
      <!-- Sidebar -->
      <el-aside width="200px" class="sidebar">
        <el-menu
          :default-active="activeMenu"
          class="el-menu-vertical-demo"
          background-color="#304156"
          text-color="#fff"
          active-text-color="#ffd04b"
          router
          :unique-opened="true"
        >
          <el-sub-menu v-for="item in menus" :key="item.path" :index="item.path">
            <template #title>
              <el-icon><component :is="item.icon" /></el-icon>
              <span>{{ item.title }}</span>
            </template>
            <el-menu-item
              v-for="child in item.children"
              :key="child.path"
              :index="child.path"
            >
              <el-icon><component :is="child.icon" /></el-icon>
              <span>{{ child.title }}</span>
            </el-menu-item>
          </el-sub-menu>
        </el-menu>
      </el-aside>

      <!-- Main content -->
      <el-container>
        <!-- Header -->
        <el-header height="60px" class="header">
          <div class="header-content">
            <h1>后台管理系统</h1>
            <el-dropdown @command="handleCommand">
              <span class="el-dropdown-link">
                {{ userStore.userInfo.username }}
                <el-icon class="el-icon--right">
                  <arrow-down />
                </el-icon>
              </span>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item command="logout">退出登录</el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
          </div>
        </el-header>

        <!-- Content -->
        <el-main>
          <router-view />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRoute } from 'vue-router'
import { useUserStore } from '../stores/user'
import { getMenuList } from '../api/user'

const route = useRoute()
const userStore = useUserStore()

const activeMenu = computed(() => route.path)
const menus = ref<any[]>([])

// Fetch menu list on mount
const fetchMenus = async () => {
  try {
    const response = await getMenuList()
    menus.value = response.data.data
  } catch (error) {
    console.error('Failed to fetch menus:', error)
  }
}

fetchMenus()

const handleCommand = (command: string) => {
  if (command === 'logout') {
    userStore.logout()
  }
}
</script>

<style scoped>
.layout {
  height: 100vh;
}

.sidebar {
  background-color: #304156;
  color: #fff;
}

.header {
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 100%;
  padding: 0 20px;
}

.header-content h1 {
  margin: 0;
  color: #333;
}

.el-dropdown-link {
  cursor: pointer;
  color: #333;
}
</style>