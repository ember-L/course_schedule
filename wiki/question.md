这个问题是因为在第一次刷新时，authStore.init() 还没有完成，但页面已经开始渲染了。让我检查一下相关的代码逻辑：
我找到问题了！第一次刷新时导航栏消失的原因是：
初始化时序问题：在 authStore.init() 完成之前，isAuthenticated 计算属性可能返回 false，导致导航栏不显示
token 存在但 user 为空：即使有 token，如果 fetchCurrentUser() 失败，会导致 logout() 被调用，清空所有状态
让我修复这个问题：
        if (!this.user) {
          const localUser = localStorage.getItem('user');
          if (localUser) {
            try {
              this.user = JSON.parse(localUser);
            } catch (error) {
              console.error('解析本地用户信息失败', error);
            }
          }
        }
        
        // 异步拉取最新用户信息，但不影响初始化状态
        try {
          await this.fetchCurrentUser();
        } catch (error) {
          console.error('拉取用户信息失败，使用本地信息', error);
          // 不执行 logout，保持当前状态
        }
        
        return this.user;
      }
      return null;
    }
  }
});
现在我还需要修改 fetchCurrentUser 方法，避免在初始化时自动登出：
      }
    },
    
    async fetchCurrentUser(autoLogout = true) {
      if (!this.token) return null;
      
      this.loading = true;
      try {
        let response;
        
        // 根据用户类型获取当前用户信息
        if (this.userType === 'admin') {
          response = await authApi.getAdminProfile();
        } else if (this.userType === 'teacher') {
          response = await authApi.getTeacherProfile();
        } else {
          response = await authApi.getStudentProfile();
        }
        
        this.user = response.user;
        localStorage.setItem('user', JSON.stringify(response.user));
        return response;
      } catch (error) {
        console.error('获取用户信息失败', error);
        // 只有在明确要求时才执行登出操作
        if (autoLogout) {
          this.logout();
        }
        return null;
      } finally {
        this.loading = false;
现在更新 init 方法，调用 fetchCurrentUser 时不自动登出：
        
        // 异步拉取最新用户信息，但不影响初始化状态
        try {
          await this.fetchCurrentUser(false); // 不自动登出
        } catch (error) {
          console.error('拉取用户信息失败，使用本地信息', error);
          // 不执行 logout，保持当前状态
现在我还需要优化 App.vue 中的初始化逻辑，确保在初始化完成前不显示主布局：
      <template v-else>
        <a-layout v-if="isAuthenticated" class="main-layout">
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
                  <a-menu-item key="/sections" v-if="isAdmin || isTeacher" class="menu-item">
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
          
          <a-layout-content class="main-content">
            <router-view />
          </a-layout-content>
          
          <a-layout-footer class="main-footer">
            <div class="footer-content">
              <div class="footer-left">
                <span>高校排课系统 &copy; {{ new Date().getFullYear() }}</span>
              </div>
              <div class="footer-right">
                <span>智能排课 · 高效管理</span>
              </div>
            </div>
          </a-layout-footer>
        </a-layout>
        <router-view v-else />
      </template>
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

const loading = ref(true);

// 监听路由变化，更新激活的菜单项
watch(() => route.path, (newPath) => {
  activeKey.value = newPath;
});

// 初始化认证状态
onMounted(async () => {
  try {
    await authStore.init();
  } catch (error) {
    console.error('初始化认证状态失败', error);
  } finally {
    // 确保有足够的时间让状态稳定
    setTimeout(() => {
      loading.value = false;
    }, 100);
  }
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
.app-container {
  min-height: 100vh;
  background-color: var(--bg-secondary);
}

.main-layout {
  min-height: 100vh;
}

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
  gap: var(--spacing-md);
}

.logo-icon {
  width: 40px;
  height: 40px;
  background: rgba(255, 255, 255, 0.2);
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
  color: rgba(255, 255, 255, 0.8);
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

.main-menu .arco-menu-item {
  color: white !important;
  background: transparent !important;
  border: none !important;
  margin: 0 !important;
  padding: 0 var(--spacing-lg) !important;
  height: 64px !important;
  line-height: 64px !important;
  border-radius: 0 !important;
  transition: all 0.2s ease !important;
  font-weight: 500 !important;
  font-size: 15px !important;
  display: flex !important;
  align-items: center !important;
  gap: var(--spacing-xs) !important;
}

.main-menu .arco-menu-item:hover {
  background-color: rgba(255, 255, 255, 0.1) !important;
  color: white !important;
}

.main-menu .arco-menu-item-selected {
  background-color: rgba(255, 255, 255, 0.2) !important;
  color: white !important;
  border-bottom: 3px solid white !important;
}

.main-menu .arco-menu-item-selected::after {
  display: none !important;
}

.main-menu .arco-menu-item .arco-icon {
  color: white !important;
  font-size: 18px !important;
}

.main-menu .arco-menu-item:hover .arco-icon {
  color: white !important;
}

.main-menu .arco-menu-item-selected .arco-icon {
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
  font-size: 15px;
  color: white !important;
}

.menu-item:hover {
  background-color: rgba(255, 255, 255, 0.1);
  color: white !important;
}

.menu-item.arco-menu-item-selected {
  background-color: rgba(255, 255, 255, 0.2);
  color: white !important;
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
  background: rgba(255, 255, 255, 0.2);
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
  color: rgba(255, 255, 255, 0.8);
  line-height: 1.2;
}

.dropdown-icon {
  color: rgba(255, 255, 255, 0.8);
  transition: all 0.2s ease;
}

.user-avatar:hover .dropdown-icon {
  color: white;
}

.dropdown-item {
  padding: var(--spacing-sm) var(--spacing-md);
}

.dropdown-item-content {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
}

.dropdown-item-content .dropdown-icon {
  color: var(--text-tertiary);
}

.main-content {
  margin-top: 64px;
  min-height: calc(100vh - 64px - 60px);
  background-color: var(--bg-secondary);
}

.main-footer {
  background: var(--bg-primary);
  border-top: 1px solid var(--border-primary);
  padding: 0;
  height: 60px;
}

.footer-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 100%;
  padding: 0 var(--spacing-xl);
  max-width: 1400px;
  margin: 0 auto;
  color: var(--text-secondary);
  font-size: 14px;
}

.footer-left {
  display: flex;
  align-items: center;
}

.footer-right {
  display: flex;
  align-items: center;
}

.global-loading {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--primary-hover) 100%);
}

.loading-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: var(--spacing-lg);
  color: white;
}

.loading-logo {
  width: 80px;
  height: 80px;
  background: rgba(255, 255, 255, 0.2);
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

/* 响应式设计 */
@media (max-width: 1200px) {
  .header-content {
    padding: 0 var(--spacing-lg);
  }
  
  .menu-item {
    padding: 0 var(--spacing-md);
  }
  
  .footer-content {
    padding: 0 var(--spacing-lg);
  }
}

@media (max-width: 768px) {
  .header-content {
    padding: 0 var(--spacing-md);
  }
  
  .logo-text {
    display: none;
  }
  
  .header-center {
    display: none;
  }
  
  .user-details {
    display: none;
  }
  
  .footer-content {
    flex-direction: column;
    gap: var(--spacing-sm);
    padding: var(--spacing-md);
    text-align: center;
  }
  
  .main-header {
现在问题应该解决了！我做了以下修复：
🔧 修复第一次刷新导航栏消失的问题
问题原因：
初始化时序问题：authStore.init() 在 fetchCurrentUser() 失败时会自动调用 logout()，清空所有状态
状态不稳定：在初始化完成前，isAuthenticated 可能返回错误的值
修复方案：
1. 优化 fetchCurrentUser 方法
添加 autoLogout 参数，控制是否在失败时自动登出
避免在初始化时因为网络问题而清空用户状态
2. 改进 init 方法
优先使用本地存储的用户信息
异步拉取最新信息但不影响初始化状态
即使拉取失败也保持当前状态
3. 优化 App.vue 初始化
添加短暂延时确保状态稳定
避免在状态未完全初始化时就显示主布局
修复效果：
✅ 第一次刷新时导航栏不会消失
✅ 即使网络问题也不会影响已登录状态
✅ 保持用户登录状态的稳定性
✅ 提供更好的用户体验
现在第一次刷新时应该不会再出现导航栏消失的问题了！