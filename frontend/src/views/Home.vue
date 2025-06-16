<template>
  <div class="page-container fade-in">
    <!-- 页面头部 -->
    <div class="page-header">
      <h1 class="page-title">欢迎回来，{{ userInfo.name }}</h1>
      <p class="page-subtitle">{{ welcomeMessage }}</p>
    </div>

    <!-- 统计卡片 -->
    <div class="card-grid">
      <div class="statistic-card">
        <div class="statistic-content">
          <div class="statistic-icon">
            <icon-calendar />
          </div>
          <div class="statistic-info">
            <div class="statistic-value">{{ formattedDate }}</div>
            <div class="statistic-label">{{ dayOfWeek }}</div>
          </div>
        </div>
      </div>
      
      <div class="statistic-card">
        <div class="statistic-content">
          <div class="statistic-icon">
            <icon-book />
          </div>
          <div class="statistic-info">
            <div class="statistic-value">{{ currentSemester }}</div>
            <div class="statistic-label">当前学期</div>
          </div>
        </div>
      </div>
      
      <div class="statistic-card">
        <div class="statistic-content">
          <div class="statistic-icon">
            <icon-clock-circle />
          </div>
          <div class="statistic-info">
            <div class="statistic-value">第{{ currentWeek }}周</div>
            <div class="statistic-label">当前周次</div>
          </div>
        </div>
      </div>
    </div>

    <!-- 主要内容区域 -->
    <div class="main-content">
      <!-- 当日课表 -->
      <a-card class="schedule-card" title="今日课表">
        <template #extra>
          <a-tag :color="todayClasses.length > 0 ? 'green' : 'orange'">
            {{ todayClasses.length > 0 ? `${todayClasses.length} 门课程` : '今日无课' }}
          </a-tag>
        </template>
        
        <a-spin :loading="loading">
          <template v-if="todayClasses.length > 0">
            <div class="timeline-container">
              <a-timeline>
                <a-timeline-item 
                  v-for="course in todayClasses" 
                  :key="course.id"
                  :dot-color="getTimeStatus(course.startTime, course.endTime)"
                >
                  <div class="timeline-content">
                    <div class="timeline-header">
                      <span class="course-time">{{ course.startTime }} - {{ course.endTime }}</span>
                      <span class="course-name">{{ course.courseName }}</span>
                    </div>
                    <div class="timeline-details">
                      <div class="detail-item">
                        <icon-user class="detail-icon" />
                        <span>{{ course.teacherName }}</span>
                      </div>
                      <div class="detail-item">
                        <icon-location class="detail-icon" />
                        <span>{{ course.classroom }}</span>
                      </div>
                    </div>
                  </div>
                </a-timeline-item>
              </a-timeline>
            </div>
          </template>
          <a-empty v-else description="今日没有课程安排，好好休息一下吧！" />
        </a-spin>
      </a-card>

      <!-- 通知公告 -->
      <a-card class="announcements-card" title="通知公告">
        <template #extra>
          <a-button type="text" size="small">
            <icon-more />
          </a-button>
        </template>
        
        <div class="announcements-list">
          <div 
            v-for="(notice, index) in announcements" 
            :key="index"
            class="announcement-item"
          >
            <div class="announcement-content">
              <div class="announcement-title">{{ notice.title }}</div>
              <div class="announcement-date">{{ notice.date }}</div>
            </div>
            <div class="announcement-badge">
              <a-tag size="small" color="blue">新</a-tag>
            </div>
          </div>
        </div>
      </a-card>
    </div>
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
.main-content {
  display: grid;
  grid-template-columns: 2fr 1fr;
  gap: var(--spacing-xl);
  margin-top: var(--spacing-xl);
}

.schedule-card {
  height: fit-content;
}

.announcements-card {
  height: fit-content;
}

.timeline-container {
  padding: var(--spacing-md) 0;
}

.timeline-content {
  padding: var(--spacing-sm) 0;
}

.timeline-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-sm);
}

.course-time {
  font-size: 14px;
  color: var(--text-secondary);
  font-weight: 500;
}

.course-name {
  font-size: 16px;
  font-weight: 600;
  color: var(--text-primary);
}

.timeline-details {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-xs);
}

.detail-item {
  display: flex;
  align-items: center;
  gap: var(--spacing-xs);
  color: var(--text-secondary);
  font-size: 14px;
}

.detail-icon {
  color: var(--primary-color);
}

.announcements-list {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
}

.announcement-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--spacing-md);
  background-color: var(--bg-tertiary);
  border-radius: var(--border-radius-medium);
  transition: all 0.2s ease;
}

.announcement-item:hover {
  background-color: var(--primary-light);
  transform: translateX(4px);
}

.announcement-content {
  flex: 1;
}

.announcement-title {
  font-size: 14px;
  font-weight: 500;
  color: var(--text-primary);
  margin-bottom: var(--spacing-xs);
}

.announcement-date {
  font-size: 12px;
  color: var(--text-tertiary);
}

.statistic-content {
  display: flex;
  align-items: center;
  gap: var(--spacing-lg);
}

.statistic-icon {
  font-size: 32px;
  opacity: 0.8;
}

.statistic-info {
  flex: 1;
}

.statistic-value {
  font-size: 20px;
  font-weight: 600;
  margin-bottom: var(--spacing-xs);
}

.statistic-label {
  font-size: 14px;
  opacity: 0.8;
}

/* 响应式设计 */
@media (max-width: 1024px) {
  .main-content {
    grid-template-columns: 1fr;
  }
  
  .card-grid {
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  }
}

@media (max-width: 768px) {
  .page-container {
    padding: var(--spacing-lg);
  }
  
  .card-grid {
    grid-template-columns: 1fr;
  }
  
  .timeline-header {
    flex-direction: column;
    align-items: flex-start;
    gap: var(--spacing-xs);
  }
  
  .statistic-content {
    flex-direction: column;
    text-align: center;
    gap: var(--spacing-md);
  }
}
</style>
