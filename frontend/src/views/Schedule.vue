<template>
  <div class="schedule-container">
    <a-card :title="pageTitle">
      <template #extra>
        <a-space>
          <!-- 管理员选择用户 -->
          <a-select
            v-if="isAdmin"
            v-model="selectedUser"
            placeholder="选择用户"
            style="width: 200px"
            @change="loadUserSchedule"
          >
            <a-option-group label="学生">
              <a-option 
                v-for="student in students" 
                :key="`student-${student.id}`" 
                :value="`student-${student.id}`"
              >
                {{ student.name }} ({{ student.studentId }}) - 学生
              </a-option>
            </a-option-group>
            <a-option-group label="教师">
              <a-option 
                v-for="teacher in teachers" 
                :key="`teacher-${teacher.id}`" 
                :value="`teacher-${teacher.id}`"
              >
                {{ teacher.name }} - 教师
              </a-option>
            </a-option-group>
          </a-select>

          <!-- 学期选择 -->
          <a-select
            v-model="selectedSemester"
            placeholder="选择学期"
            style="width: 150px"
            @change="filterBySemester"
          >
            <a-option v-for="sem in semesters" :key="sem" :value="sem">{{ sem }}</a-option>
          </a-select>

          <!-- 导出功能 -->
          <a-dropdown>
            <a-button type="primary">
              导出课表
              <icon-down />
            </a-button>
            <template #content>
              <a-menu @click="handleExport">
                <a-menu-item key="pdf">导出为 PDF</a-menu-item>
                <a-menu-item key="excel">导出为 Excel</a-menu-item>
              </a-menu>
            </template>
          </a-dropdown>
        </a-space>
      </template>
      
      <!-- 加载状态 -->
      <div v-if="loading" class="loading-container">
        <a-spin size="large" />
        <div class="loading-text">正在加载课表...</div>
      </div>

      <!-- 空状态 -->
      <div v-else-if="!currentUser && !selectedUser && !loading" class="empty-container">
        <a-empty description="请选择用户查看课表" />
      </div>

      <!-- 课表内容 -->
      <div v-else-if="currentUser || selectedUser">
        <!-- 用户信息 -->
        <div class="user-info-section">
          <a-descriptions :column="3" bordered>
            <a-descriptions-item label="姓名">
              {{ currentUser?.name || '未知' }}
            </a-descriptions-item>
            <a-descriptions-item label="类型">
              {{ userTypeText }}
            </a-descriptions-item>
            <a-descriptions-item label="学期">
              {{ selectedSemester || '全部' }}
            </a-descriptions-item>
          </a-descriptions>
        </div>

        <!-- 课表统计 -->
        <div class="schedule-stats">
          <a-row :gutter="16">
            <a-col :span="6">
              <a-statistic title="总课程数" :value="filteredSections.length" />
            </a-col>
            <a-col :span="6">
              <a-statistic title="本周课程" :value="thisWeekCourses" />
            </a-col>
            <a-col :span="6">
              <a-statistic title="今日课程" :value="todayCourses" />
            </a-col>
            <a-col :span="6">
              <a-statistic  title="总课时" :value="totalHours" />
            </a-col>
          </a-row>
        </div>

        <!-- 课表组件 -->
        <time-table :sections="filteredSections" />
      </div>
    </a-card>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { Message } from '@arco-design/web-vue';
import { useAuthStore } from '@/store/authStore';
import TimeTable from '@/components/TimeTable.vue';
import { studentApi, teacherApi, enrollmentApi, sectionApi, scheduleApi } from '@/api/modules';
import html2pdf from 'html2pdf.js';
import * as XLSX from 'xlsx';

// 状态管理
const authStore = useAuthStore();
const loading = ref(false);
const currentUser = ref(null);
const students = ref([]);
const teachers = ref([]);
const sections = ref([]);
const selectedUser = ref(null);
const selectedSemester = ref('');
const semesters = ref([]);

// 计算属性
const isAdmin = computed(() => authStore.isAdmin);
const isTeacher = computed(() => authStore.isTeacher);
const isStudent = computed(() => authStore.isStudent);

const pageTitle = computed(() => {
  if (isAdmin.value) {
    return selectedUser.value ? '用户课表' : '课表管理';
  }
  return '我的课表';
});

const userTypeText = computed(() => {
  if (!currentUser.value) return '';
  if (isStudent.value) return '学生';
  if (isTeacher.value) return '教师';
  return '未知';
});

