<template>
  <div class="page-container fade-in">
    <!-- 页面头部 -->
    <div class="page-header">
      <h1 class="page-title">欢迎回来，{{ userInfo.name }}</h1>
      <p class="page-subtitle">{{ welcomeMessage }}</p>
      <a-button v-if="isAdmin" type="primary" size="small" style="margin-top: 8px;" @click="editConfig">
        <icon-settings style="margin-right:4px"/>学期/周次管理
      </a-button>
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
            <div class="statistic-value">
              {{ currentSemester }}
              <a-button v-if="isAdmin" type="text" size="mini" @click="editConfig" style="margin-left: 4px;">
                <icon-edit />
              </a-button>
            </div>
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

    <!-- 学期/周次编辑弹窗 -->
    <a-modal v-model:visible="editConfigVisible" title="编辑学期与周次" @ok="handleConfigSave" :confirm-loading="configLoading" @cancel="editConfigVisible=false">
      <a-form :model="{editSemester,editStartDate}">
        <a-form-item label="学期">
          <a-input v-model="editSemester" placeholder="如2023-2024-2" />
        </a-form-item>
        <a-form-item label="学期起始日期">
          <a-date-picker v-model="editStartDate" format="YYYY-MM-DD" style="width:100%" />
        </a-form-item>
        <a-form-item label="当前周次">
          <a-input v-model="currentWeek" readonly />
        </a-form-item>
      </a-form>
    </a-modal>

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
          <a-button v-if="isAdmin" type="primary" size="small" @click="newAnnVisible=true">
            <icon-plus /> 发布公告
          </a-button>
        </template>
        
        <div class="announcements-list">
          <div 
            v-for="notice in announcements" 
            :key="notice.id"
            class="announcement-item"
            @click="showAnnDetail(notice)"
            style="cursor:pointer;"
          >
            <div class="announcement-content">
              <div class="announcement-title">{{ notice.title }}</div>
              <div class="announcement-date">{{ notice.date }}</div>
            </div>
            <div class="announcement-badge">
              <a-tag size="small" color="blue">新</a-tag>
              <a-popconfirm v-if="isAdmin" title="确定删除该公告？" @ok="() => handleAnnDelete(notice.id)">
                <a-button type="text" size="mini" status="danger"><icon-delete /></a-button>
              </a-popconfirm>
            </div>
          </div>
        </div>
      </a-card>
      <!-- 发布公告弹窗 -->
      <a-modal v-model:visible="newAnnVisible" title="发布公告" @ok="handleAnnCreate" :confirm-loading="annLoading" @cancel="newAnnVisible=false">
        <a-form :model="{newAnnTitle,newAnnContent}">
          <a-form-item label="标题">
            <a-input v-model="newAnnTitle" placeholder="请输入公告标题" />
          </a-form-item>
          <a-form-item label="内容">
            <a-textarea v-model="newAnnContent" placeholder="请输入公告内容" :auto-size="{minRows:3,maxRows:6}" />
          </a-form-item>
        </a-form>
      </a-modal>
      <!-- 公告详情弹窗 -->
      <a-modal v-model:visible="annDetailVisible" :title="annDetail.title" footer="null" @cancel="annDetailVisible=false">
        <div style="white-space:pre-line;min-height:60px;">{{ annDetail.content }}</div>
        <div style="margin-top:16px;color:#888;font-size:13px;">发布时间：{{ annDetail.date }}</div>
      </a-modal>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { useAuthStore } from '@/store/authStore';
import {scheduleApi} from '@/api/modules.js'
import { systemApi, announcementApi } from '@/api/modules.js'


const authStore = useAuthStore();
const userInfo = computed(() => authStore.currentUser || {});
const loading = ref(true);
const todayClasses = ref([]);
const currentSemester = ref('');
const currentWeek = ref(1);
const announcements = ref([]);
const configLoading = ref(false);
const annLoading = ref(false);

// 计算属性
const isAdmin = computed(() => authStore.isAdmin);
const isTeacher = computed(() => authStore.isTeacher);
const isStudent = computed(() => authStore.isStudent);
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

// 管理员编辑面板
const editConfigVisible = ref(false);
const editSemester = ref('');
const editStartDate = ref('');
const editConfig = () => {
  editSemester.value = currentSemester.value;
  editStartDate.value = startDate.value;
  editConfigVisible.value = true;
};
const handleConfigSave = async () => {
  configLoading.value = true;
  try {
    await systemApi.updateConfig({ semester: editSemester.value, startDate: editStartDate.value });
    currentSemester.value = editSemester.value;
    startDate.value = editStartDate.value;
    editConfigVisible.value = false;
    await fetchConfig();
  } catch (e) {
    window.$message?.error('更新失败');
  } finally {
    configLoading.value = false;
  }
};

// 公告管理
const newAnnVisible = ref(false);
const newAnnTitle = ref('');
const newAnnContent = ref('');
const handleAnnCreate = async () => {
  annLoading.value = true;
  try {
    await announcementApi.create({ title: newAnnTitle.value, content: newAnnContent.value });
    await fetchAnnouncements();
    newAnnVisible.value = false;
    newAnnTitle.value = '';
    newAnnContent.value = '';
  } catch (e) {
    window.$message?.error('发布失败');
  } finally {
    annLoading.value = false;
  }
};
const handleAnnDelete = async (id) => {
  annLoading.value = true;
  try {
    await announcementApi.delete(id);
    await fetchAnnouncements();
  } catch (e) {
    window.$message?.error('删除失败');
  } finally {
    annLoading.value = false;
  }
};

// 获取系统参数
const fetchConfig = async () => {
  try {
    const res = await systemApi.getConfig();
    currentSemester.value = res.semester;
    startDate.value = res.startDate;
    currentWeek.value = res.currentWeek;
  } catch (e) {
    currentSemester.value = '未知';
    startDate.value = '';
    currentWeek.value = 1;
  }
};
// 获取公告
const fetchAnnouncements = async () => {
  try {
    const res = await announcementApi.getAll();
    announcements.value = (res || []).map(a => ({ ...a, date: a.createdAt?.slice(0,10) }));
  } catch (e) {
    announcements.value = [];
  }
};

// 获取当日课表
const fetchTodaySchedule = async () => {
  loading.value = true;
  try {
    if(!isAdmin.value){
      const response = await scheduleApi.getTodaySchedule();
      // 兼容后端返回 Section 结构
      if (Array.isArray(response)) {
        todayClasses.value = response.map(section => ({
          id: section.id,
          courseName: section.course?.name || '未知课程',
          teacherName: section.teacher?.name || '未知教师',
          classroom: (section.classroom?.building || '') + (section.classroom?.room || ''),
          startTime: section.timeSlot?.startTime || '',
          endTime: section.timeSlot?.endTime || ''
        }));
      } else {
        todayClasses.value = [];
      }
    }
  } catch (error) {
    console.error('获取今日课表失败:', error);
    todayClasses.value = [];
  } finally {
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

const startDate = ref('');
const annDetailVisible = ref(false);
const annDetail = ref({});
const showAnnDetail = (notice) => {
  annDetail.value = notice;
  annDetailVisible.value = true;
};

onMounted(() => {
  fetchConfig();
  fetchAnnouncements();
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
