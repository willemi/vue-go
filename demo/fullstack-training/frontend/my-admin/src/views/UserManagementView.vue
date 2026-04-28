<template>
  <div class="user-management">
    <h2>用户管理</h2>

    <!-- 搜索 -->
    <el-form inline class="search-form">
      <el-form-item label="用户名">
        <el-input v-model="searchUsername" clearable placeholder="请输入用户名" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="fetchUsers">搜索</el-button>
        <el-button type="success" @click="openAddDialog">新增用户</el-button>
      </el-form-item>
    </el-form>

    <!-- 用户表格 -->
    <el-table :data="users" border style="width: 100%">
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column prop="username" label="用户名" />
      <el-table-column prop="role" label="角色">
        <template #default="{ row }">
          <el-tag :type="row.role === 'admin' ? 'danger' : 'info'">
            {{ row.role === 'admin' ? '管理员' : '普通用户' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="created_at" label="创建时间">
        <template #default="{ row }">
          {{ formatTime(row.created_at) }}
        </template>
      </el-table-column>
      <el-table-column label="操作" width="180">
        <template #default="{ row }">
          <el-button type="primary" size="small" @click="openEditDialog(row)">编辑</el-button>
          <el-button type="danger" size="small" @click="handleDelete(row.id)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 分页 -->
    <el-pagination
      v-model:current-page="currentPage"
      v-model:page-size="pageSize"
      :total="total"
      :page-sizes="[10, 20, 50, 100]"
      layout="total, sizes, prev, pager, next, jumper"
      class="pagination"
      @current-change="fetchUsers"
      @size-change="fetchUsers"
    />

    <!-- 新增/编辑对话框 -->
    <el-dialog v-model="dialogVisible" :title="isEdit ? '编辑用户' : '新增用户'" width="400px">
      <el-form :model="form" :rules="formRules" ref="formRef" label-width="80px">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="form.username" />
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="form.password" type="password" />
        </el-form-item>
        <el-form-item label="角色" prop="role">
          <el-select v-model="form.role">
            <el-option label="管理员" value="admin" />
            <el-option label="普通用户" value="user" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="onSubmit">确认</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { getUserList, createUser, updateUser, deleteUser } from '../api/user'

const formatTime = (time: string) => {
  if (!time) return ''
  const date = new Date(time)
  const y = date.getFullYear()
  const m = String(date.getMonth() + 1).padStart(2, '0')
  const d = String(date.getDate()).padStart(2, '0')
  const h = String(date.getHours()).padStart(2, '0')
  const min = String(date.getMinutes()).padStart(2, '0')
  const s = String(date.getSeconds()).padStart(2, '0')
  return `${y}-${m}-${d} ${h}:${min}:${s}`
}

const users = ref<any[]>([])
const total = ref(0)
const currentPage = ref(1)
const pageSize = ref(10)
const searchUsername = ref('')

const dialogVisible = ref(false)
const isEdit = ref(false)
const formRef = ref()

const form = ref({
  id: 0,
  username: '',
  password: '',
  role: 'user'
})

const formRules = {
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }]
}

const fetchUsers = async () => {
  try {
    const response = await getUserList({
      username: searchUsername.value,
      page: currentPage.value,
      page_size: pageSize.value
    })
    const { users: list, total: count } = response.data
    users.value = list
    total.value = count
  } catch (error) {
    console.error('Failed to fetch users:', error)
  }
}

const openAddDialog = () => {
  isEdit.value = false
  form.value = { id: 0, username: '', password: '', role: 'user' }
  dialogVisible.value = true
}

const openEditDialog = (row: any) => {
  isEdit.value = true
  form.value = { id: row.id, username: row.username, password: '', role: row.role }
  dialogVisible.value = true
}

const onSubmit = async () => {
  if (!formRef.value) return

  const valid = await formRef.value.validate()
  if (!valid) return

  try {
    if (isEdit.value) {
      await updateUser(form.value)
    } else {
      await createUser(form.value)
    }
    dialogVisible.value = false
    fetchUsers()
  } catch (error) {
    console.error('Failed to save user:', error)
  }
}

const handleDelete = async (id: number) => {
  try {
    await deleteUser(id)
    fetchUsers()
  } catch (error) {
    console.error('Failed to delete user:', error)
  }
}

onMounted(fetchUsers)
</script>

<style scoped>
.user-management {
  padding: 20px;
}

.user-management h2 {
  text-align: left;
  margin-bottom:20px
}

.search-form {
  margin-bottom: 20px;
}

.pagination {
  margin-top: 20px;
  justify-content: flex-end;
}
</style>