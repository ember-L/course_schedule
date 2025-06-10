<template>
  <div class="teachers-container">
    <a-card title="教师管理">
      <template #extra>
        <a-button type="primary" @click="showAddModal">添加教师</a-button>
      </template>
      
      <a-table
        :loading="loading"
        :columns="columns"
        :data="teachers"
        :pagination="{
          showTotal: true,
          pageSize: 10
        }"
      >
        <template #operations="{ record }">
          <a-space>
            <a-button type="text" @click="showEditModal(record)">编辑</a-button>
            <a-popconfirm
              content="确定要删除这个教师吗？"
              @ok="deleteTeacher(record.id)"
            >
              <a-button type="text" status="danger">删除</a-button>
            </a-popconfirm>
          </a-space>
        </template>
      </a-table>
    </a-card>
    
    <a-modal
      v-model:visible="modalVisible"
      :title="isEdit ? '编辑教师' : '添加教师'"
      @ok="handleModalOk"
      @cancel="modalVisible = false"
    >
      <a-form :model="form" ref="formRef" :rules="rules">
        <a-form-item field="name" label="姓名">
          <a-input v-model="form.name" placeholder="请输入姓名" />
        </a-form-item>
        <a-form-item field="title" label="职称">
          <a-input v-model="form.title" placeholder="请输入职称" />
        </a-form-item>
        <a-form-item field="email" label="邮箱">
          <a-input v-model="form.email" placeholder="请输入邮箱" />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue';
import { Message } from '@arco-design/web-vue';
import { teacherApi } from '../api/modules';

const loading = ref(false);
const teachers = ref([]);
const modalVisible = ref(false);
const isEdit = ref(false);
const formRef = ref(null);

const form = reactive({
  id: null,
  name: '',
  title: '',
  email: ''
});

const rules = {
  name: [{ required: true, message: '请输入姓名' }],
  email: [
    { required: true, message: '请输入邮箱' },
    { type: 'email', message: '请输入有效的邮箱地址' }
  ]
};

const columns = [
  {
    title: '姓名',
    dataIndex: 'name'
  },
  {
    title: '职称',
    dataIndex: 'title'
  },
  {
    title: '邮箱',
    dataIndex: 'email'
  },
  {
    title: '操作',
    slotName: 'operations',
    width: 150
  }
];

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
</script>

<style scoped>
.teachers-container {
  padding: 20px;
}
</style>