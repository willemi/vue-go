// MenuManagementView.vue 菜单管理页面
// 管理员可对菜单进行 CRUD 操作（创建、查看、编辑、删除）
// 普通用户无此页面访问权限（路由守卫已在 router/index.ts 中控制）
<template>
  <div class="menu-management">
    <h2>菜单管理</h2>

    <!-- 操作栏：新增按钮 -->
    <el-form inline class="search-form">
      <el-form-item>
        <el-button type="success" @click="openAddDialog">新增菜单</el-button>
      </el-form-item>
    </el-form>

    <!-- 菜单表格：展示所有菜单（含父级关系） -->
    <el-table :data="menus" border style="width: 100%">
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column prop="title" label="菜单名称" />
      <el-table-column prop="path" label="路径" />
      <el-table-column prop="icon" label="图标" />
      <!-- 父级列：根据 parent_id 查找对应的菜单名称 -->
      <el-table-column label="父级" width="140">
        <template #default="{ row }">
          <el-tag :type="row.parent_id === 0 ? 'success' : 'warning'" size="small">
            {{ getParentTitle(row.parent_id) }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="sort" label="排序" width="80" />
      <!-- 角色列：将逗号分隔的 role 字符串拆分为多个标签 -->
      <el-table-column label="角色" width="160">
        <template #default="{ row }">
          <template v-if="row.role">
            <el-tag
              v-for="r in row.role.split(',')"
              :key="r"
              :type="getRoleTagType(r.trim())"
              size="small"
              style="margin-right: 4px"
            >
              {{ getRoleLabel(r.trim()) }}
            </el-tag>
          </template>
          <span v-else>-</span>
        </template>
      </el-table-column>
      <!-- 隐藏列：控制菜单是否在侧边栏显示 -->
      <el-table-column label="隐藏" width="80">
        <template #default="{ row }">
          <el-tag :type="row.hidden ? 'danger' : 'success'">{{
            row.hidden ? '是' : '否'
          }}</el-tag>
        </template>
      </el-table-column>
      <!-- 操作列 -->
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
      <!-- 父级菜单选择：排除自身（防止将自己设为父级） -->
      <el-form-item label="父级菜单" prop="parent_id">
        <el-select v-model="form.parent_id" placeholder="请选择父级菜单" clearable style="width: 100%">
          <el-option label="无（顶级菜单）" :value="0" />
          <el-option
            v-for="menu in parentMenuOptions"
            :key="menu.id"
            :label="menu.title"
            :value="menu.id"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="排序" prop="sort">
        <el-select v-model="form.sort" placeholder="请选择排序" style="width: 100%">
          <el-option v-for="n in 10" :key="n - 1" :label="String(n - 1)" :value="n - 1" />
        </el-select>
      </el-form-item>
      <!-- 角色多选：多个角色用逗号分隔存储到后端 -->
      <el-form-item label="角色" prop="role">
        <el-select v-model="formRoles" multiple placeholder="请选择角色" style="width: 100%">
          <el-option label="管理员" value="admin" />
          <el-option label="普通用户" value="user" />
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
import { ref, computed, onMounted } from 'vue'
import { getMenuList, createMenu, updateMenu, deleteMenu } from '../api/user'

const menus = ref<any[]>([])

// 对话框状态
const dialogVisible = ref(false)
const isEdit = ref(false)
const formRef = ref()

// 表单数据
const form = ref({
  id: 0,
  title: '',
  path: '',
  icon: '',
  parent_id: 0,
  sort: 0,
  hidden: false,
  role: ''
})

// 角色多选：数组形式用于 el-select multiple，提交前转为逗号分隔字符串
const formRoles = ref<string[]>([])

// 表单验证规则：菜单名称和路径必填
const formRules = {
  title: [{ required: true, message: '请输入菜单名称', trigger: 'blur' }],
  path: [{ required: true, message: '请输入路径', trigger: 'blur' }]
}

// parentMenuOptions：父级菜单选项
// 只显示顶级菜单（parent_id === 0），且排除自身（编辑时不能将自己设为父级）
const parentMenuOptions = computed(() => {
  return menus.value.filter(m => m.parent_id === 0 && m.id !== form.value.id)
})

// getParentTitle 根据 parent_id 返回父级菜单名称
const getParentTitle = (parentId: number) => {
  if (!parentId) return '无'
  const parent = menus.value.find(m => m.id === parentId)
  return parent ? parent.title : `ID: ${parentId}`
}

// 角色标签映射：role 值到中文标签名
const getRoleLabel = (role: string) => {
  const map: Record<string, string> = {
    admin: '管理员',
    user: '普通用户',
    all: '所有用户'
  }
  return map[role] || role
}

// getRoleTagType 返回角色标签的颜色类型
const getRoleTagType = (role: string) => {
  const map: Record<string, string> = {
    admin: 'danger',
    user: 'info',
    all: 'success'
  }
  return map[role] || ''
}

// fetchMenus 获取菜单列表
const fetchMenus = async () => {
  try {
    const response = await getMenuList()
    menus.value = response.data
  } catch (error) {
    console.error('Failed to fetch menus:', error)
  }
}

// openAddDialog 打开新增对话框
const openAddDialog = () => {
  isEdit.value = false
  form.value = {
    id: 0,
    title: '',
    path: '',
    icon: '',
    parent_id: 0,
    sort: 0,
    hidden: false,
    role: ''
  }
  formRoles.value = []
  dialogVisible.value = true
}

// openEditDialog 打开编辑对话框
const openEditDialog = (row: any) => {
  isEdit.value = true
  form.value = { ...row }
  // 将逗号分隔的 role 字符串转为数组，用于 el-select multiple 显示
  formRoles.value = row.role ? row.role.split(',').map((r: string) => r.trim()) : []
  dialogVisible.value = true
}

// onSubmit 提交表单
const onSubmit = async () => {
  if (!formRef.value) return

  const valid = await formRef.value.validate()
  if (!valid) return

  try {
    // 将角色数组转为逗号分隔字符串
    form.value.role = formRoles.value.join(',')

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

// handleDelete 删除菜单
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