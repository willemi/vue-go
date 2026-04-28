<template>
  <div class="menu-management">
    <h2>菜单管理</h2>

    <!-- 搜索 -->
    <el-form inline class="search-form">
      <el-form-item>
        <el-button type="success" @click="openAddDialog">新增菜单</el-button>
      </el-form-item>
    </el-form>

    <!-- 菜单表格 -->
    <el-table :data="menus" border style="width: 100%">
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column prop="title" label="菜单名称" />
      <el-table-column prop="path" label="路径" />
      <el-table-column prop="icon" label="图标" />
      <el-table-column prop="parent_id" label="父级ID" width="100" />
      <el-table-column prop="sort" label="排序" width="80" />
      <el-table-column prop="role" label="角色" width="120" />
      <el-table-column label="隐藏" width="80">
        <template #default="{ row }">
          <el-tag :type="row.hidden ? 'danger' : 'success'">{{
            row.hidden ? '是' : '否'
          }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="180">
        <template #default="{ row }">
          <el-button type="primary" size="small" @click="openEditDialog(row)">编辑</el-button>
          <el-button type="danger" size="small" @click="handleDelete(row.id)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>

  <!-- 新增/编辑对话框 -->
  <el-dialog v-model="dialogVisible" :title="isEdit ? '编辑菜单' : '新增菜单'" width="400px">
    <el-form :model="form" :rules="formRules" ref="formRef" label-width="100px">
      <el-form-item label="菜单名称" prop="title">
        <el-input v-model="form.title" />
      </el-form-item>
      <el-form-item label="路径" prop="path">
        <el-input v-model="form.path" />
      </el-form-item>
      <el-form-item label="图标" prop="icon">
        <el-input v-model="form.icon" />
      </el-form-item>
      <el-form-item label="父级ID" prop="parent_id">
        <el-input-number v-model="form.parent_id" :min="0" />
      </el-form-item>
      <el-form-item label="排序" prop="sort">
        <el-input-number v-model="form.sort" :min="0" />
      </el-form-item>
      <el-form-item label="角色" prop="role">
        <el-select v-model="form.role">
          <el-option label="管理员" value="admin" />
          <el-option label="普通用户" value="user" />
          <el-option label="所有用户" value="all" />
        </el-select>
      </el-form-item>
      <el-form-item label="隐藏" prop="hidden">
        <el-switch v-model="form.hidden" />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="dialogVisible = false">取消</el-button>
      <el-button type="primary" @click="onSubmit">确认</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { getMenuList, createMenu, updateMenu, deleteMenu } from '../api/user'

const menus = ref<any[]>([])

const dialogVisible = ref(false)
const isEdit = ref(false)
const formRef = ref()

const form = ref({
  title: '',
  path: '',
  icon: '',
  parent_id: 0,
  sort: 0,
  hidden: false,
  role: 'user'
})

const formRules = {
  title: [{ required: true, message: '请输入菜单名称', trigger: 'blur' }],
  path: [{ required: true, message: '请输入路径', trigger: 'blur' }]
}

const fetchMenus = async () => {
  try {
    const response = await getMenuList()
    menus.value = response.data
  } catch (error) {
    console.error('Failed to fetch menus:', error)
  }
}

const openAddDialog = () => {
  isEdit.value = false
  form.value = {
    title: '',
    path: '',
    icon: '',
    parent_id: 0,
    sort: 0,
    hidden: false,
    role: 'user'
  }
  dialogVisible.value = true
}

const openEditDialog = (row: any) => {
  isEdit.value = true
  form.value = { ...row }
  dialogVisible.value = true
}

const onSubmit = async () => {
  if (!formRef.value) return

  const valid = await formRef.value.validate()
  if (!valid) return

  try {
    if (isEdit.value) {
      await updateMenu(form.value)
    } else {
      await createMenu(form.value)
    }
    dialogVisible.value = false
    fetchMenus()
  } catch (error) {
    console.error('Failed to save menu:', error)
  }
}

const handleDelete = async (id: number) => {
  try {
    await deleteMenu(id)
    fetchMenus()
  } catch (error) {
    console.error('Failed to delete menu:', error)
  }
}

onMounted(fetchMenus)
</script>

<style scoped>
.menu-management {
  padding: 20px;
}


.menu-management h2 {
  text-align: left;
  margin-bottom:20px
}
.search-form {
  margin-bottom: 20px;
  text-align: right;
}
</style>