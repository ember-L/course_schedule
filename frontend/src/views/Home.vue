<template>
  <div class="dashboard-container">
    <!-- 日期和欢迎信息 -->
    <a-row :gutter="16" class="welcome-section">
      <a-col :span="16">
        <a-card>
          <div class="welcome-content">
            <div class="date-info">
              <div class="current-date">{{ formattedDate }}</div>
              <div class="day-of-week">{{ dayOfWeek }}</div>
            </div>
            <div class="welcome-text">
              <h2>欢迎回来，{{ userInfo.name }}</h2>
              <p>{{ welcomeMessage }}</p>
            </div>
          </div>
        </a-card>
      </a-col>
      <a-col :span="8">
        <a-card title="系统概览">
          <a-statistic title="学期" :value="currentSemester" />
          <a-divider style="margin: 16px 0" />
          <a-statistic title="周次" :value="currentWeek" />
        </a-card>
      </a-col>
    </a-row>

    <!-- 当日课表 -->
    <a-card title="今日课表" class="today-schedule">
      <a-spin :loading="loading">
        <template v-if="todayClasses.length > 0">
          <a-timeline>
            <a-timeline-item 
              v-for="course in todayClasses" 
              :key="course.id"
              :dot-color="getTimeStatus(course.startTime, course.endTime)"
            >
              <div class="timeline-title">
                <span class="course-time">{{ course.startTime }} - {{ course.endTime }}</span>
                <span class="course-name">{{ course.courseName }}</span>
              </div>
              <div class="timeline-content">
                <p><icon-user /> 教师: {{ course.teacherName }}</p>
                <p><icon-location /> 教室: {{ course.classroom }}</p>
              </div>
            </a-timeline-item>
          </a-timeline>
        </template>
        <a-empty v-else description="今日没有课程安排" />
      </a-spin>
    </a-card>

    <!-- 通知公告 -->
    <a-card title="通知公告" class="announcements">
      <a-list>
        <a-list-item v-for="(notice, index) in announcements" :key="index">
          <a-list-item-meta :title="notice.title">
            <template #description>
              <span>{{ notice.date }}</span>
            </template>
          </a-list-item-meta>
        </a-list-item>
      </a-list>
    </a-card>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { useAuthStore } from '../store/authStore';
import axios from 'axios';

const authStore = useAuthStore();
const userInfo = computed(() => authStore.currentUser || {});
const loading = ref(true);
const todayClasses = ref([]);
const currentSemester = ref('2023-2024-2');
const currentWeek = ref(12);

// 日期相关
const now = new Date();
const formattedDate = computed(() => {
  return now.toLocaleDateString('zh-CN', { year: 'numeric', month: 'long', day: 'numeric' });
});

const dayOfWeek = computed(() => {
  const days = ['星期日', '星期一', '星期二', '星期三', '星期四', '星期五', '星期六'];
  return days[now.getDay()];
});

const welcomeMessage = computed(() => {
  const hour = now.getHours();
  if (hour < 6) return '夜深了，请注意休息！';
  if (hour < 9) return '早上好，祝您有个美好的一天！';
  if (hour < 12) return '上午好，工作顺利！';
  if (hour < 14) return '中午好，记得午休哦！';
  if (hour < 18) return '下午好，继续加油！';
  if (hour < 22) return '晚上好，今天过得如何？';
  return '夜深了，请注意休息！';
});

// 模拟数据 - 实际应用中应从API获取
const announcements = ref([
  { title: '关于2023-2024学年第二学期选课的通知', date: '2024-02-15' },
  { title: '期中考试安排', date: '2024-03-20' },
  { title: '关于五一假期调课安排', date: '2024-04-25' }
]);

// 获取当日课表
const fetchTodaySchedule = async () => {
  loading.value = true;
  try {
    // 实际应用中应调用API
    // const response = await axios.get('/api/schedule/today');
    // todayClasses.value = response.data;
    
    // 模拟数据
    setTimeout(() => {
      const dayIndex = now.getDay();
      if (dayIndex === 0 || dayIndex === 6) {
        todayClasses.value = [];
      } else {
        todayClasses.value = [
          {
            id: 1,
            courseName: '高等数学',
            teacherName: '张教授',
            classroom: '教学楼A-101',
            startTime: '08:00',
            endTime: '09:40'
          },
          {
            id: 2,
            courseName: '数据结构',
            teacherName: '李副教授',
            classroom: '教学楼B-203',
            startTime: '10:00',
            endTime: '11:40'
          },
          {
            id: 3,
            courseName: '计算机网络',
            teacherName: '王讲师',
            classroom: '实验楼C-305',
            startTime: '14:00',
            endTime: '15:40'
          }
        ];
      }
      loading.value = false;
    }, 500);
  } catch (error) {
    console.error('获取今日课表失败:', error);
    loading.value = false;
  }
};

// 根据时间判断课程状态（未开始/进行中/已结束）
const getTimeStatus = (startTime, endTime) => {
  const now = new Date();
  const currentHour = now.getHours();
  const currentMinute = now.getMinutes();
  const currentTimeInMinutes = currentHour * 60 + currentMinute;
  
  const [startHour, startMinute] = startTime.split(':').map(Number);
  const startTimeInMinutes = startHour * 60 + startMinute;
  
  const [endHour, endMinute] = endTime.split(':').map(Number);
  const endTimeInMinutes = endHour * 60 + endMinute;
  
  if (currentTimeInMinutes < startTimeInMinutes) {
    return 'blue'; // 未开始
  } else if (currentTimeInMinutes > endTimeInMinutes) {
    return 'gray'; // 已结束
  } else {
    return 'green'; // 进行中
  }
};

onMounted(() => {
  fetchTodaySchedule();
});
</script>

<style scoped>
.dashboard-container {
  padding: 20px;
  max-width: 1200px;
  margin: 0 auto;
}

.welcome-section {
  margin-bottom: 20px;
}

.welcome-content {
  display: flex;
  align-items: center;
}

.date-info {
  margin-right: 24px;
  text-align: center;
  padding-right: 24px;
  border-right: 1px solid #e5e5e5;
}

.current-date {
  font-size: 16px;
  color: #333;
}

.day-of-week {
  font-size: 24px;
  font-weight: bold;
  color: #165dff;
  margin-top: 4px;
}

.welcome-text h2 {
  margin-top: 0;
  margin-bottom: 8px;
}

.welcome-text p {
  margin: 0;
  color: #666;
}

.today-schedule {
  margin-bottom: 20px;
}

.timeline-title {
  display: flex;
  align-items: center;
  margin-bottom: 8px;
}

.course-time {
  font-weight: bold;
  margin-right: 12px;
  color: #666;
}

.course-name {
  font-size: 16px;
  font-weight: bold;
}

.timeline-content {
  color: #666;
}

.timeline-content p {
  margin: 4px 0;
}

.announcements {
  margin-bottom: 20px;
}
</style>
