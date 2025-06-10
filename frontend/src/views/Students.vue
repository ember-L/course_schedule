<template>
  <div class="students-container">
    <a-card title="学生管理">
      <template #extra>
        <a-button type="primary" @click="showAddModal">添加学生</a-button>
      </template>
      
      <a-table
        :loading="loading"
        :columns="columns"
        :data="students"
        :pagination="{
          showTotal: true,
          pageSize: 10
        }"
      >
        <template #operations="{ record }">
          <a-space>
            <a-button type="text" @click="showEditModal(record)">编辑</a-button>
            <a-button type="text" @click="viewEnrollments(record)">选课记录</a-button>
            <a-popconfirm
              content="确定要删除这个学生吗？"
              @ok="deleteStudent(record.id)"
            >
              <a-button type="text" status="danger">删除</a-button>
            </a-popconfirm>
          </a-space>
        </template>
      </a-table>
    </a-card>
    
    <a-modal
      v-model:visible="modalVisible"
      :title="isEdit ? '编辑学生' : '添加学生'"
      @ok="handleModalOk"
      @cancel="modalVisible = false"
    >
      <a-form :model="form" ref="formRef" :rules="rules">
        <a-form-item field="studentId" label="学号">
          <a-input v-model="form.studentId" placeholder="请输入学号" />
        </a-form-item>
        <a-form-item field="name" label="姓名">
          <a-input v-model="form.name" placeholder="请输入姓名" />
        </a-form-item>
        <a-form-item field="grade" label="年级">
          <a-input v-model="form.grade" placeholder="请输入年级" />
        </a-form-item>
        <a-form-item field="major" label="专业">
          <a-input v-model="form.major" placeholder="请输入专业" />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue';
import { Message } from '@arco-design/web-vue';
import { useRouter } from 'vue-router';
import { studentApi } from '../api/modules';

const router = useRouter();
const loading = ref(false);
const students = ref([]);
const modalVisible = ref(false);
const isEdit = ref(false);
const formRef = ref(null);

const form = reactive({
  id: null,
  studentId: '',
  name: '',
  grade: '',
  major: ''
});

const rules = {
  studentId: [{ required: true, message: '请输入学号' }],
  name: [{ required: true, message: '请输入姓名' }],
  grade: [{ required: true, message: '请输入年级' }],
  major: [{ required: true, message: '请输入专业' }]
};

const columns = [
  {
    title: '学号',
    dataIndex: 'studentId'
  },
  {
    title: '姓名',
    dataIndex: 'name'
  },
  {
    title: '年级',
    dataIndex: 'grade'
  },
  {
    title: '专业',
    dataIndex: 'major'
  },
  {
    title: '操作',
    slotName: 'operations',
    width: 250
  }
];

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
</script>

<style scoped>
.students-container {
  padding: 20px;
}
</style>