<template>
  <div class="layout">
    <el-container style="height:100%">
      <!-- 侧边栏 -->
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
          <template v-for="item in menus" :key="item.path">
            <!-- 没有子菜单的一级菜单 -->
            <el-menu-item v-if="!item.children || item.children.length === 0" :index="item.path">
              <el-icon><component :is="item.icon" /></el-icon>
              <span>{{ item.title }}</span>
            </el-menu-item>
            <!-- 有子菜单的一级菜单 -->
            <el-sub-menu v-else :index="item.path">
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
          </template>
        </el-menu>
      </el-aside>

      <!-- 主内容区 -->
      <el-container>
        <!-- 顶部导航 -->
        <el-header height="60px" class="header">
          <div class="header-content">
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

        <!-- 内容区 -->
        <el-main>
          <router-view />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { ArrowDown } from '@element-plus/icons-vue'
import { useUserStore } from '../stores/user'
import { getMenuTree } from '../api/user'

interface MenuNode {
  id: number
  title: string
  path: string
  icon: string
  parent_id: number
  sort: number
  hidden: boolean
  role: string
  children: MenuNode[]
}

const route = useRoute()
const userStore = useUserStore()

const activeMenu = computed(() => route.path)
const menus = ref<MenuNode[]>([])

// 挂载时获取菜单树
const fetchMenus = async () => {
  try {
    const response = await getMenuTree()
    menus.value = response.data
  } catch (error) {
    console.error('Failed to fetch menu tree:', error)
  }
}

onMounted(() => {
  fetchMenus()
})

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
  justify-content: flex-end;
  align-items: center;
  height: 100%;
  padding: 0 20px;
}
.el-main>div{
  height:100%;
  box-sizing: border-box;
}

.header-content h1 {
  margin:0;
  color: #333;
}

.el-dropdown-link {
  cursor: pointer;
  color: #333;
}
</style>
