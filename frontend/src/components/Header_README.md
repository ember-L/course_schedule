# Header 组件优化说明

## 功能特性

### 1. 增强的用户体验
- **Logo点击导航**: 点击Logo可快速返回首页
- **动画效果**: 添加了悬停动画和过渡效果
- **视觉反馈**: 菜单项选中时有滑动动画效果
- **确认对话框**: 退出登录时显示确认对话框

### 2. 通知系统
- **通知图标**: 显示未读通知数量
- **徽章显示**: 使用红点标识新通知
- **点击交互**: 点击通知图标可查看通知详情

### 3. 用户菜单增强
- **个人信息**: 快速访问个人信息页面
- **系统设置**: 访问系统设置功能
- **视觉分隔**: 使用分割线区分不同功能区域
- **退出登录**: 突出显示退出登录选项

### 4. 移动端适配
- **响应式设计**: 完美适配各种屏幕尺寸
- **侧边菜单**: 移动端使用抽屉式侧边菜单
- **触摸优化**: 针对触摸设备优化交互体验
- **自适应布局**: 根据屏幕大小自动调整布局

## 技术实现

### 1. 响应式断点
```css
/* 大屏幕 */
@media (max-width: 1200px) {
  /* 调整间距和菜单项 */
}

/* 平板 */
@media (max-width: 768px) {
  /* 隐藏桌面菜单，显示移动菜单按钮 */
}

/* 手机 */
@media (max-width: 480px) {
  /* 进一步优化小屏幕显示 */
}
```

### 2. 动画效果
```css
/* 悬停动画 */
.menu-item:hover {
  transform: translateY(-1px);
  transition: all 0.3s ease;
}

/* 选中动画 */
.menu-item.arco-menu-item-selected::after {
  animation: slideIn 0.3s ease forwards;
}

@keyframes slideIn {
  to { width: 80%; }
}
```

### 3. 状态管理
```javascript
// 响应式数据
const mobileMenuVisible = ref(false);
const notificationCount = ref(3);
const hasNotifications = computed(() => notificationCount.value > 0);

// 权限控制
const isAdmin = computed(() => authStore.isAdmin);
const isTeacher = computed(() => authStore.isTeacher);
const isStudent = computed(() => authStore.isStudent);
```

## 使用方式

### 1. 基本使用
```vue
<template>
  <Header />
</template>

<script setup>
import Header from '@/components/Header.vue';
</script>
```

### 2. 权限控制
组件会根据用户角色自动显示/隐藏相应的菜单项：
- **管理员**: 显示所有菜单项
- **教师**: 显示课程管理、排课管理、选课管理、课表查询
- **学生**: 显示选课管理、课表查询

### 3. 事件处理
```javascript
// 菜单点击
const handleMenuClick = (key) => {
  router.push(key);
};

// 移动端菜单点击
const handleMobileMenuClick = (key) => {
  router.push(key);
  mobileMenuVisible.value = false;
};

// 退出登录
const handleLogout = async () => {
  Modal.confirm({
    title: '确认退出',
    content: '您确定要退出登录吗？',
    onOk: async () => {
      await authStore.logout();
      router.push('/login');
    }
  });
};
```

## 样式定制

### 1. CSS变量
组件使用了ArcoDesign的CSS变量系统：
```css
:root {
  --primary-color: #165dff;
  --primary-hover: #4080ff;
  --spacing-xs: 4px;
  --spacing-sm: 8px;
  --spacing-md: 12px;
  --spacing-lg: 16px;
  --spacing-xl: 24px;
  --border-radius-medium: 6px;
  --border-radius-large: 8px;
  --shadow-2: 0 4px 10px rgba(0, 0, 0, 0.1);
}
```

### 2. 主题定制
可以通过修改CSS变量来自定义主题：
```css
.main-header {
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--primary-hover) 100%);
}
```

## 性能优化

### 1. 事件监听器管理
```javascript
onMounted(() => {
  const handleResize = () => {
    if (window.innerWidth > 768 && mobileMenuVisible.value) {
      mobileMenuVisible.value = false;
    }
  };
  
  window.addEventListener('resize', handleResize);
  
  onUnmounted(() => {
    window.removeEventListener('resize', handleResize);
  });
});
```

### 2. 计算属性优化
使用计算属性避免不必要的重复计算：
```javascript
const userTypeText = computed(() => {
  if (isAdmin.value) return '管理员';
  if (isTeacher.value) return '教师';
  if (isStudent.value) return '学生';
  return '用户';
});
```

## 扩展功能

### 1. 通知系统
可以扩展通知功能：
```javascript
const showNotifications = () => {
  // 实现通知面板
  // 可以集成WebSocket实时通知
  // 支持通知分类和标记已读
};
```

### 2. 用户设置
可以扩展用户设置功能：
```javascript
const showSettings = () => {
  // 实现系统设置面板
  // 支持主题切换、语言设置等
};
```

### 3. 搜索功能
可以添加全局搜索功能：
```vue
<div class="search-section">
  <a-input-search 
    placeholder="搜索课程、教师、学生..." 
    @search="handleSearch"
  />
</div>
```

## 最佳实践

1. **无障碍访问**: 确保所有交互元素都有适当的ARIA标签
2. **键盘导航**: 支持键盘导航和快捷键
3. **加载状态**: 在异步操作时显示加载状态
4. **错误处理**: 优雅处理各种错误情况
5. **国际化**: 支持多语言切换
6. **主题切换**: 支持明暗主题切换 