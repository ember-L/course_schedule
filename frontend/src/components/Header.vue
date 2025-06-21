<template>
  <a-layout-header class="main-header">
    <div class="header-content">
      <div class="header-left">
        <div class="logo-section">
          <div class="logo-icon">
            <icon-book />
          </div>
          <div class="logo-text">
            <h1 class="logo-title">高校排课系统</h1>
            <p class="logo-subtitle">智能排课管理系统</p>
          </div>
        </div>
      </div>
      <div class="header-center">
        <a-menu
          mode="horizontal"
          :selected-keys="[activeKey]"
          @menu-item-click="handleMenuClick"
          class="main-menu"
        >
          <a-menu-item key="/" class="menu-item">
            <icon-home />
            <span>首页</span>
          </a-menu-item>
          <a-menu-item key="/courses" v-if="isAdmin || isTeacher" class="menu-item">
            <icon-book />
            <span>课程管理</span>
          </a-menu-item>
          <a-menu-item key="/teachers" v-if="isAdmin" class="menu-item">
            <icon-user />
            <span>教师管理</span>
          </a-menu-item>
          <a-menu-item key="/classrooms" v-if="isAdmin" class="menu-item">
            <icon-home />
            <span>教室管理</span>
          </a-menu-item>
          <a-menu-item key="/sections" v-if="isAdmin" class="menu-item">
            <icon-calendar />
            <span>排课管理</span>
          </a-menu-item>
          <a-menu-item key="/students" v-if="isAdmin" class="menu-item">
            <icon-user-group />
            <span>学生管理</span>
          </a-menu-item>
          <a-menu-item key="/enrollments" class="menu-item">
            <icon-file />
            <span>选课管理</span>
          </a-menu-item>
          <a-menu-item key="/schedule" class="menu-item">
            <icon-clock-circle />
            <span>课表查询</span>
          </a-menu-item>
        </a-menu>
      </div>
      <div class="header-right">
        <div class="user-info">
          <a-dropdown trigger="click" class="user-dropdown">
            <div class="user-avatar">
              <div class="avatar-icon">
                <icon-user />
              </div>
              <div class="user-details">
                <div class="user-name">{{ userInfo.name || '用户' }}</div>
                <div class="user-type">{{ userTypeText }}</div>
              </div>
              <icon-down class="dropdown-icon" />
            </div>
            <template #content>
              <a-doption class="dropdown-item">
                <div class="dropdown-item-content">
                  <icon-user class="dropdown-icon" />
                  <span>{{ userTypeText }}</span>
                </div>
              </a-doption>
              <a-doption @click="$router.push('/change-password')" class="dropdown-item">
                <div class="dropdown-item-content">
                  <icon-lock class="dropdown-icon" />
                  <span>修改密码</span>
                </div>
              </a-doption>
              <a-doption @click="handleLogout" class="dropdown-item">
                <div class="dropdown-item-content">
                  <icon-export class="dropdown-icon" />
                  <span>退出登录</span>
                </div>
              </a-doption>
            </template>
          </a-dropdown>
        </div>
      </div>
    </div>
  </a-layout-header>
</template>

<script setup>
import { computed } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { Message } from '@arco-design/web-vue';
import { useAuthStore } from '@/store/authStore';

const router = useRouter();
const route = useRoute();
const authStore = useAuthStore();
const activeKey = computed(() => route.path);

const isAdmin = computed(() => authStore.isAdmin);
const isTeacher = computed(() => authStore.isTeacher);
const isStudent = computed(() => authStore.isStudent);
const userInfo = computed(() => authStore.currentUser || {});
const userTypeText = computed(() => {
  if (isAdmin.value) return '管理员';
  if (isTeacher.value) return '教师';
  if (isStudent.value) return '学生';
  return '用户';
});

const handleMenuClick = (key) => {
  router.push(key);
};

