<template>
  <div class="schedule-container">
    <a-card title="我的课表">
      <template #extra>
        <a-space>
          <a-select
            v-if="isAdmin"
            v-model="selectedStudent"
            placeholder="选择学生"
            style="width: 200px"
            @change="loadStudentSchedule"
          >
            <a-option v-for="student in students" :key="student.id" :value="student.id">
              {{ student.name }} ({{ student.studentId }})
            </a-option>
          </a-select>
          <a-select
            v-model="selectedSemester"
            placeholder="选择学期"
            style="width: 150px"
            @change="filterBySemester"
          >
            <a-option v-for="sem in semesters" :key="sem" :value="sem">{{ sem }}</a-option>
          </a-select>
                    <a-dropdown>
            <a-button type="primary">
              导出课表
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
      
      <div v-if="loading" class="loading-container">
        <a-spin />
      </div>
      <div v-else-if="!student && !selectedStudent" class="empty-container">
        <a-empty description="请选择学生查看课表" />
      </div>
      <div v-else>
        <time-table :sections="filteredSections" />
      </div>
    </a-card>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { Message } from '@arco-design/web-vue';
import TimeTable from '@/components/TimeTable.vue';
import { studentApi, enrollmentApi, sectionApi } from '@/api/modules';
import html2pdf from 'html2pdf.js';
import * as XLSX from 'xlsx';

const loading = ref(false);
const student = ref()
const students = ref([]);
const enrollments = ref([]);
const sections = ref([]);
const selectedStudent = ref(null);
const selectedSemester = ref('');
const semesters = ref([]);
const isAdmin = computed(() => localStorage.getItem('userType') === 'admin'); 

// 根据学期筛选课程
const filteredSections = computed(() => {
  if (!selectedSemester.value) {
    return sections.value;
  }
  return sections.value.filter(section => section.semester === selectedSemester.value);
});

onMounted(async () => {
  if(localStorage.getItem('userType') === 'student'){
      await loadStudents();
      await loadStudentSchedule();
  } else if (localStorage.getItem('userType') === "admin"){
    await loadStudents();
  }
});

const loadStudents = async () => {
  loading.value = true;
  try {
    // console.log(localStorage.getItem('user'));
      if(localStorage.getItem('userType') === 'student'){
        const userId = JSON.parse(localStorage.getItem('user')).id
         console.log(userId);
        const data = await studentApi.getById(userId);
        student.value = data;
      } else if (localStorage.getItem('userType') === "admin"){
        const data = await studentApi.getAll();
        students.value = data;
      }



  } catch (error) {
    Message.error('获取课程列表失败');
    console.error(error);
  } finally {
    loading.value = false;
  }
};

const loadStudentSchedule = async () => {
  
  loading.value = true;
  try{
      let enrollmentsData
      if(localStorage.getItem('userType') === 'student')
        enrollmentsData = await enrollmentApi.getByStudent(student.value.id);
      else if (localStorage.getItem('userType') === "admin")
        enrollmentsData = await enrollmentApi.getByStudent(selectedStudent.value);
      

      enrollments.value = enrollmentsData;
      // console.log(enrollmentsData)
      // 获取课程安排详情
      const sectionIds = enrollmentsData.map(e => e.sectionId);
      if (sectionIds.length === 0) {
        sections.value = [];
        semesters.value = [];
        return;
      }
    
      // 获取所有课程安排
      const allSections = await sectionApi.getAll();
    
      // 过滤出学生选择的课程安排
      sections.value = allSections.filter(section => sectionIds.includes(section.id));
      
      // 提取唯一的学期列表
      const semesterSet = new Set(sections.value.map(section => section.semester));
      semesters.value = Array.from(semesterSet);
    
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

const filterBySemester = () => {
  // 通过计算属性自动筛选
};// 导出处理函数
const handleExport = async ({ key }) => {
  if (key === 'pdf') {
    await exportToPDF();
  } else {
    exportToExcel();
  }
};

// 定义时间段映射表
const timeSlots = [
  "08:00-09:40",
  "10:00-11:40",
  "13:30-15:10",
  "15:30-17:10",
  "18:30-20:10"
];
// 导出PDF
const exportToPDF = () => {
  const element = document.querySelector('.time-table-container'); // 确保 TimeTable 有这个类名
  const opt = {
    margin: 0.5,
    filename: `课表_${student.value.name}_${selectedSemester.value}.pdf`,
    image: { type: 'jpeg', quality: 0.98 },
    html2canvas: { scale: 2 },
    jsPDF: { unit: 'in', format: 'a4', orientation: 'portrait' }
  };

  html2pdf().set(opt).from(element).save();
};

const exportToExcel = () => {
  try {
    const days = ["周一", "周二", "周三", "周四", "周五", "周六", "周天"];
    const timeMap = {};
    
    // 构建二维数组（行：时间段，列：星期）
    const sheetData = timeSlots.map(time => {
      const row = {};
      days.forEach(day => {
        const section = filteredSections.value.find(
          sec => sec.dayOfWeek === day && sec.startTime === time.split("-")[0]
        );
        row[day] = section ? `${section.courseName}\n${section.classroom}\n${section.teacherName}` : "";
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
    XLSX.writeFile(workbook, `课表_${student.value.name}_${selectedSemester.value}.xlsx`);
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
  justify-content: center;
  padding: 40px 0;
}

.empty-container {
  padding: 40px 0;
}

/* 为导出添加特殊样式 */
@media print, pdf {
  .schedule-container {
    padding: 10px;
  }
  .arco-table {
    font-size: 12px;
  }
}
</style>