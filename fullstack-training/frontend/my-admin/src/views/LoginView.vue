// LoginView.vue 登录页面
// 提供用户名密码输入，提交后调用 /api/user/login
// 登录成功后保存 token 和 userInfo 到 Pinia store，然后跳转到 Dashboard
<template>
  <div class="login-container">
    <el-card class="login-card">
      <h2 class="login-title">后台管理系统</h2>
      <el-form
        :model="form"
        :rules="rules"
        ref="formRef"
        label-width="80px"
        class="login-form"
      >
        <el-form-item label="用户名" prop="username">
          <el-input v-model="form.username" autocomplete="off" />
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input
            v-model="form.password"
            type="password"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            @click="onSubmit"
            :loading="loading"
            >登录</el-button
          >
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { login } from '../api/user'
import { useUserStore } from '../stores/user'

const router = useRouter()
const userStore = useUserStore()

// form 表单数据
const form = reactive({
  username: '',
  password: ''
})

// rules 表单验证规则
const rules = reactive({
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' }
  ]
})

const formRef = ref()
const loading = ref(false)

// onSubmit 提交登录表单
const onSubmit = async () => {
  if (!formRef.value) return

  // 先进行表单验证
  const valid = await formRef.value.validate()
  if (!valid) return

  loading.value = true
  try {
    // 调用登录 API
    const response = await login(form.username, form.password)
    const { token, user } = response.data

    // 保存 token 和用户信息到 Pinia store（自动持久化到 localStorage）
    userStore.setToken(token)
    userStore.setUserInfo(user)

    // 跳转到 Dashboard
    router.push('/dashboard')
  } catch (error) {
    console.error('Login failed:', error)
    // TODO: 使用 Element Plus 的 ElMessage 显示错误提示
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f0f2f5;
}

.login-card {
  width: 360px;
  padding: 40px;
}

.login-title {
  text-align: center;
  margin-bottom: 30px;
  color: #333;
}

.login-form {
  width: 100%;
}
</style>