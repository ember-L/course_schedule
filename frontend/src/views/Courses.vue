<template>
  <div class="page-container fade-in">
    <!-- 页面头部 -->
    <div class="page-header">
      <div class="header-content">
        <div class="header-left">
          <h1 class="page-title">课程管理</h1>
          <p class="page-subtitle">管理系统中的所有课程信息</p>
        </div>
        <div class="header-right">
          <a-button type="primary" @click="showAddModal" class="add-button">
            <icon-plus />
            添加课程
          </a-button>
        </div>
      </div>
    </div>

    <!-- 统计卡片 -->
    <div class="card-grid">
      <div class="statistic-card">
        <div class="statistic-content">
          <div class="statistic-icon">
            <icon-book />
          </div>
          <div class="statistic-info">
            <div class="statistic-value">{{ courses.length }}</div>
            <div class="statistic-label">总课程数</div>
          </div>
        </div>
      </div>
      
      <div class="statistic-card">
        <div class="statistic-content">
          <div class="statistic-icon">
            <icon-clock-circle />
          </div>
          <div class="statistic-info">
            <div class="statistic-value">{{ totalCredits }}</div>
            <div class="statistic-label">总学分</div>
          </div>
        </div>
      </div>
    </div>

    <!-- 课程列表 -->
    <a-card class="courses-card" title="课程列表">
      <template #extra>
        <div class="card-actions">
          <a-input-search
            v-model="searchKeyword"
            placeholder="搜索课程..."
            style="width: 250px"
            @search="handleSearch"
          />
        </div>
      </template>
      
      <a-table
        :loading="loading"
        :columns="columns"
        :data="filteredCourses"
        :pagination="{
          showTotal: true,
          pageSize: 10,
          showJumper: true,
          showPageSize: true
        }"
        class="courses-table"
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
              content="确定要删除这个课程吗？"
              @ok="deleteCourse(record.id)"
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
        
        <template #credits="{ record }">
          <a-tag :color="getCreditsColor(record.credits)">
            {{ record.credits }} 学分
          </a-tag>
        </template>
        
        <template #description="{ record }">
          <div class="description-cell">
            <span class="description-text">{{ record.description || '暂无描述' }}</span>
          </div>
        </template>
      </a-table>
    </a-card>
    
    <!-- 添加/编辑课程模态框 -->
    <a-modal
      v-model:visible="modalVisible"
      :title="isEdit ? '编辑课程' : '添加课程'"
      @ok="handleModalOk"
      @cancel="modalVisible = false"
      :width="600"
      class="course-modal"
    >
      <a-form :model="form" ref="formRef" :rules="rules" layout="vertical">
        <a-row :gutter="16">
          <a-col :span="12">
            <a-form-item field="code" label="课程代码">
              <a-input 
                v-model="form.code" 
                placeholder="请输入课程代码"
                :maxLength="20"
              />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item field="credits" label="学分">
              <a-input-number 
                v-model="form.credits" 
                placeholder="请输入学分"
                :min="0"
                :max="20"
                style="width: 100%"
              />
            </a-form-item>
          </a-col>
        </a-row>
        
        <a-form-item field="name" label="课程名称">
          <a-input 
            v-model="form.name" 
            placeholder="请输入课程名称"
            :maxLength="50"
          />
        </a-form-item>
        
        <a-form-item field="description" label="课程描述">
          <a-textarea 
            v-model="form.description" 
            placeholder="请输入课程描述"
            :maxLength="200"
            :auto-size="{ minRows: 3, maxRows: 5 }"
            show-word-limit
          />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, computed } from 'vue';
import { Message } from '@arco-design/web-vue';
import { courseApi } from '../api/modules';

const loading = ref(false);
const courses = ref([]);
const modalVisible = ref(false);
const isEdit = ref(false);
const formRef = ref(null);
const searchKeyword = ref('');

const form = reactive({
  id: null,
  code: '',
  name: '',
  credits: 0,
  description: ''
});

const rules = {
  code: [
    { required: true, message: '请输入课程代码' },
    { minLength: 2, message: '课程代码至少2个字符' }
  ],
  name: [
    { required: true, message: '请输入课程名称' },
    { minLength: 2, message: '课程名称至少2个字符' }
  ],
  credits: [
    { required: true, type: 'number', message: '请输入学分' },
    { type: 'number', min: 0, max: 20, message: '学分范围0-20' }
  ]
};

const columns = [
  {
    title: '课程代码',
    dataIndex: 'code',
    width: 120,
    fixed: 'left'
  },
  {
    title: '课程名称',
    dataIndex: 'name',
    width: 200
  },
  {
    title: '学分',
    dataIndex: 'credits',
    slotName: 'credits',
    width: 100
  },
  {
    title: '课程描述',
    dataIndex: 'description',
    slotName: 'description',
    ellipsis: true
  },
  {
    title: '操作',
    slotName: 'operations',
    width: 150,
    fixed: 'right'
  }
];

// 计算属性
const totalCredits = computed(() => {
  return courses.value.reduce((sum, course) => sum + (course.credits || 0), 0);
});

const filteredCourses = computed(() => {
  if (!searchKeyword.value) return courses.value;
  const keyword = searchKeyword.value.toLowerCase();
  return courses.value.filter(course => 
    course.name.toLowerCase().includes(keyword) ||
    course.code.toLowerCase().includes(keyword) ||
    (course.description && course.description.toLowerCase().includes(keyword))
  );
});

onMounted(async () => {
  await fetchCourses();
});

const fetchCourses = async () => {
  loading.value = true;
  try {
    const data = await courseApi.getAll();
    courses.value = data;
  } catch (error) {
    Message.error('获取课程列表失败');
    console.error(error);
  } finally {
    loading.value = false;
  }
};

const showAddModal = () => {
  isEdit.value = false;
  form.id = null;
  form.code = '';
  form.name = '';
  form.credits = 0;
  form.description = '';
  modalVisible.value = true;
};

const showEditModal = (record) => {
  isEdit.value = true;
  form.id = record.id;
  form.code = record.code;
  form.name = record.name;
  form.credits = record.credits;
  form.description = record.description;
  modalVisible.value = true;
};

const handleModalOk = async () => {
  formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        if (isEdit.value) {
          await courseApi.update(form.id, form);
          Message.success('更新课程成功');
        } else {
          await courseApi.create(form);
          Message.success('添加课程成功');
        }
        modalVisible.value = false;
        await fetchCourses();
      } catch (error) {
        Message.error(isEdit.value ? '更新课程失败' : '添加课程失败');
        console.error(error);
      }
    }
  });
};

const deleteCourse = async (id) => {
  try {
    await courseApi.delete(id);
    Message.success('删除课程成功');
    await fetchCourses();
  } catch (error) {
    Message.error('删除课程失败');
    console.error(error);
  }
};

const handleSearch = () => {
  // 搜索逻辑已在 computed 中处理
};

const getCreditsColor = (credits) => {
  if (credits <= 2) return 'green';
  if (credits <= 4) return 'blue';
  if (credits <= 6) return 'orange';
  return 'red';
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

.courses-card {
  margin-top: var(--spacing-xl);
}

.card-actions {
  display: flex;
  gap: var(--spacing-sm);
  align-items: center;
}

.courses-table {
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

.description-cell {
  max-width: 300px;
}

.description-text {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
  color: var(--text-secondary);
  font-size: 14px;
}

.course-modal .arco-form-item {
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
  
  .description-cell {
    max-width: 200px;
  }
}
</style>