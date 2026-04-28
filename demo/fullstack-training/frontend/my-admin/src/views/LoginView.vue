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

const form = reactive({
  username: '',
  password: ''
})

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

const onSubmit = async () => {
  if (!formRef.value) return
debugger
  const valid = await formRef.value.validate()
  if (!valid) return

  loading.value = true
  try {
    const response = await login(form.username, form.password)
    const { token, user } = response.data.data

    userStore.setToken(token)
    userStore.setUserInfo(user)

    router.push('/dashboard')
  } catch (error) {
    console.error('Login failed:', error)
    // 处理错误
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