<template>
  <a-config-provider>
    <div class="app-container">
      <div v-if="loading" class="global-loading">
        <div class="loading-content">
          <div class="loading-logo">
            <icon-book />
          </div>
          <div class="loading-text">高校排课系统</div>
          <a-spin size="large" />
        </div>
      </div>
      <template v-else>
        <!-- 登录页面不显示 Header 和 Footer -->
        <template v-if="isLoginPage">
          <router-view />
        </template>
        <!-- 其他页面显示完整的布局 -->
        <template v-else>
          <Header v-if="isAuthenticated" />
          <div class="main-content">
            <router-view />
          </div>
          <Footer v-if="isAuthenticated" />
        </template>
      </template>
    </div>
  </a-config-provider>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { useAuthStore } from './store/authStore';
import Header from './components/Header.vue';
import Footer from './components/Footer.vue';

const route = useRoute();
const authStore = useAuthStore();
const isAuthenticated = computed(() => authStore.isAuthenticated);
const isLoginPage = computed(() => route.name === 'Login');
const loading = ref(true);

onMounted(async () => {
  try {
    await authStore.init();
    if (authStore.token && !authStore.user) {
      await authStore.fetchCurrentUser(false);
    }
  } catch (error) {
    console.error('初始化认证状态失败', error);
  } finally {
    setTimeout(() => {
      loading.value = false;
    }, 300);
  }
});
</script>

<style scoped>
.app-container {
  min-height: 100vh;
  background-color: #f7f8fa;
}
.main-content {
  margin-top: 64px;
  min-height: calc(100vh - 64px - 60px);
  background-color: #f7f8fa;
}
.global-loading {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: linear-gradient(135deg, #165dff 0%, #4080ff 100%);
}
.loading-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 24px;
  color: white;
}
.loading-logo {
  width: 80px;
  height: 80px;
  background: rgba(255,255,255,0.2);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 36px;
  color: white;
  backdrop-filter: blur(10px);
}
.loading-text {
  font-size: 24px;
  font-weight: 600;
  color: white;
}
</style>