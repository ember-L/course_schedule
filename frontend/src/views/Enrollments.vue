<template>
  <div class="enrollments-container">
    <a-card :title="getPageTitle()">
      <template #extra>
        <a-space>
          <!-- 管理员可以选择学生查看选课记录 -->
          <a-select
            v-if="isAdmin && !studentId"
            v-model="selectedStudent"
            placeholder="选择学生"
            style="width: 200px"
            @change="handleStudentChange"
          >
            <a-option v-for="student in students" :key="student.id" :value="student.id">
              {{ student.name }} ({{ student.studentId }})
            </a-option>
          </a-select>
          
          <!-- 教师可以选择自己的课程查看学生选课记录 -->
          <a-select
            v-if="isTeacher && !sectionId"
            v-model="selectedSection"
            placeholder="选择课程"
            style="width: 300px"
            @change="handleSectionChange"
          >
            <a-option v-for="section in mySections" :key="section.id" :value="section.id">
              {{ section.course.name }} - {{ getDayName(section.timeSlot.dayOfWeek) }} {{ section.timeSlot.startTime }}-{{ section.timeSlot.endTime }}
            </a-option>
          </a-select>
          
          <!-- 管理员和教师可以添加选课 -->
          <a-button 
            v-if="(isAdmin && studentId) || (isTeacher && sectionId)" 
            type="primary" 
            @click="showAddModal"
          >
            添加选课
          </a-button>
          
          <!-- 返回按钮 -->
          <a-button v-if="studentId || sectionId" @click="backToList">返回列表</a-button>
        </a-space>
      </template>
      
      <a-table
        :loading="loading"
        :columns="getColumns()"
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
            <a-button 
              v-if="isTeacher || isAdmin" 
              type="text" 
              size="mini" 
              @click="editGrade(record)"
            >
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
        <!-- 管理员添加选课时需要选择学生 -->
        <a-form-item v-if="isAdmin" field="studentId" label="学生">
          <a-select v-model="form.studentId" placeholder="请选择学生">
            <a-option v-for="student in students" :key="student.id" :value="student.id">
              {{ student.name }} ({{ student.studentId }})
            </a-option>
          </a-select>
        </a-form-item>
        
        <!-- 教师添加选课时需要选择学生 -->
        <a-form-item v-if="isTeacher" field="studentId" label="学生">
          <a-select v-model="form.studentId" placeholder="请选择学生">
            <a-option v-for="student in availableStudents" :key="student.id" :value="student.id">
              {{ student.name }} ({{ student.studentId }})
            </a-option>
          </a-select>
        </a-form-item>
        
        <!-- 管理员添加选课时需要选择课程 -->
        <a-form-item v-if="isAdmin" field="sectionId" label="课程">
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
import { useRoute } from 'vue-router';
import { useAuthStore } from '../store/authStore';
import { enrollmentApi, studentApi, sectionApi } from '../api/modules';

const route = useRoute();
const authStore = useAuthStore();

const loading = ref(false);
const enrollments = ref([]);
const sections = ref([]);
const students = ref([]);
const mySections = ref([]); // 教师的课程
const modalVisible = ref(false);
const formRef = ref(null);
const studentId = ref(null);
const studentName = ref('');
const sectionId = ref(null);
const sectionName = ref('');
const selectedStudent = ref(null);
const selectedSection = ref(null);
const editingGrade = ref(null);
const gradeValue = ref('');

const form = reactive({
  studentId: null,
  sectionId: null
});

const rules = {
  studentId: [{ required: true, message: '请选择学生' }],
  sectionId: [{ required: true, message: '请选择课程' }]
};

// 计算属性
const isAdmin = computed(() => authStore.isAdmin);
const isTeacher = computed(() => authStore.isTeacher);
const isStudent = computed(() => authStore.isStudent);

const availableSections = computed(() => {
  // 过滤掉已经选择的课程
  const enrolledSectionIds = enrollments.value.map(e => e.sectionId);
  return sections.value.filter(section => !enrolledSectionIds.includes(section.id));
});