// 根据学期筛选课程
const filteredSections = computed(() => {
  if (!selectedSemester.value) {
    return sections.value;
  }
  return sections.value.filter(section => section.semester === selectedSemester.value);
});

// 统计信息
const thisWeekCourses = computed(() => {
  const today = new Date();
  const startOfWeek = new Date(today.setDate(today.getDate() - today.getDay() + 1));
  const endOfWeek = new Date(today.setDate(today.getDate() + 6));
  
  return filteredSections.value.filter(section => {
    const sectionDate = new Date(section.semester);
    return sectionDate >= startOfWeek && sectionDate <= endOfWeek;
  }).length;
});

const todayCourses = computed(() => {
  const today = new Date();
  const dayOfWeek = today.getDay() || 7; // 转换为1-7
  
  return filteredSections.value.filter(section => 
    section.timeSlot?.dayOfWeek === dayOfWeek
  ).length;
});

const totalHours = computed(() => {
  return filteredSections.value.reduce((total, section) => {
    const startTime = section.timeSlot?.startTime;
    const endTime = section.timeSlot?.endTime;
    if (startTime && endTime) {
      const start = new Date(`2000-01-01 ${startTime}`);
      const end = new Date(`2000-01-01 ${endTime}`);
      return total + (end - start) / (1000 * 60 * 60);
    }
    return total;
  }, 0);
});

// 生命周期
onMounted(async () => {
  await initializeData();
});

// 初始化数据
const initializeData = async () => {
  if (isAdmin.value) {
    await loadAllUsers();
  } else {
    const userLoaded = await loadCurrentUser();
    if (userLoaded) {
      await loadUserSchedule();
    }
  }
};

// 加载当前用户信息
const loadCurrentUser = async () => {
  loading.value = true;
  try {
    const userId = authStore.currentUser?.id;
    if (!userId) {
      Message.error('用户信息获取失败');
      return false; // 返回false表示加载失败
    }
    currentUser.value = userId;
    return true; // 返回true表示加载成功
  } catch (error) {
    Message.error('获取用户信息失败');
    console.error(error);
    return false; // 返回false表示加载失败
  } finally {
    loading.value = false;
  }
};

// 加载所有用户（管理员）
const loadAllUsers = async () => {
  loading.value = true;
  try {
    const [studentsData, teachersData] = await Promise.all([
      studentApi.getAll(),
      teacherApi.getAll()
    ]);
    
    students.value = studentsData || [];
    teachers.value = teachersData || [];
  } catch (error) {
    Message.error('获取用户列表失败');
    console.error(error);
    students.value = [];
    teachers.value = [];
  } finally {
    loading.value = false;
  }
};

// 加载用户课表
const loadUserSchedule = async () => {
  loading.value = true;
  try {
    let sectionsData = [];

    if (isAdmin.value && selectedUser.value) {
      // 管理员查看指定用户的课表
      const [userType, userId] = selectedUser.value.split('-');
      
      if (!userId) {
        Message.error('无效的用户ID');
        return;
      }
      
      if (userType === 'student') {
        // 获取学生的选课记录
        const enrollments = await enrollmentApi.getByStudent(parseInt(userId));
        const sectionIds = enrollments.map(e => e.sectionId);
        if (sectionIds.length > 0) {
          const allSections = await sectionApi.getAll();
          sectionsData = allSections.filter(section => sectionIds.includes(section.id));
        }
      } else if (userType === 'teacher') {
        // 获取教师的课程安排
        const allSections = await sectionApi.getAll();
        sectionsData = allSections.filter(section => section.teacherId === parseInt(userId));
      } else {
        Message.error('无效的用户类型');
        return;
      }
    } else if (isStudent.value) {
      // 学生查看自己的课表
      if (!currentUser.value) {
        Message.error('用户信息不完整，请重新登录');
        return;
      }
      // const enrollments = await enrollmentApi.getByStudent(currentUser.value);
      // const sectionIds = enrollments.map(e => e.sectionId);
      // if (sectionIds.length > 0) {
      //   const allSections = await sectionApi.getAll();
      //   sectionsData = allSections.filter(section => sectionIds.includes(section.id));
      // }
      sectionsData = await scheduleApi.getMySchedule();
    } else if (isTeacher.value) {
      // 教师查看自己的课表
      if (!currentUser.value) {
        Message.error('用户信息不完整，请重新登录');
        return;
      }
      // const allSections = await sectionApi.getAll();
      // sectionsData = allSections.filter(section => section.teacherId === currentUser.value);
      sectionsData = await scheduleApi.getMySchedule();
    }

    sections.value = sectionsData;
    
    // 提取唯一的学期列表
    const semesterSet = new Set(sectionsData.map(section => section.semester));
    semesters.value = Array.from(semesterSet).sort();
    
    // 默认选择第一个学期
    if (semesters.value.length > 0 && !selectedSemester.value) {
      selectedSemester.value = semesters.value[0];
    }
  } catch (error) {
    Message.error('获取课表失败');
    console.error(error);
  } finally {
    loading.value = false;
  }
};

