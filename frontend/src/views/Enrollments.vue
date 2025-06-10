<template>
  <div class="enrollments-container">
    <a-card :title="studentId ? `${studentName}的选课记录` : '选课管理'">
      <template #extra>
        <a-space>
          <a-select
            v-if="!studentId"
            v-model="selectedStudent"
            placeholder="选择学生"
            style="width: 200px"
            @change="handleStudentChange"
          >
            <a-option v-for="student in students" :key="student.id" :value="student.id">
              {{ student.name }} ({{ student.studentId }})
            </a-option>
          </a-select>
          <a-button type="primary" @click="showAddModal" :disabled="!studentId">添加选课</a-button>
          <a-button v-if="studentId" @click="backToList">返回列表</a-button>
        </a-space>
      </template>
      
      <a-table
        :loading="loading"
        :columns="columns"
        :data="enrollments"
        :pagination="{
          showTotal: true,
          pageSize: 10
        }"
      >
        <template #grade="{ record }">
          <a-input
            v-if="editingGrade === record.id"
            v-model="gradeValue"
            placeholder="输入成绩"
            @blur="saveGrade(record.id)"
            @keyup.enter="saveGrade(record.id)"
            style="width: 60px"
          />
          <a-space v-else>
            <span>{{ record.grade || '-' }}</span>
            <a-button type="text" size="mini" @click="editGrade(record)">
              <icon-edit />
            </a-button>
          </a-space>
        </template>
        <template #operations="{ record }">
          <a-popconfirm
            content="确定要退选这门课程吗？"
            @ok="deleteEnrollment(record.id)"
          >
            <a-button type="text" status="danger">退选</a-button>
          </a-popconfirm>
        </template>
      </a-table>
    </a-card>
    
    <a-modal
      v-model:visible="modalVisible"
      title="添加选课"
      @ok="handleModalOk"
      @cancel="modalVisible = false"
    >
      <a-form :model="form" ref="formRef" :rules="rules">
        <a-form-item field="sectionId" label="课程">
          <a-select v-model="form.sectionId" placeholder="请选择课程">
            <a-option v-for="section in availableSections" :key="section.id" :value="section.id">
              {{ section.course.name }} - {{ section.teacher.name }} - 
              {{ getDayName(section.timeSlot.dayOfWeek) }} {{ section.timeSlot.startTime }}-{{ section.timeSlot.endTime }}
            </a-option>
          </a-select>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted, watch } from 'vue';
import { Message } from '@arco-design/web-vue';
import { useRoute, useRouter } from 'vue-router';
import { enrollmentApi, studentApi, sectionApi } from '../api/modules';

const route = useRoute();
const router = useRouter();
const loading = ref(false);
const enrollments = ref([]);
const sections = ref([]);
const students = ref([]);
const modalVisible = ref(false);
const formRef = ref(null);
const studentId = ref(null);
const studentName = ref('');
const selectedStudent = ref(null);
const editingGrade = ref(null);
const gradeValue = ref('');

const form = reactive({
  studentId: null,
  sectionId: null
});

const rules = {
  sectionId: [{ required: true, message: '请选择课程' }]
};

const columns = [
  {
    title: '课程',
    dataIndex: 'section.course.name'
  },
  {
    title: '教师',
    dataIndex: 'section.teacher.name'
  },
  {
    title: '时间',
    render: ({ record }) => {
      const timeSlot = record.section.timeSlot;
      return `${getDayName(timeSlot.dayOfWeek)} ${timeSlot.startTime}-${timeSlot.endTime}`;
    }
  },
  {
    title: '地点',
    render: ({ record }) => {
      const classroom = record.section.classroom;
      return `${classroom.building}-${classroom.room}`;
    }
  },
  {
    title: '学期',
    dataIndex: 'section.semester'
  },
  {
    title: '成绩',
    slotName: 'grade'
  },
  {
    title: '操作',
    slotName: 'operations',
    width: 100
  }
];

const availableSections = computed(() => {
  // 过滤掉已经选择的课程
  const enrolledSectionIds = enrollments.value.map(e => e.sectionId);
  return sections.value.filter(section => !enrolledSectionIds.includes(section.id));
});

const getDayName = (day) => {
  const days = ['周日', '周一', '周二', '周三', '周四', '周五', '周六'];
  return days[day] || '';
};

onMounted(async () => {
  // 检查URL参数中是否有studentId
  if (route.query.studentId) {
    studentId.value = Number(route.query.studentId);
    selectedStudent.value = studentId.value;
  }
  
  await loadData();
});

watch(selectedStudent, async (newVal) => {
  if (newVal) {
    studentId.value = newVal;
    await loadEnrollments();
  }
});

const loadData = async () => {
  loading.value = true;
  try {
    // 加载学生和课程安排数据
    const [studentsData, sectionsData] = await Promise.all([
      studentApi.getAll(),
      sectionApi.getAll()
    ]);
    
    students.value = studentsData;
    sections.value = sectionsData;
    
    // 如果有选定的学生，加载该学生的选课记录
    if (studentId.value) {
      await loadEnrollments();
    }
  } catch (error) {
    Message.error('加载数据失败');
    console.error(error);
  } finally {
    loading.value = false;
  }
};

const loadEnrollments = async () => {
  if (!studentId.value) return;
  
  loading.value = true;
  try {
    const enrollmentsData = await enrollmentApi.getByStudent(studentId.value);
    enrollments.value = enrollmentsData;
    
    // 获取学生姓名
    const student = students.value.find(s => s.id === studentId.value);
    if (student) {
      studentName.value = student.name;
    }
  } catch (error) {
    Message.error('加载选课记录失败');
    console.error(error);
  } finally {
    loading.value = false;
  }
};

const handleStudentChange = (value) => {
  studentId.value = value;
};

const showAddModal = () => {
  form.studentId = studentId.value;
  form.sectionId = null;
  modalVisible.value = true;
};

const handleModalOk = async () => {
  formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        await enrollmentApi.create({
          studentId: form.studentId,
          sectionId: form.sectionId,
          enrollTime: new Date().toISOString()
        });
        Message.success('选课成功');
        modalVisible.value = false;
        await loadEnrollments();
      } catch (error) {
        Message.error('选课失败');
        console.error(error);
      }
    }
  });
};

const deleteEnrollment = async (id) => {
  try {
    await enrollmentApi.delete(id);
    Message.success('退选成功');
    await loadEnrollments();
  } catch (error) {
    Message.error('退选失败');
    console.error(error);
  }
};

const editGrade = (record) => {
  editingGrade.value = record.id;
  gradeValue.value = record.grade || '';
};

const saveGrade = async (id) => {
  try {
    await enrollmentApi.updateGrade(id, gradeValue.value);
    Message.success('成绩更新成功');
    editingGrade.value = null;
    await loadEnrollments();
  } catch (error) {
    Message.error('成绩更新失败');
    console.error(error);
  }
};

const backToList = () => {
  studentId.value = null;
  selectedStudent.value = null;
  router.replace('/enrollments');
};
</script>

<style scoped>
.enrollments-container {
  padding: 20px;
}
</style>