const availableStudents = computed(() => {
  if (!isTeacher.value || !sectionId.value) return students.value;
  
  // 过滤掉已经选过该课程的学生
  const enrolledStudentIds = enrollments.value.map(e => e.studentId);
  return students.value.filter(student => !enrolledStudentIds.includes(student.id));
});

const getDayName = (day) => {
  const days = ['周日', '周一', '周二', '周三', '周四', '周五', '周六'];
  return days[day] || '';
};

const getPageTitle = () => {
  if (isStudent.value) {
    return '我的选课记录';
  } else if (isTeacher.value) {
    if (sectionId.value) {
      return `${sectionName.value}的学生选课记录`;
    }
    return '我的课程选课管理';
  } else if (isAdmin.value) {
    if (studentId.value) {
      return `${studentName.value}的选课记录`;
    }
    return '选课管理';
  }
  return '选课管理';
};

const getColumns = () => {
  const baseColumns = [
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
    }
  ];

  // 根据用户角色和查看模式添加不同的列
  if (isStudent.value) {
    // 学生查看自己的选课记录
    baseColumns.unshift({
      title: '课程',
      dataIndex: 'section.course.name'
    });
    baseColumns.splice(1, 0, {
      title: '教师',
      dataIndex: 'section.teacher.name'
    });
  } else if (isTeacher.value) {
    if (sectionId.value) {
      // 教师查看指定课程的学生选课记录
      baseColumns.unshift({
        title: '学生',
        dataIndex: 'student.name'
      });
      baseColumns.unshift({
        title: '课程',
        dataIndex: 'section.course.name'
      });
    } else {
      // 教师查看自己所有课程的选课记录
      baseColumns.unshift({
        title: '学生',
        dataIndex: 'student.name'
      });
      baseColumns.unshift({
        title: '课程',
        dataIndex: 'section.course.name'
      });
    }
  } else if (isAdmin.value) {
    // 管理员查看选课记录
    baseColumns.unshift({
      title: '学生',
      dataIndex: 'student.name'
    });
    baseColumns.unshift({
      title: '课程',
      dataIndex: 'section.course.name'
    });
    baseColumns.splice(2, 0, {
      title: '教师',
      dataIndex: 'section.teacher.name'
    });
  }

  // 管理员、教师和学生都可以看到操作列
  if (isAdmin.value || isTeacher.value || isStudent.value) {
    baseColumns.push({
      title: '操作',
      slotName: 'operations',
      width: 100
    });
  }

  return baseColumns;
};

onMounted(async () => {
  // 检查URL参数中是否有studentId或sectionId
  if (route.query.studentId) {
    studentId.value = Number(route.query.studentId);
    selectedStudent.value = studentId.value;
  }
  if (route.query.sectionId) {
    sectionId.value = Number(route.query.sectionId);
    selectedSection.value = sectionId.value;
  }
  
  await loadData();
});

watch(selectedStudent, async (newVal) => {
  if (newVal) {
    studentId.value = newVal;
    await loadEnrollments();
  }
});

watch(selectedSection, async (newVal) => {
  if (newVal) {
    sectionId.value = newVal;
    await loadEnrollments();
  }
});

const loadData = async () => {
  loading.value = true;
  try {
    if (isAdmin.value) {
      // 管理员加载所有学生和课程安排数据
      const [studentsData, sectionsData] = await Promise.all([
        studentApi.getAll(),
        sectionApi.getAll()
      ]);
      students.value = studentsData;
      sections.value = sectionsData;
    } else if (isTeacher.value) {
      // 教师加载自己的课程和所有学生
      const teacherId = authStore.currentUser.id;
      const [sectionsData, studentsData] = await Promise.all([
        sectionApi.getByTeacher(teacherId),
        studentApi.getAll()
      ]);
      mySections.value = sectionsData;
      students.value = studentsData;
    }
    
    // 如果有选定的学生或课程，加载选课记录
    if (studentId.value || sectionId.value) {
      await loadEnrollments();
    } else if (isStudent.value) {
      // 学生直接加载自己的选课记录
      await loadEnrollments();
    } else if (isTeacher.value) {
      // 教师加载自己所有课程的选课记录
      await loadMyCourseEnrollments();
    }
  } catch (error) {
    Message.error('加载数据失败');
    console.error(error);
  } finally {
    loading.value = false;
  }
};

