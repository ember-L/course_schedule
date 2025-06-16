<template>
  <div class="page-container fade-in">
    <!-- 页面头部 -->
    <div class="page-header">
      <div class="header-content">
        <div class="header-left">
          <h1 class="page-title">教师管理</h1>
          <p class="page-subtitle">管理系统中的所有教师信息</p>
        </div>
        <div class="header-right">
          <a-button type="primary" @click="showAddModal" class="add-button">
            <icon-plus />
            添加教师
          </a-button>
        </div>
      </div>
    </div>

    <!-- 统计卡片 -->
    <div class="card-grid">
      <div class="statistic-card">
        <div class="statistic-content">
          <div class="statistic-icon">
            <icon-user />
          </div>
          <div class="statistic-info">
            <div class="statistic-value">{{ teachers.length }}</div>
            <div class="statistic-label">总教师数</div>
          </div>
        </div>
      </div>
      
      <div class="statistic-card">
        <div class="statistic-content">
          <div class="statistic-icon">
            <icon-star />
          </div>
          <div class="statistic-info">
            <div class="statistic-value">{{ professorCount }}</div>
            <div class="statistic-label">教授人数</div>
          </div>
        </div>
      </div>
    </div>

    <!-- 教师列表 -->
    <a-card class="teachers-card" title="教师列表">
      <template #extra>
        <div class="card-actions">
          <a-input-search
            v-model="searchKeyword"
            placeholder="搜索教师..."
            style="width: 250px"
            @search="handleSearch"
          />
        </div>
      </template>
      
      <a-table
        :loading="loading"
        :columns="columns"
        :data="filteredTeachers"
        :pagination="{
          showTotal: true,
          pageSize: 10,
          showJumper: true,
          showPageSize: true
        }"
        class="teachers-table"
      >
        <template #operations="{ record }">
          <div class="action-buttons">
            <a-button 
              type="text" 
              size="small"
              @click="showEditModal(record)"
              class="action-button edit"
            >
              <icon-edit />
              编辑
            </a-button>
            <a-popconfirm
              content="确定要删除这个教师吗？"
              @ok="deleteTeacher(record.id)"
            >
              <a-button 
                type="text" 
                size="small"
                status="danger"
                class="action-button delete"
              >
                <icon-delete />
                删除
              </a-button>
            </a-popconfirm>
          </div>
        </template>
        
        <template #title="{ record }">
          <a-tag :color="getTitleColor(record.title)">
            {{ record.title }}
          </a-tag>
        </template>
        
        <template #email="{ record }">
          <div class="email-cell">
            <icon-email />
            <a :href="`mailto:${record.email}`">{{ record.email }}</a>
          </div>
        </template>
      </a-table>
    </a-card>
    
    <!-- 添加/编辑教师模态框 -->
    <a-modal
      v-model:visible="modalVisible"
      :title="isEdit ? '编辑教师' : '添加教师'"
      @ok="handleModalOk"
      @cancel="modalVisible = false"
      :width="600"
      class="teacher-modal"
    >
      <a-form :model="form" ref="formRef" :rules="rules" layout="vertical">
        <a-row :gutter="16">
          <a-col :span="12">
            <a-form-item field="name" label="姓名">
              <a-input 
                v-model="form.name" 
                placeholder="请输入姓名"
                :maxLength="20"
              />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item field="title" label="职称">
              <a-select 
                v-model="form.title" 
                placeholder="请选择职称"
                allow-clear
              >
                <a-option value="教授">教授</a-option>
                <a-option value="副教授">副教授</a-option>
                <a-option value="讲师">讲师</a-option>
                <a-option value="助教">助教</a-option>
              </a-select>
            </a-form-item>
          </a-col>
        </a-row>
        
        <a-form-item field="email" label="邮箱">
          <a-input 
            v-model="form.email" 
            placeholder="请输入邮箱"
            type="email"
          />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, computed } from 'vue';
import { Message } from '@arco-design/web-vue';
import { teacherApi } from '../api/modules';

const loading = ref(false);
const teachers = ref([]);
const modalVisible = ref(false);
const isEdit = ref(false);
const formRef = ref(null);
const searchKeyword = ref('');

const form = reactive({
  id: null,
  name: '',
  title: '',
  email: ''
});

const rules = {
  name: [
    { required: true, message: '请输入姓名' },
    { minLength: 2, message: '姓名至少2个字符' }
  ],
  title: [
    { required: true, message: '请选择职称' }
  ],
  email: [
    { required: true, message: '请输入邮箱' },
    { type: 'email', message: '请输入有效的邮箱地址' }
  ]
};

