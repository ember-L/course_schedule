<template>
  <div class="page-container fade-in">
    <!-- 页面头部 -->
    <div class="page-header">
      <div class="header-content">
        <div class="header-left">
          <h1 class="page-title">教室管理</h1>
          <p class="page-subtitle">管理系统中的所有教室信息</p>
        </div>
        <div class="header-right">
          <a-button type="primary" @click="showAddModal" class="add-button">
            <icon-plus />
            添加教室
          </a-button>
        </div>
      </div>
    </div>

    <!-- 统计卡片 -->
    <div class="card-grid">
      <div class="statistic-card">
        <div class="statistic-content">
          <div class="statistic-icon">
            <icon-home />
          </div>
          <div class="statistic-info">
            <div class="statistic-value">{{ classrooms.length }}</div>
            <div class="statistic-label">总教室数</div>
          </div>
        </div>
      </div>
      
      <div class="statistic-card">
        <div class="statistic-content">
          <div class="statistic-icon">
            <icon-user-group />
          </div>
          <div class="statistic-info">
            <div class="statistic-value">{{ totalCapacity }}</div>
            <div class="statistic-label">总容量</div>
          </div>
        </div>
      </div>
      
      <div class="statistic-card">
        <div class="statistic-content">
          <div class="statistic-icon">
            <icon-bulb />
          </div>
          <div class="statistic-info">
            <div class="statistic-value">{{ buildingCount }}</div>
            <div class="statistic-label">教学楼数量</div>
          </div>
        </div>
      </div>
    </div>

    <!-- 教室列表 -->
    <a-card class="classrooms-card" title="教室列表">
      <template #extra>
        <div class="card-actions">
          <a-input-search
            v-model="searchKeyword"
            placeholder="搜索教室..."
            style="width: 250px"
            @search="handleSearch"
          />
        </div>
      </template>
      
      <a-table
        :loading="loading"
        :columns="columns"
        :data="filteredClassrooms"
        :pagination="{
          showTotal: true,
          pageSize: 10,
          showJumper: true,
          showPageSize: true
        }"
        class="classrooms-table"
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
              content="确定要删除这个教室吗？"
              @ok="deleteClassroom(record.id)"
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
        
        <template #capacity="{ record }">
          <a-tag :color="getCapacityColor(record.capacity)">
            {{ record.capacity }} 人
          </a-tag>
        </template>
        
        <template #type="{ record }">
          <a-tag :color="getTypeColor(record.type)">
            {{ record.type }}
          </a-tag>
        </template>
        
        <template #location="{ record }">
          <div class="location-cell">
            <icon-location />
            <span>{{ record.building }} - {{ record.room }}</span>
          </div>
        </template>
      </a-table>
    </a-card>
    
    <!-- 添加/编辑教室模态框 -->
    <a-modal
      v-model:visible="modalVisible"
      :title="isEdit ? '编辑教室' : '添加教室'"
      @ok="handleModalOk"
      @cancel="modalVisible = false"
      :width="600"
      class="classroom-modal"
    >
      <a-form :model="form" ref="formRef" :rules="rules" layout="vertical">
        <a-row :gutter="16">
          <a-col :span="12">
            <a-form-item field="building" label="教学楼">
              <a-input 
                v-model="form.building" 
                placeholder="请输入教学楼"
                :maxLength="20"
              />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item field="room" label="房间号">
              <a-input 
                v-model="form.room" 
                placeholder="请输入房间号"
                :maxLength="20"
              />
            </a-form-item>
          </a-col>
        </a-row>
        
        <a-row :gutter="16">
          <a-col :span="12">
            <a-form-item field="capacity" label="容量">
              <a-input-number 
                v-model="form.capacity" 
                placeholder="请输入容量"
                :min="1"
                :max="500"
                style="width: 100%"
              />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item field="type" label="教室类型">
              <a-select 
                v-model="form.type" 
                placeholder="请选择教室类型"
                allow-clear
              >
                <a-option value="普通教室">普通教室</a-option>
                <a-option value="实验室">实验室</a-option>
                <a-option value="多媒体教室">多媒体教室</a-option>
                <a-option value="阶梯教室">阶梯教室</a-option>
                <a-option value="计算机房">计算机房</a-option>
                <a-option value="语音室">语音室</a-option>
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
import { classroomApi } from '../api/modules';

