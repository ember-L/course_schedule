<template>
  <a-config-provider>
    <div class="app-container">
      <a-layout v-if="isAuthenticated">
        <a-layout-header>
          <div class="header-content">
            <div class="logo">高校排课系统</div>
            <a-menu
              mode="horizontal"
              :selected-keys="[activeKey]"
              @menu-item-click="handleMenuClick"
            >
              <a-menu-item key="/">首页</a-menu-item>
              <a-menu-item key="/courses" v-if="isAdmin || isTeacher">课程管理</a-menu-item>
              <a-menu-item key="/teachers" v-if="isAdmin">教师管理</a-menu-item>
              <a-menu-item key="/classrooms" v-if="isAdmin">教室管理</a-menu-item>
              <a-menu-item key="/sections" v-if="isAdmin || isTeacher">排课管理</a-menu-item>
              <a-menu-item key="/students" v-if="isAdmin">学生管理</a-menu-item>
              <a-menu-item key="/enrollments">选课管理</a-menu-item>
              <a-menu-item key="/schedule">课表查询</a-menu-item>
              <a-menu-item key="/api-test" v-if="isAdmin">API测试</a-menu-item>
            </a-menu>
            <div class="user-info">
              <a-dropdown trigger="click">
                <a-button type="text">
                  {{ userInfo.name || '用户' }}
                  <icon-down />
                </a-button>
                <template #content>
                  <a-doption>
                    <a-space>
                      <icon-user />
                      <span>{{ userTypeText }}</span>
                    </a-space>
                  </a-doption>
                  <a-doption @click="$router.push('/change-password')">
                    <a-space>
                      <icon-lock />
                      <span>修改密码</span>
                    </a-space>
                  </a-doption>
                  <a-doption @click="handleLogout">
                    <a-space>
                      <icon-export />
                      <span>退出登录</span>
                    </a-space>
                  </a-doption>
                </template>
              </a-dropdown>
            </div>
          </div>
        </a-layout-header>
        
        <a-layout-content>
          <router-view />
        </a-layout-content>
        
        <a-layout-footer>
          <div class="footer-content">
            高校排课系统 &copy; {{ new Date().getFullYear() }}
          </div>
        </a-layout-footer>
      </a-layout>
      
      <router-view v-else />
    </div>
  </a-config-provider>
</template>

<script setup>
import { ref, watch, computed, onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { Message } from '@arco-design/web-vue';
import { useAuthStore } from './store/authStore';

const router = useRouter();
const route = useRoute();
const authStore = useAuthStore();
const activeKey = ref(route.path);

// 计算属性
const isAuthenticated = computed(() => authStore.isAuthenticated);
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

// 监听路由变化，更新激活的菜单项
watch(() => route.path, (newPath) => {
  activeKey.value = newPath;
});

// 初始化认证状态
onMounted(() => {
  authStore.init();
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

<style>
body {
  margin: 0;
  padding: 0;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
}

.app-container {
  min-height: 100vh;
}

.header-content {
  display: flex;
  align-items: center;
  color: #fff;
  padding: 0 20px;
  width: 100%;
}

.logo {
  font-size: 18px;
  font-weight: bold;
  margin-right: 40px;
}

.user-info {
  margin-left: auto;
  color: #fff;
}

.footer-content {
  text-align: center;
  color: #666;
  padding: 20px 0;
}
</style>