const handleLogout = async () => {
  try {
    await authStore.logout();
    Message.success('已退出登录');
    router.push('/login');
  } catch (error) {
    Message.error('退出登录失败');
  }
};
</script>

<style scoped>
.main-header {
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--primary-hover) 100%);
  box-shadow: var(--shadow-2);
  padding: 0;
  height: 64px;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1000;
}
.header-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 100%;
  padding: 0 var(--spacing-xl);
  max-width: 1400px;
  margin: 0 auto;
}

.header-left {
  display: flex;
  align-items: center;
}
.logo-section {
  display: flex;
  align-items: center;
  gap: 16px;
}
.logo-icon {
  width: 40px;
  height: 40px;
  background: rgba(255,255,255,0.2);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  color: white;
  backdrop-filter: blur(10px);
}
.logo-text {
  display: flex;
  flex-direction: column;
}
.logo-title {
  font-size: 18px;
  font-weight: 600;
  color: white;
  margin: 0;
  line-height: 1.2;
}
.logo-subtitle {
  font-size: 12px;
  color: rgba(255,255,255,0.8);
  margin: 0;
  line-height: 1.2;
}
.header-center {
  flex: 1;
  display: flex;
  justify-content: center;
}
.main-menu {
background: transparent !important;
  border: none !important;
  color: white !important;
}
.menu-item {
  display: flex;
  align-items: center;
  gap: var(--spacing-xs);
  padding: 0 var(--spacing-lg);
  height: 64px;
  border-radius: 0;
  transition: all 0.2s ease;
  font-weight: 500;
}
.menu-item .arco-icon {
  font-size: 18px !important;
}
.menu-item:hover {
  background-color: rgba(255, 255, 255, 0.1);
  color: white;
}
.menu-item.arco-menu-item-selected {
  background-color: rgba(255, 255, 255, 0.2);
  color: white;
  border-bottom: 3px solid white;
}
.header-right {
  display: flex;
  align-items: center;
}
.user-dropdown {
  cursor: pointer;
}
.user-avatar {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  padding: var(--spacing-sm) var(--spacing-md);
  background: rgba(255, 255, 255, 0.1);
  border-radius: var(--border-radius-large);
  transition: all 0.2s ease;
  backdrop-filter: blur(10px);
}
.user-avatar:hover {
  background: rgba(255,255,255,0.2);
}
.avatar-icon {
  width: 32px;
  height: 32px;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
  color: white;
}
.user-details {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}
.user-name {
  font-size: 14px;
  font-weight: 500;
  color: white;
  line-height: 1.2;
}
.user-type {
  font-size: 12px;
  color: rgba(255,255,255,0.8);
  line-height: 1.2;
}
.dropdown-icon {
  color: rgba(255,255,255,0.8);
  transition: all 0.2s ease;
}
.user-avatar:hover .dropdown-icon {
  color: white;
}
.dropdown-item {
  padding: 8px 16px;
}
.dropdown-item-content {
  display: flex;
  align-items: center;
  gap: 8px;
}
.dropdown-item-content .dropdown-icon {
  color: #86909c;
}
@media (max-width: 1200px) {
  .header-content {
    padding: 0 24px;
  }
  .main-menu .menu-item {
    padding: 0 16px !important;
  }
}
@media (max-width: 768px) {
  .header-content {
    padding: 0 16px;
  }
  .logo-text {
    display: none;
  }
  .header-center {
    flex: 1;
    display: flex;
    justify-content: center;
    overflow-x: auto;
    -webkit-overflow-scrolling: touch;
  }
  .main-menu {
    min-width: auto !important;
    flex-wrap: nowrap !important;
    overflow-x: auto !important;
    -webkit-overflow-scrolling: touch !important;
  }
  .main-menu .menu-item {
    padding: 0 12px !important;
    font-size: 13px !important;
    white-space: nowrap !important;
    flex-shrink: 0 !important;
  }
  .main-menu .arco-icon {
    font-size: 16px !important;
  }
  .user-details {
    display: none;
  }
}
</style> 