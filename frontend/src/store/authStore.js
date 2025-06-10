import { defineStore } from 'pinia';
import { authApi } from '../api/auth';

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null,
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
      }
    },
    
    async fetchCurrentUser() {
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
        
        this.user = response;
        return response;
      } catch (error) {
        console.error('获取用户信息失败', error);
        // 如果获取用户信息失败，可能是token过期，执行登出操作
        this.logout();
        return null;
      } finally {
        this.loading = false;
      }
    },
    
    // 初始化认证状态
    init() {
      if (this.token) {
        // 获取当前用户信息
        this.fetchCurrentUser();
      }
    }
  }
});