const columns = [
  {
    title: '姓名',
    dataIndex: 'name',
    width: 120,
    fixed: 'left'
  },
  {
    title: '职称',
    dataIndex: 'title',
    slotName: 'title',
    width: 120
  },
  {
    title: '邮箱',
    dataIndex: 'email',
    slotName: 'email',
    width: 200
  },
  {
    title: '操作',
    slotName: 'operations',
    width: 150,
    fixed: 'right'
  }
];

// 计算属性
const professorCount = computed(() => {
  return teachers.value.filter(teacher => teacher.title === '教授').length;
});

const filteredTeachers = computed(() => {
  if (!searchKeyword.value) return teachers.value;
  const keyword = searchKeyword.value.toLowerCase();
  return teachers.value.filter(teacher => 
    teacher.name.toLowerCase().includes(keyword) ||
    teacher.title.toLowerCase().includes(keyword) ||
    teacher.email.toLowerCase().includes(keyword)
  );
});

onMounted(async () => {
  await fetchTeachers();
});

const fetchTeachers = async () => {
  loading.value = true;
  try {
    const data = await teacherApi.getAll();
    teachers.value = data;
  } catch (error) {
    Message.error('获取教师列表失败');
    console.error(error);
  } finally {
    loading.value = false;
  }
};

const showAddModal = () => {
  isEdit.value = false;
  form.id = null;
  form.name = '';
  form.title = '';
  form.email = '';
  modalVisible.value = true;
};

const showEditModal = (record) => {
  isEdit.value = true;
  form.id = record.id;
  form.name = record.name;
  form.title = record.title;
  form.email = record.email;
  modalVisible.value = true;
};

const handleModalOk = async () => {
  formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        if (isEdit.value) {
          await teacherApi.update(form.id, form);
          Message.success('更新教师成功');
        } else {
          await teacherApi.create(form);
          Message.success('添加教师成功');
        }
        modalVisible.value = false;
        await fetchTeachers();
      } catch (error) {
        Message.error(isEdit.value ? '更新教师失败' : '添加教师失败');
        console.error(error);
      }
    }
  });
};

const deleteTeacher = async (id) => {
  try {
    await teacherApi.delete(id);
    Message.success('删除教师成功');
    await fetchTeachers();
  } catch (error) {
    Message.error('删除教师失败');
    console.error(error);
  }
};

const handleSearch = () => {
  // 搜索逻辑已在 computed 中处理
};

const getTitleColor = (title) => {
  switch (title) {
    case '教授': return 'red';
    case '副教授': return 'orange';
    case '讲师': return 'blue';
    case '助教': return 'green';
    default: return 'gray';
  }
};
</script>

<style scoped>
.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-left {
  flex: 1;
}

.header-right {
  display: flex;
  gap: var(--spacing-sm);
}

.add-button {
  display: flex;
  align-items: center;
  gap: var(--spacing-xs);
  font-weight: 500;
}

.teachers-card {
  margin-top: var(--spacing-xl);
}

.card-actions {
  display: flex;
  gap: var(--spacing-sm);
  align-items: center;
}

.teachers-table {
  margin-top: var(--spacing-md);
}

.action-buttons {
  display: flex;
  gap: var(--spacing-xs);
}

.action-button {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 12px;
  padding: 4px 8px;
  border-radius: var(--border-radius-small);
  transition: all 0.2s ease;
}

.action-button.edit:hover {
  background-color: var(--primary-light);
  color: var(--primary-color);
}

.action-button.delete:hover {
  background-color: var(--danger-light);
  color: var(--danger-color);
}

.email-cell {
  display: flex;
  align-items: center;
  gap: var(--spacing-xs);
  color: var(--primary-color);
}

.email-cell a {
  color: var(--primary-color);
  text-decoration: none;
}

.email-cell a:hover {
  text-decoration: underline;
}

.teacher-modal .arco-form-item {
  margin-bottom: var(--spacing-lg);
}

/* 响应式设计 */
@media (max-width: 768px) {
  .header-content {
    flex-direction: column;
    align-items: flex-start;
    gap: var(--spacing-md);
  }
  
  .header-right {
    width: 100%;
  }
  
  .add-button {
    width: 100%;
    justify-content: center;
  }
  
  .card-actions {
    flex-direction: column;
    align-items: stretch;
  }
  
  .action-buttons {
    flex-direction: column;
  }
}
</style>