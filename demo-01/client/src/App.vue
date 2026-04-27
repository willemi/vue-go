<template>
  <div style="max-width: 600px; margin: 40px auto">
    <h1>Vue3 + Golang 演示</h1>

    <div style="margin: 20px 0">
      <h3>接口返回：</h3>
      <p>{{ msg }}</p>
    </div>

    <div style="margin: 20px 0">
      <h3>用户信息：</h3>
      <div>昵称：{{ user.nickname }}</div>
      <div>年龄：{{ user.age }}</div>
    </div>

    <div style="margin: 40px 0">
      <h3>登录测试：</h3>
      <div>
        <input v-model="loginForm.username" placeholder="用户名" />
        <input v-model="loginForm.password" placeholder="密码" type="password" />
        <button @click="handleLogin">登录</button>
      </div>
      <p style="color: green" v-if="loginResult">{{ loginResult }}</p>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import request from './utils/request.js'

const msg = ref('')
const user = ref({})
const loginResult = ref('')

const loginForm = ref({
  username: 'admin',
  password: '123456'
})

// 页面加载获取数据
onMounted(async () => {
  const hello = await request.get('/api/hello')
  msg.value = hello.message

  const info = await request.get('/api/user/info')
  user.value = info.data
})

// 登录
const handleLogin = async () => {
  try {
    const res = await request.post('/api/login', loginForm.value)
    if (res.code === 0) {
      loginResult.value = '登录成功！token：' + res.token
    } else {
      loginResult.value = '登录失败：' + res.msg
    }
  } catch (err) {
    loginResult.value = '请求失败'
  }
}
</script>