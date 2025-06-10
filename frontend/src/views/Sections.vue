<template>
  <div class="sections-container">
    <a-card title="排课管理">
      <template #extra>
        <a-button type="primary" @click="showAddModal">添加排课</a-button>
      </template>
      
      <a-row :gutter="16" style="margin-bottom: 16px">
        <a-col :span="6">
          <a-select
            v-model="filterSemester"
            placeholder="选择学期"
            allow-clear
            style="width: 100%"
          >
            <a-option v-for="sem in semesters" :key="sem" :value="sem">{{ sem }}</a-option>
          </a-select>
        </a-col>
        <a-col :span="6">
          <a-select
            v-model="filterCourse"
            placeholder="选择课程"
            allow-clear
            style="width: 100%"
          >
            <a-option v-for="course in courses" :key="course.id" :value="course.id">
              {{ course.name }}
            </a-option>
          </a-select>
        </a-col>
        <a-col :span="6">
          <a-button type="outline" @click="handleFilter">筛选</a-button>
        </a-col>
      </a-row>
      
      <a-table
        :loading="loading"
        :columns="columns"
        :data="filteredSections"
        :pagination="{
          showTotal: true,
          pageSize: 10
        }"
      >
        <template #operations="{ record }">
          <a-space>
            <a-button type="text" @click="showEditModal(record)">编辑</a-button>
            <a-popconfirm
              content="确定要删除这个排课吗？"
              @ok="deleteSection(record.id)"
            >
              <a-button type="text" status="danger">删除</a-button>
            </a-popconfirm>
          </a-space>
        </template>
      </a-table>
    </a-card>
    
    <a-modal
      v-model:visible="modalVisible"
      :title="isEdit ? '编辑排课' : '添加排课'"
      @ok="handleModalOk"
      @cancel="modalVisible = false"
    >
      <a-form :model="form" ref="formRef" :rules="rules">
        <a-form-item field="courseId" label="课程">
          <a-select v-model="form.courseId" placeholder="请选择课程">
            <a-option v-for="course in courses" :key="course.id" :value="course.id">
              {{ course.name }}
            </a-option>
          </a-select>
        </a-form-item>
        <a-form-item field="teacherId" label="教师">
          <a-select v-model="form.teacherId" placeholder="请选择教师">
            <a-option v-for="teacher in teachers" :key="teacher.id" :value="teacher.id">
              {{ teacher.name }}
            </a-option>
          </a-select>
        </a-form-item>
        <a-form-item field="classroomId" label="教室">
          <a-select v-model="form.classroomId" placeholder="请选择教室">
            <a-option v-for="classroom in classrooms" :key="classroom.id" :value="classroom.id">
              {{ classroom.building }}-{{ classroom.room }}
            </a-option>
          </a-select>
        </a-form-item>
        <a-form-item field="timeSlotId" label="时间段">
          <a-select v-model="form.timeSlotId" placeholder="请选择时间段">
            <a-option v-for="timeSlot in timeSlots" :key="timeSlot.id" :value="timeSlot.id">
              {{ getDayName(timeSlot.dayOfWeek) }} {{ timeSlot.startTime }}-{{ timeSlot.endTime }}
            </a-option>
          </a-select>
        </a-form-item>
        <a-form-item field="semester" label="学期">
          <a-input v-model="form.semester" placeholder="请输入学期，如：2023-2024-1" />
        </a-form-item>
        <a-form-item field="maxStudents" label="最大学生数">
          <a-input-number v-model="form.maxStudents" placeholder="请输入最大学生数" />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue';
import { Message } from '@arco-design/web-vue';
import { sectionApi, courseApi, teacherApi, classroomApi, timeSlotApi } from '../api/modules';

const loading = ref(false);
const sections = ref([]);
const courses = ref([]);
const teachers = ref([]);
const classrooms = ref([]);
const timeSlots = ref([]);
const semesters = ref([]);
const modalVisible = ref(false);
const isEdit = ref(false);
const formRef = ref(null);
const filterSemester = ref('');
const filterCourse = ref('');

const form = reactive({
  id: null,
  courseId: null,
  teacherId: null,
  classroomId: null,
  timeSlotId: null,
  semester: '',
  maxStudents: 0
});

