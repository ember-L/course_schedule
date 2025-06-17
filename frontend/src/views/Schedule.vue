<template>
  <div class="schedule-container">
    <a-card title="我的课表">
      <template #extra>
        <a-space>
          <a-select
            v-model="selectedSemester"
            placeholder="选择学期"
            style="width: 150px"
            @change="filterBySemester"
          >
            <a-option v-for="sem in semesters" :key="sem" :value="sem">{{ sem }}</a-option>
          </a-select>
        </a-space>
      </template>
      
      <div v-if="loading" class="loading-container">
        <a-spin />
      </div>
      <div v-else-if="!student" class="empty-container">
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

const loading = ref(false);
const student = ref()
const enrollments = ref([]);
const sections = ref([]);
const selectedStudent = ref(null);
const selectedSemester = ref('');
const semesters = ref([]);


// 根据学期筛选课程
const filteredSections = computed(() => {
  if (!selectedSemester.value) {
    return sections.value;
  }
  return sections.value.filter(section => section.semester === selectedSemester.value);
});

onMounted(async () => {
  await loadStudents();
  await loadStudentSchedule();
});

const loadStudents = async () => {
  loading.value = true;
  try {
      const userId = JSON.parse(localStorage.getItem('user')).id
      const data = await studentApi.getById(userId);
      student.value = data;
      console.log(student.value)

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
      console.log(student.value.id)
      const enrollmentsData = await enrollmentApi.getByStudent(1);
      enrollments.value = enrollmentsData;
      console.log(enrollmentsData)
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
</style>