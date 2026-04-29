// DashboardView.vue 仪表盘页面
// 展示系统概览数据（当前为静态数据，可扩展为从后端获取实时统计）
<template>
  <div class="dashboard">
    <h1>欢迎来到后台管理系统</h1>

    <!-- 统计卡片：使用 el-row 和 el-col 实现响应式布局 -->
    <el-row :gutter="20" class="stats-row">
      <el-col :span="12">
        <el-card class="stat-card">
          <h3>用户总数</h3>
          <!-- TODO: 从后端 API 获取真实数据 -->
          <p class="stat-value">128</p>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card class="stat-card">
          <h3>菜单总数</h3>
          <p class="stat-value">24</p>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { useUserStore } from '../stores/user'

const userStore = useUserStore()
</script>

<style scoped>
.dashboard {
  padding: 20px;
}

.stats-row {
  margin: 20px 0;
}

.stat-card {
  text-align: center;
  height: 100px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.stat-value {
  font-size: 24px;
  font-weight: bold;
  color: #409EFF;
  margin: 10px 0 0;
}
</style>