const loading = ref(false);
const classrooms = ref([]);
const modalVisible = ref(false);
const isEdit = ref(false);
const formRef = ref(null);
const searchKeyword = ref('');

const form = reactive({
  id: null,
  building: '',
  room: '',
  capacity: 0,
  type: '普通教室'
});

const rules = {
  building: [
    { required: true, message: '请输入教学楼' },
    { minLength: 2, message: '教学楼名称至少2个字符' }
  ],
  room: [
    { required: true, message: '请输入房间号' },
    { minLength: 2, message: '房间号至少2个字符' }
  ],
  capacity: [
    { required: true, type: 'number', message: '请输入容量' },
    { type: 'number', min: 1, max: 500, message: '容量范围1-500' }
  ],
  type: [
    { required: true, message: '请选择教室类型' }
  ]
};

const columns = [
  {
    title: '位置',
    dataIndex: 'location',
    slotName: 'location',
    width: 200,
    fixed: 'left'
  },
  {
    title: '容量',
    dataIndex: 'capacity',
    slotName: 'capacity',
    width: 100
  },
  {
    title: '教室类型',
    dataIndex: 'type',
    slotName: 'type',
    width: 120
  },
  {
    title: '操作',
    slotName: 'operations',
    width: 150,
    fixed: 'right'
  }
];

// 计算属性
const totalCapacity = computed(() => {
  return classrooms.value.reduce((sum, classroom) => sum + (classroom.capacity || 0), 0);
});

const buildingCount = computed(() => {
  const buildings = new Set(classrooms.value.map(classroom => classroom.building));
  return buildings.size;
});

const filteredClassrooms = computed(() => {
  if (!searchKeyword.value) return classrooms.value;
  const keyword = searchKeyword.value.toLowerCase();
  return classrooms.value.filter(classroom => 
    classroom.building.toLowerCase().includes(keyword) ||
    classroom.room.toLowerCase().includes(keyword) ||
    classroom.type.toLowerCase().includes(keyword)
  );
});

onMounted(async () => {
  await fetchClassrooms();
});

const fetchClassrooms = async () => {
  loading.value = true;
  try {
    const data = await classroomApi.getAll();
    classrooms.value = data;
  } catch (error) {
    Message.error('获取教室列表失败');
    console.error(error);
  } finally {
    loading.value = false;
  }
};

const showAddModal = () => {
  isEdit.value = false;
  form.id = null;
  form.building = '';
  form.room = '';
  form.capacity = 0;
  form.type = '普通教室';
  modalVisible.value = true;
};

const showEditModal = (record) => {
  isEdit.value = true;
  form.id = record.id;
  form.building = record.building;
  form.room = record.room;
  form.capacity = record.capacity;
  form.type = record.type;
  modalVisible.value = true;
};

const handleModalOk = async () => {
  formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        if (isEdit.value) {
          await classroomApi.update(form.id, form);
          Message.success('更新教室成功');
        } else {
          await classroomApi.create(form);
          Message.success('添加教室成功');
        }
        modalVisible.value = false;
        await fetchClassrooms();
      } catch (error) {
        Message.error(isEdit.value ? '更新教室失败' : '添加教室失败');
        console.error(error);
      }
    }
  });
};

const deleteClassroom = async (id) => {
  try {
    await classroomApi.delete(id);
    Message.success('删除教室成功');
    await fetchClassrooms();
  } catch (error) {
    Message.error('删除教室失败');
    console.error(error);
  }
};

const handleSearch = () => {
  // 搜索逻辑已在 computed 中处理
};

const getCapacityColor = (capacity) => {
  if (capacity <= 30) return 'green';
  if (capacity <= 60) return 'blue';
  if (capacity <= 100) return 'orange';
  return 'red';
};

const getTypeColor = (type) => {
  switch (type) {
    case '普通教室': return 'blue';
    case '实验室': return 'green';
    case '多媒体教室': return 'purple';
    case '阶梯教室': return 'orange';
    case '计算机房': return 'cyan';
    case '语音室': return 'magenta';
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

.classrooms-card {
  margin-top: var(--spacing-xl);
}

.card-actions {
  display: flex;
  gap: var(--spacing-sm);
  align-items: center;
}

.classrooms-table {
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

.location-cell {
  display: flex;
  align-items: center;
  gap: var(--spacing-xs);
  color: var(--text-primary);
  font-weight: 500;
}

.location-cell .arco-icon {
  color: var(--primary-color);
}

.classroom-modal .arco-form-item {
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