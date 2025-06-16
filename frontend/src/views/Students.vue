<template>
  <div class="page-container fade-in">
    <!-- 页面头部 -->
    <div class="page-header">
      <div class="header-content">
        <div class="header-left">
          <h1 class="page-title">学生管理</h1>
          <p class="page-subtitle">管理系统中的所有学生信息</p>
        </div>
        <div class="header-right">
          <a-button type="primary" @click="showAddModal" class="add-button">
            <icon-plus />
            添加学生
          </a-button>
        </div>
      </div>
    </div>

    <!-- 统计卡片 -->
    <div class="card-grid">
      <div class="statistic-card">
        <div class="statistic-content">
          <div class="statistic-icon">
            <icon-user-group />
          </div>
          <div class="statistic-info">
            <div class="statistic-value">{{ students.length }}</div>
            <div class="statistic-label">总学生数</div>
          </div>
        </div>
      </div>
      
      <div class="statistic-card">
        <div class="statistic-content">
          <div class="statistic-icon">
            <icon-calendar />
          </div>
          <div class="statistic-info">
            <div class="statistic-value">{{ gradeCount }}</div>
            <div class="statistic-label">年级数量</div>
          </div>
        </div>
      </div>
      
      <div class="statistic-card">
        <div class="statistic-content">
          <div class="statistic-icon">
            <icon-book />
          </div>
          <div class="statistic-info">
            <div class="statistic-value">{{ majorCount }}</div>
            <div class="statistic-label">专业数量</div>
          </div>
        </div>
      </div>
    </div>

    <!-- 学生列表 -->
    <a-card class="students-card" title="学生列表">
      <template #extra>
        <div class="card-actions">
          <a-input-search
            v-model="searchKeyword"
            placeholder="搜索学生..."
            style="width: 250px"
            @search="handleSearch"
          />
        </div>
      </template>
      
      <a-table
        :loading="loading"
        :columns="columns"
        :data="filteredStudents"
        :pagination="{
          showTotal: true,
          pageSize: 10,
          showJumper: true,
          showPageSize: true
        }"
        class="students-table"
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
            <a-button 
              type="text" 
              size="small"
              @click="viewEnrollments(record)"
              class="action-button view"
            >
              <icon-eye />
              选课记录
            </a-button>
            <a-popconfirm
              content="确定要删除这个学生吗？"
              @ok="deleteStudent(record.id)"
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
        
        <template #grade="{ record }">
          <a-tag :color="getGradeColor(record.grade)">
            {{ record.grade }}
          </a-tag>
        </template>
        
        <template #major="{ record }">
          <a-tag :color="getMajorColor(record.major)">
            {{ record.major }}
          </a-tag>
        </template>
        
        <template #studentId="{ record }">
          <div class="student-id-cell">
            <icon-idcard />
            <span>{{ record.studentId }}</span>
          </div>
        </template>
      </a-table>
    </a-card>
    
    <!-- 添加/编辑学生模态框 -->
    <a-modal
      v-model:visible="modalVisible"
      :title="isEdit ? '编辑学生' : '添加学生'"
      @ok="handleModalOk"
      @cancel="modalVisible = false"
      :width="600"
      class="student-modal"
    >
      <a-form :model="form" ref="formRef" :rules="rules" layout="vertical">
        <a-row :gutter="16">
          <a-col :span="12">
            <a-form-item field="studentId" label="学号">
              <a-input 
                v-model="form.studentId" 
                placeholder="请输入学号"
                :maxLength="20"
              />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item field="name" label="姓名">
              <a-input 
                v-model="form.name" 
                placeholder="请输入姓名"
                :maxLength="20"
              />
            </a-form-item>
          </a-col>
        </a-row>
        
        <a-row :gutter="16">
          <a-col :span="12">
            <a-form-item field="grade" label="年级">
              <a-select 
                v-model="form.grade" 
                placeholder="请选择年级"
                allow-clear
              >
                <a-option value="2020级">2020级</a-option>
                <a-option value="2021级">2021级</a-option>
                <a-option value="2022级">2022级</a-option>
                <a-option value="2023级">2023级</a-option>
                <a-option value="2024级">2024级</a-option>
              </a-select>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item field="major" label="专业">
              <a-select 
                v-model="form.major" 
                placeholder="请选择专业"
                allow-clear
              >
                <a-option value="计算机科学与技术">计算机科学与技术</a-option>
                <a-option value="软件工程">软件工程</a-option>
                <a-option value="信息安全">信息安全</a-option>
                <a-option value="人工智能">人工智能</a-option>
                <a-option value="数据科学">数据科学</a-option>
                <a-option value="网络工程">网络工程</a-option>
                <a-option value="电子信息工程">电子信息工程</a-option>
                <a-option value="通信工程">通信工程</a-option>
              </a-select>
            </a-form-item>
          </a-col>
        </a-row>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, computed } from 'vue';
