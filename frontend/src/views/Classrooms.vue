<template>
  <div class="classrooms-container">
    <a-card title="教室管理">
      <template #extra>
        <a-button type="primary" @click="showAddModal">添加教室</a-button>
      </template>
      
      <a-table
        :loading="loading"
        :columns="columns"
        :data="classrooms"
        :pagination="{
          showTotal: true,
          pageSize: 10
        }"
      >
        <template #operations="{ record }">
          <a-space>
            <a-button type="text" @click="showEditModal(record)">编辑</a-button>
            <a-popconfirm
              content="确定要删除这个教室吗？"
              @ok="deleteClassroom(record.id)"
            >
              <a-button type="text" status="danger">删除</a-button>
            </a-popconfirm>
          </a-space>
        </template>
      </a-table>
    </a-card>
    
    <a-modal
      v-model:visible="modalVisible"
      :title="isEdit ? '编辑教室' : '添加教室'"
      @ok="handleModalOk"
      @cancel="modalVisible = false"
    >
      <a-form :model="form" ref="formRef" :rules="rules">
        <a-form-item field="building" label="教学楼">
          <a-input v-model="form.building" placeholder="请输入教学楼" />
        </a-form-item>
        <a-form-item field="room" label="房间号">
          <a-input v-model="form.room" placeholder="请输入房间号" />
        </a-form-item>
        <a-form-item field="capacity" label="容量">
          <a-input-number v-model="form.capacity" placeholder="请输入容量" />
        </a-form-item>
        <a-form-item field="type" label="教室类型">
          <a-select v-model="form.type" placeholder="请选择教室类型">
            <a-option value="普通教室">普通教室</a-option>
            <a-option value="实验室">实验室</a-option>
            <a-option value="多媒体教室">多媒体教室</a-option>
            <a-option value="阶梯教室">阶梯教室</a-option>
          </a-select>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue';
import { Message } from '@arco-design/web-vue';
import { classroomApi } from '../api/modules';

const loading = ref(false);
const classrooms = ref([]);
const modalVisible = ref(false);
const isEdit = ref(false);
const formRef = ref(null);

const form = reactive({
  id: null,
  building: '',
  room: '',
  capacity: 0,
  type: '普通教室'
});

const rules = {
  building: [{ required: true, message: '请输入教学楼' }],
  room: [{ required: true, message: '请输入房间号' }],
  capacity: [{ required: true, type: 'number', message: '请输入容量' }]
};

const columns = [
  {
    title: '教学楼',
    dataIndex: 'building'
  },
  {
    title: '房间号',
    dataIndex: 'room'
  },
  {
    title: '容量',
    dataIndex: 'capacity'
  },
  {
    title: '教室类型',
    dataIndex: 'type'
  },
  {
    title: '操作',
    slotName: 'operations',
    width: 150
  }
];

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
</script>

<style scoped>
.classrooms-container {
  padding: 20px;
}
</style>