// 学期筛选
const filterBySemester = () => {
  // 通过计算属性自动筛选
};

// 导出处理
const handleExport = async ({ key }) => {
  if (key === 'pdf') {
    await exportToPDF();
  } else {
    exportToExcel();
  }
};

// 时间段定义
const timeSlots = [
  "08:00-09:40",
  "10:00-11:40",
  "13:30-15:10",
  "15:30-17:10",
  "18:30-20:10"
];

// 导出PDF
const exportToPDF = () => {
  const element = document.querySelector('.timetable-container');
  if (!element) {
    Message.error('未找到课表内容');
    return;
  }

  const fileName = isAdmin.value 
    ? `课表_${selectedUser.value}_${selectedSemester.value}.pdf`
    : `课表_${currentUser.value?.name || '未知用户'}_${selectedSemester.value}.pdf`;

  const opt = {
    margin: 0.5,
    filename: fileName,
    image: { type: 'jpeg', quality: 0.98 },
    html2canvas: { scale: 2 },
    jsPDF: { unit: 'in', format: 'a4', orientation: 'portrait' }
  };

  html2pdf().set(opt).from(element).save();
};

// 导出Excel
const exportToExcel = () => {
  try {
    const days = ["周一", "周二", "周三", "周四", "周五", "周六", "周日"];
    
    // 构建二维数组（行：时间段，列：星期）
    const sheetData = timeSlots.map(time => {
      const row = {};
      days.forEach((day, index) => {
        const dayNumber = index + 1;
        const section = filteredSections.value.find(
          sec => sec.timeSlot?.dayOfWeek === dayNumber && 
                 sec.timeSlot?.startTime === time.split("-")[0]
        );
        row[day] = section ? `${section.course?.name || '未知课程'}\n${section.classroom?.building || ''}${section.classroom?.room || ''}\n${section.teacher?.name || '未知教师'}` : "";
      });
      return row;
    });

    // 设置表头
    const worksheet = XLSX.utils.json_to_sheet(sheetData, {
      header: ["时间/星期", ...days],
      skipHeader: false
    });

    // 设置列宽
    worksheet['!cols'] = [
      { wch: 15 }, // 时间段列
      { wch: 20 }, // Monday
      { wch: 20 }, // Tuesday
      { wch: 20 }, // Wednesday
      { wch: 20 }, // Thursday
      { wch: 20 }, // Friday
      { wch: 20 }, // Saturday
      { wch: 20 }  // Sunday
    ];

    // 创建工作簿并保存
    const workbook = XLSX.utils.book_new();
    XLSX.utils.book_append_sheet(workbook, worksheet, '课表');
    
    const fileName = isAdmin.value 
      ? `课表_${selectedUser.value}_${selectedSemester.value}.xlsx`
      : `课表_${currentUser.value?.name || '未知用户'}_${selectedSemester.value}.xlsx`;
    
    XLSX.writeFile(workbook, fileName);
    Message.success('导出成功');
  } catch (error) {
    Message.error('导出失败，请重试');
    console.error('导出错误:', error);
  }
};
</script>

<style scoped>
.schedule-container {
  padding: 20px;
}

.loading-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 60px 0;
  gap: 16px;
}

.loading-text {
  color: var(--color-text-2);
  font-size: 14px;
}

.empty-container {
  padding: 60px 0;
  text-align: center;
}

.user-info-section {
  margin-bottom: 24px;
}

.schedule-stats {
  margin-bottom: 24px;
  padding: 16px;
  background: var(--color-fill-1);
  border-radius: var(--border-radius-medium);
}

/* 响应式设计 */
@media (max-width: 768px) {
  .schedule-container {
    padding: 12px;
  }
  
  .schedule-stats {
    padding: 12px;
  }
  
  .schedule-stats .arco-col {
    margin-bottom: 12px;
  }
}

/* 导出样式 */
@media print, pdf {
  .schedule-container {
    padding: 10px;
  }
  
  .user-info-section,
  .schedule-stats {
    display: none;
  }
  
  .timetable-container {
    font-size: 12px;
  }
}
</style>