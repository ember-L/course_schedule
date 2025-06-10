import api from '@/api/index';

// 认证相关API
export const authApi = {
  // 管理员登录
  adminLogin: (username, password) => 
    api.post('/admin/login', { username, password }),
  
  // 教师登录
  teacherLogin: (username, password) => 
    api.post('/teacher/login', { username, password }),
  
  // 学生登录
  studentLogin: (username, password) => 
    api.post('/student/login', { username, password }),
  
  // 登出
  logout: () => api.post('/logout'),
  
  // 获取管理员信息
  getAdminProfile: () => api.get('/admin/profile'),
  
  // 获取教师信息
  getTeacherProfile: () => api.get('/teacher/profile'),
  
  // 获取学生信息
  getStudentProfile: () => api.get('/student/profile'),
  
  // 修改密码
  changePassword: (oldPassword, newPassword) => 
    api.post('/change-password', { oldPassword, newPassword })
};