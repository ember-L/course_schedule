import { createRouter, createWebHistory } from 'vue-router';
import { useAuthStore } from '../store/authStore';

// 页面组件
const routes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue'),
    meta: { requiresAuth: false }
  },
  {
    path: '/',
    name: 'Home',
    component: () => import('../views/Home.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/courses',
    name: 'Courses',
    component: () => import('../views/Courses.vue'),
    meta: { requiresAuth: true, roles: ['admin', 'teacher'] }
  },
  {
    path: '/teachers',
    name: 'Teachers',
    component: () => import('../views/Teachers.vue'),
    meta: { requiresAuth: true, roles: ['admin'] }
  },
  {
    path: '/classrooms',
    name: 'Classrooms',
    component: () => import('../views/Classrooms.vue'),
    meta: { requiresAuth: true, roles: ['admin'] }
  },
  {
    path: '/sections',
    name: 'Sections',
    component: () => import('../views/Sections.vue'),
    meta: { requiresAuth: true, roles: ['admin', 'teacher'] }
  },
  {
    path: '/students',
    name: 'Students',
    component: () => import('../views/Students.vue'),
    meta: { requiresAuth: true, roles: ['admin'] }
  },
  {
    path: '/enrollments',
    name: 'Enrollments',
    component: () => import('../views/Enrollments.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/schedule',
    name: 'Schedule',
    component: () => import('../views/Schedule.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/change-password',
    name: 'ChangePassword',
    component: () => import('../views/ChangePassword.vue'),
    meta: { requiresAuth: true }
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

// 路由守卫
router.beforeEach(async (to, from, next) => {
  const authStore = useAuthStore();
  const requiresAuth = to.matched.some(record => record.meta.requiresAuth !== false);
  
  // 如果有token但没有用户信息，尝试拉取用户信息但不自动登出
  if (authStore.token && !authStore.user) {
    try {
      await authStore.fetchCurrentUser(false); // 不自动登出
    } catch (error) {
      console.error('拉取用户信息失败', error);
      // 只有在明确需要认证的页面才登出
      if (requiresAuth) {
        await authStore.logout();
      }
    }
  }
  
  const isAuthenticated = authStore.isAuthenticated;
  
  // 如果需要认证但未登录，重定向到登录页
  if (requiresAuth && !isAuthenticated) {
    next({ name: 'Login' });
    return;
  }
  
  // 如果已登录但访问登录页，重定向到首页
  if (to.name === 'Login' && isAuthenticated) {
    next({ name: 'Home' });
    return;
  }
  
  // 检查角色权限
  if (to.meta.roles && isAuthenticated) {
    const userType = authStore.userType;
    if (!to.meta.roles.includes(userType)) {
      // 如果用户角色不符合要求，重定向到首页
      next({ name: 'Home' });
      return;
    }
  }
  
  next();
});

export default router;