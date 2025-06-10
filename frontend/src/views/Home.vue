<template>
  <div class="home-container">
    <a-row :gutter="16">
      <a-col :span="24">
        <a-card class="welcome-card">
          <template #title>高校排课系统</template>
          <template #extra>
            <a-space>
              <a-button type="primary" @click="$router.push('/courses')">课程管理</a-button>
              <a-button type="primary" @click="$router.push('/sections')">排课管理</a-button>
              <a-button type="primary" @click="$router.push('/schedule')">课表查询</a-button>
            </a-space>
          </template>
          <p>欢迎使用高校排课系统，本系统提供课程管理、教师管理、教室管理、排课管理和学生选课等功能。</p>
          <a-divider />
          <a-row :gutter="16">
            <a-col :span="8">
              <a-statistic title="课程数量" :value="statistics.courses" />
            </a-col>
            <a-col :span="8">
              <a-statistic title="教师数量" :value="statistics.teachers" />
            </a-col>
            <a-col :span="8">
              <a-statistic title="学生数量" :value="statistics.students" />
            </a-col>
          </a-row>
        </a-card>
      </a-col>
    </a-row>
    
    <a-row :gutter="16" style="margin-top: 16px">
      <a-col :span="8" v-for="menu in menus" :key="menu.path">
        <a-card hoverable @click="$router.push(menu.path)">
          <template #title>{{ menu.title }}</template>
          <p>{{ menu.description }}</p>
        </a-card>
      </a-col>
    </a-row>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { courseApi, teacherApi, studentApi } from '../api/modules';

const statistics = ref({
  courses: 0,
  teachers: 0,
  students: 0
});

const menus = [
  {
    title: '课程管理',
    path: '/courses',
    description: '添加、编辑、删除课程信息'
  },
  {
    title: '教师管理',
    path: '/teachers',
    description: '添加、编辑、删除教师信息'
  },
  {
    title: '教室管理',
    path: '/classrooms',
    description: '添加、编辑、删除教室信息'
  },
  {
    title: '排课管理',
    path: '/sections',
    description: '安排课程、教师、教室、时间段'
  },
  {
    title: '学生管理',
    path: '/students',
    description: '添加、编辑、删除学生信息'
  },
  {
    title: '选课管理',
    path: '/enrollments',
    description: '学生选课、查看课表'
  },
  {
    title: '课表查询',
    path: '/schedule',
    description: '查看学生个人课表'
  }
];

onMounted(async () => {
  try {
    // 并行获取统计数据
    const [courses, teachers, students] = await Promise.all([
      courseApi.getAll(),
      teacherApi.getAll(),
      studentApi.getAll()
    ]);
    
    statistics.value = {
      courses: courses.length,
      teachers: teachers.length,
      students: students.length
    };
  } catch (error) {
    console.error('获取统计数据失败', error);
  }
});
</script>

<style scoped>
.home-container {
  padding: 20px;
}
.welcome-card {
  margin-bottom: 20px;
}
</style>