const rules = {
  courseId: [{ required: true, message: '请选择课程' }],
  teacherId: [{ required: true, message: '请选择教师' }],
  classroomId: [{ required: true, message: '请选择教室' }],
  timeSlotId: [{ required: true, message: '请选择时间段' }],
  semester: [{ required: true, message: '请输入学期' }],
  maxStudents: [{ required: true, type: 'number', message: '请输入最大学生数' }]
};

const columns = [
  {
    title: '课程',
    dataIndex: 'course.name'
  },
  {
    title: '教师',
    dataIndex: 'teacher.name'
  },
  {
    title: '教室',
    render: ({ record }) => `${record.classroom.building}-${record.classroom.room}`
  },
  {
    title: '时间',
    render: ({ record }) => `${getDayName(record.timeSlot.dayOfWeek)} ${record.timeSlot.startTime}-${record.timeSlot.endTime}`
  },
  {
    title: '学期',
    dataIndex: 'semester'
  },
  {
    title: '最大学生数',
    dataIndex: 'maxStudents'
  },
  {
    title: '操作',
    slotName: 'operations',
    width: 150
  }
];

const filteredSections = computed(() => {
  let result = sections.value;
  
  if (filterSemester.value) {
    result = result.filter(section => section.semester === filterSemester.value);
  }
  
  if (filterCourse.value) {
    result = result.filter(section => section.courseId === filterCourse.value);
  }
  
  return result;
});

const getDayName = (day) => {
  const days = ['周日', '周一', '周二', '周三', '周四', '周五', '周六'];
  return days[day] || '';
};

onMounted(async () => {
  await fetchAllData();
});

const fetchAllData = async () => {
  loading.value = true;
  try {
    // 并行加载所有需要的数据
    const [sectionsData, coursesData, teachersData, classroomsData, timeSlotsData] = await Promise.all([
      sectionApi.getAll(),
      courseApi.getAll(),
      teacherApi.getAll(),
      classroomApi.getAll(),
      timeSlotApi.getAll()
    ]);
    
    sections.value = sectionsData;
    courses.value = coursesData;
    teachers.value = teachersData;
    classrooms.value = classroomsData;
    timeSlots.value = timeSlotsData;
    
    // 提取唯一的学期列表
    const semesterSet = new Set(sections.value.map(section => section.semester));
    semesters.value = Array.from(semesterSet);
  } catch (error) {
    Message.error('加载数据失败');
    console.error(error);
  } finally {
    loading.value = false;
  }
};

const handleFilter = () => {
  // 筛选已在计算属性中处理
};

const showAddModal = () => {
  isEdit.value = false;
  form.id = null;
  form.courseId = null;
  form.teacherId = null;
  form.classroomId = null;
  form.timeSlotId = null;
  form.semester = '';
  form.maxStudents = 0;
  modalVisible.value = true;
};

const showEditModal = (record) => {
  isEdit.value = true;
  form.id = record.id;
  form.courseId = record.courseId;
  form.teacherId = record.teacherId;
  form.classroomId = record.classroomId;
  form.timeSlotId = record.timeSlotId;
  form.semester = record.semester;
  form.maxStudents = record.maxStudents;
  modalVisible.value = true;
};

const handleModalOk = async () => {
  formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        // 先检查冲突
        const checkResult = await sectionApi.checkConflict({
          id: form.id, // 如果是编辑，需要排除自身
          teacherId: form.teacherId,
          classroomId: form.classroomId,
          timeSlotId: form.timeSlotId,
          semester: form.semester
        });
        
        if (checkResult.hasConflict) {
          Message.error(`排课冲突: ${checkResult.message}`);
          return;
        }
        
        if (isEdit.value) {
          await sectionApi.update(form.id, form);
          Message.success('更新排课成功');
        } else {
          await sectionApi.create(form);
          Message.success('添加排课成功');
        }
        modalVisible.value = false;
        await fetchAllData();
      } catch (error) {
        Message.error(isEdit.value ? '更新排课失败' : '添加排课失败');
        console.error(error);
      }
    }
  });
};

const deleteSection = async (id) => {
  try {
    await sectionApi.delete(id);
    Message.success('删除排课成功');
    await fetchAllData();
  } catch (error) {
    Message.error('删除排课失败');
    console.error(error);
  }
};
</script>

<style scoped>
.sections-container {
  padding: 20px;
}
</style>