import { Message } from '@arco-design/web-vue';
import { useRouter } from 'vue-router';
import { studentApi } from '../api/modules';

const router = useRouter();
const loading = ref(false);
const students = ref([]);
const modalVisible = ref(false);
const isEdit = ref(false);
const formRef = ref(null);
const searchKeyword = ref('');

const form = reactive({
  id: null,
  studentId: '',
  name: '',
  grade: '',
  major: ''
});

const rules = {
  studentId: [
    { required: true, message: '请输入学号' },
    { minLength: 8, message: '学号至少8个字符' }
  ],
  name: [
    { required: true, message: '请输入姓名' },
    { minLength: 2, message: '姓名至少2个字符' }
  ],
  grade: [
    { required: true, message: '请选择年级' }
  ],
  major: [
    { required: true, message: '请选择专业' }
  ]
};

const columns = [
  {
    title: '学号',
    dataIndex: 'studentId',
    slotName: 'studentId',
    width: 150,
    fixed: 'left'
  },
  {
    title: '姓名',
    dataIndex: 'name',
    width: 120
  },
  {
    title: '年级',
    dataIndex: 'grade',
    slotName: 'grade',
    width: 100
  },
  {
    title: '专业',
    dataIndex: 'major',
    slotName: 'major',
    width: 200
  },
  {
    title: '操作',
    slotName: 'operations',
    width: 200,
    fixed: 'right'
  }
];

// 计算属性
const gradeCount = computed(() => {
  const grades = new Set(students.value.map(student => student.grade));
  return grades.size;
});

const majorCount = computed(() => {
  const majors = new Set(students.value.map(student => student.major));
  return majors.size;
});

const filteredStudents = computed(() => {
  if (!searchKeyword.value) return students.value;
  const keyword = searchKeyword.value.toLowerCase();
  return students.value.filter(student => 
    student.studentId.toLowerCase().includes(keyword) ||
    student.name.toLowerCase().includes(keyword) ||
    student.grade.toLowerCase().includes(keyword) ||
    student.major.toLowerCase().includes(keyword)
  );
});

onMounted(async () => {
  await fetchStudents();
});

const fetchStudents = async () => {
  loading.value = true;
  try {
    const data = await studentApi.getAll();
    students.value = data;
  } catch (error) {
    Message.error('获取学生列表失败');
    console.error(error);
  } finally {
    loading.value = false;
  }
};

const showAddModal = () => {
  isEdit.value = false;
  form.id = null;
  form.studentId = '';
  form.name = '';
  form.grade = '';
  form.major = '';
  modalVisible.value = true;
};

const showEditModal = (record) => {
  isEdit.value = true;
  form.id = record.id;
  form.studentId = record.studentId;
  form.name = record.name;
  form.grade = record.grade;
  form.major = record.major;
  modalVisible.value = true;
};

const viewEnrollments = (record) => {
  router.push({
    path: '/enrollments',
    query: { studentId: record.id }
  });
};

const handleModalOk = async () => {
  formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        if (isEdit.value) {
          await studentApi.update(form.id, form);
          Message.success('更新学生成功');
        } else {
          await studentApi.create(form);
          Message.success('添加学生成功');
        }
        modalVisible.value = false;
        await fetchStudents();
      } catch (error) {
        Message.error(isEdit.value ? '更新学生失败' : '添加学生失败');
        console.error(error);
      }
    }
  });
};

const deleteStudent = async (id) => {
  try {
    await studentApi.delete(id);
    Message.success('删除学生成功');
    await fetchStudents();
  } catch (error) {
    Message.error('删除学生失败');
    console.error(error);
  }
};

const handleSearch = () => {
  // 搜索逻辑已在 computed 中处理
};

const getGradeColor = (grade) => {
  switch (grade) {
    case '2020级': return 'red';
    case '2021级': return 'orange';
    case '2022级': return 'blue';
    case '2023级': return 'green';
    case '2024级': return 'purple';
    default: return 'gray';
  }
};

const getMajorColor = (major) => {
  const colors = ['blue', 'green', 'orange', 'purple', 'cyan', 'magenta', 'red', 'gold'];
  const index = major.length % colors.length;
  return colors[index];
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

.students-card {
  margin-top: var(--spacing-xl);
}

.card-actions {
  display: flex;
  gap: var(--spacing-sm);
  align-items: center;
}

.students-table {
  margin-top: var(--spacing-md);
}

.action-buttons {
  display: flex;
  gap: var(--spacing-xs);
  flex-wrap: wrap;
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

.action-button.view:hover {
  background-color: var(--success-light);
  color: var(--success-color);
}

.action-button.delete:hover {
  background-color: var(--danger-light);
  color: var(--danger-color);
}

.student-id-cell {
  display: flex;
  align-items: center;
  gap: var(--spacing-xs);
  color: var(--text-primary);
  font-weight: 500;
}

.student-id-cell .arco-icon {
  color: var(--primary-color);
}

.student-modal .arco-form-item {
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