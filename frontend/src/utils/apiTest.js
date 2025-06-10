import api from '../api';
import { authApi } from '../api/auth';

// 测试API连接
export const testApiConnection = async () => {
  try {
    // 测试管理员登录
    console.log('测试管理员登录...');
    const adminLoginResult = await authApi.adminLogin('admin', '123456');
    console.log('管理员登录成功:', adminLoginResult);
    
    // 保存token
    localStorage.setItem('token', adminLoginResult.token);
    localStorage.setItem('userType', 'admin');
    
    // 测试获取管理员信息
    console.log('测试获取管理员信息...');
    const adminProfile = await authApi.getAdminProfile();
    console.log('管理员信息:', adminProfile);
    
    // 测试获取课程列表
    console.log('测试获取课程列表...');
    const courses = await api.get('/courses');
    console.log('课程列表:', courses);
    
    return {
      success: true,
      message: '前后端API连接测试成功',
      data: {
        adminLoginResult,
        adminProfile,
        courses
      }
    };
  } catch (error) {
    console.error('API连接测试失败:', error);
    return {
      success: false,
      message: '前后端API连接测试失败',
      error: error.response?.data || error.message
    };
  }
};