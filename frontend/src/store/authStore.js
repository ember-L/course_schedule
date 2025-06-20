import { defineStore } from 'pinia';
import { authApi } from '@/api/auth';

// 安全解析 localStorage
const safeParse = (key, defaultValue = null) => {
  try {
    const item = localStorage.getItem(key);
    console.log(`解析 ${key}`, item);
    return item ? JSON.parse(item) : defaultValue;
  } catch (error) {
    console.error(`解析 ${key} 失败`, error);
    return defaultValue;
  }
};

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: safeParse('user'), 
    token: localStorage.getItem('token') || null,
    userType: localStorage.getItem('userType') || null,
    loading: false
  }),
  
  getters: {
    isAuthenticated: (state) => !!state.token,
    isAdmin: (state) => state.userType === 'admin',
    isTeacher: (state) => state.userType === 'teacher',
    isStudent: (state) => state.userType === 'student',
    currentUser: (state) => state.user
  },
  
  actions: {
    async login(userType, username, password) {
      this.loading = true;
      try {
        let response;
        
        // 根据用户类型调用不同的登录API
        if (userType === 'admin') {
          response = await authApi.adminLogin(username, password);
        } else if (userType === 'teacher') {
          response = await authApi.teacherLogin(username, password);
        } else {
          response = await authApi.studentLogin(username, password);
        }
        
        // 保存认证信息
        this.token = response.token;
        this.userType = userType;
        this.user = response.user;
        
        // 存储到本地存储
        localStorage.setItem('token', response.token);
        localStorage.setItem('userType', userType);
        localStorage.setItem('user', JSON.stringify(response.user));
        return response;
      } finally {
        this.loading = false;
      }
    },
    
    async logout() {
      try {
        // 调用登出API
        await authApi.logout();
      } catch (error) {
        console.error('登出失败', error);
      } finally {
        // 清除认证信息
        this.token = null;
        this.userType = null;
        this.user = null;
        
        // 清除本地存储
        localStorage.removeItem('token');
        localStorage.removeItem('userType');
        localStorage.removeItem('user');
      }
    },
    
    async fetchCurrentUser(autoLogout = true) {
      if (!this.token) return null;
      
      this.loading = true;
      try {
        let response;
         // 确保用户 ID 存在
        if (!this.user?.id) {
          throw new Error('用户 ID 不存在');
        }

        const userId = this.user.id;
        console.log('获取当前用户信息', userId);
        // 根据用户类型获取当前用户信息
        if (this.userType === 'admin') {
          response = await authApi.getAdminProfile(userId);
        } else if (this.userType === 'teacher') {
          response = await authApi.getTeacherProfile(userId);
        } else {
          response = await authApi.getStudentProfile(userId);
        }
        
        // 更新用户信息
        if (response?.user) {
          this.user = response.user;
          localStorage.setItem('user', JSON.stringify(response.user));
        }

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
      }
    },
    
    // 初始化认证状态
    async init() {
      if (this.token) {
        // 优先用本地 user 信息
        if (!this.user) {
          this.user = safeParse('user');
        }
        
        // 异步拉取最新用户信息，但不影响初始化状态
        try {
          await this.fetchCurrentUser(false); // 不自动登出
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