const loadEnrollments = async () => {
  loading.value = true;
  try {
    let enrollmentsData;
    
    if (isStudent.value) {
      // 学生查看自己的选课记录
      enrollmentsData = await enrollmentApi.getMy();
    } else if (isTeacher.value && sectionId.value) {
      // 教师查看指定课程的学生选课记录
      enrollmentsData = await enrollmentApi.getByTeacherSection(sectionId.value);
    } else if (isAdmin.value && studentId.value) {
      // 管理员查看指定学生的选课记录
      enrollmentsData = await enrollmentApi.getByStudent(studentId.value);
    }
    
    enrollments.value = enrollmentsData || [];
    
    // 获取学生姓名或课程名称
    if (isAdmin.value && studentId.value) {
      const student = students.value.find(s => s.id === studentId.value);
      if (student) {
        studentName.value = student.name;
      }
    } else if (isTeacher.value && sectionId.value) {
      const section = mySections.value.find(s => s.id === sectionId.value);
      if (section) {
        sectionName.value = section.course.name;
      }
    }
  } catch (error) {
    Message.error('加载选课记录失败');
    console.error(error);
  } finally {
    loading.value = false;
  }
};

const loadMyCourseEnrollments = async () => {
  loading.value = true;
  try {
    const enrollmentsData = await enrollmentApi.getMyCourses();
    enrollments.value = enrollmentsData || [];
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

const handleSectionChange = (value) => {
  sectionId.value = value;
};

const showAddModal = () => {
  // 重置表单数据
  form.studentId = null;
  form.sectionId = null;
  
  if (isAdmin.value) {
    form.studentId = studentId.value;
  } else if (isTeacher.value) {
    form.sectionId = sectionId.value; // 教师已经选择了课程，所以设置sectionId
  }
  modalVisible.value = true;
};

const handleModalOk = async () => {
  try {
    // 验证表单数据
    if (isAdmin.value && (!form.studentId || !form.sectionId)) {
      Message.error('请选择学生和课程');
      return;
    } else if (isTeacher.value && (!form.studentId || !form.sectionId)) {
      Message.error('请选择学生');
      return;
    }
    
    if (isAdmin.value) {
      await enrollmentApi.create({
        studentId: form.studentId,
        sectionId: form.sectionId,
        enrollTime: new Date().toISOString()
      });
    } else if (isTeacher.value) {
      // 教师添加学生到自己的课程
      await enrollmentApi.create({
        studentId: form.studentId,
        sectionId: form.sectionId,
        enrollTime: new Date().toISOString()
      });
    }
    
    Message.success('选课成功');
    modalVisible.value = false;
    await loadEnrollments();
  } catch (error) {
    Message.error('选课失败');
    console.error(error);
  }
};

const deleteEnrollment = async (id) => {
  try {
    if (isStudent.value) {
      await enrollmentApi.drop(id);
    } else {
      await enrollmentApi.delete(id);
    }
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
    if (isTeacher.value) {
      await enrollmentApi.updateGrade(id, gradeValue.value);
    } else {
      await enrollmentApi.updateGrade(id, gradeValue.value);
    }
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
  sectionId.value = null;
  selectedStudent.value = null;
  selectedSection.value = null;
  enrollments.value = [];
  studentName.value = '';
  sectionName.value = '';
};
</script>

<style scoped>
.enrollments-container {
  padding: 20px;
}
</style>