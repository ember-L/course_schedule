<template>
  <div class="courses-container">
    <a-card title="课程管理">
      <template #extra>
        <a-button type="primary" @click="showAddModal">添加课程</a-button>
      </template>
      
      <a-table
        :loading="loading"
        :columns="columns"
        :data="courses"
        :pagination="{
          showTotal: true,
          pageSize: 10
        }"
      >
        <template #operations="{ record }">
          <a-space>
            <a-button type="text" @click="showEditModal(record)">编辑</a-button>
            <a-popconfirm
              content="确定要删除这个课程吗？"
              @ok="deleteCourse(record.id)"
            >
              <a-button type="text" status="danger">删除</a-button>
            </a-popconfirm>
          </a-space>
        </template>
      </a-table>
    </a-card>
    
    <a-modal
      v-model:visible="modalVisible"
      :title="isEdit ? '编辑课程' : '添加课程'"
      @ok="handleModalOk"
      @cancel="modalVisible = false"
    >
      <a-form :model="form" ref="formRef" :rules="rules">
        <a-form-item field="code" label="课程代码">
          <a-input v-model="form.code" placeholder="请输入课程代码" />
        </a-form-item>
        <a-form-item field="name" label="课程名称">
          <a-input v-model="form.name" placeholder="请输入课程名称" />
        </a-form-item>
        <a-form-item field="credits" label="学分">
          <a-input-number v-model="form.credits" placeholder="请输入学分" />
        </a-form-item>
        <a-form-item field="description" label="课程描述">
          <a-textarea v-model="form.description" placeholder="请输入课程描述" />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue';
import { Message } from '@arco-design/web-vue';
import { courseApi } from '../api/modules';

const loading = ref(false);
const courses = ref([]);
const modalVisible = ref(false);
const isEdit = ref(false);
const formRef = ref(null);

const form = reactive({
  id: null,
  code: '',
  name: '',
  credits: 0,
  description: ''
});

const rules = {
  code: [{ required: true, message: '请输入课程代码' }],
  name: [{ required: true, message: '请输入课程名称' }],
  credits: [{ required: true, type: 'number', message: '请输入学分' }]
};

const columns = [
  {
    title: '课程代码',
    dataIndex: 'code'
  },
  {
    title: '课程名称',
    dataIndex: 'name'
  },
  {
    title: '学分',
    dataIndex: 'credits'
  },
  {
    title: '课程描述',
    dataIndex: 'description',
    ellipsis: true
  },
  {
    title: '操作',
    slotName: 'operations',
    width: 150
  }
];

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
</script>

<style scoped>
.courses-container {
  padding: 20px;
}
</style>