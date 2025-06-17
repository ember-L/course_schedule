<template>
  <div class="login-container">
    <div class="login-background">
      <div class="background-shapes">
        <div class="shape shape-1"></div>
        <div class="shape shape-2"></div>
        <div class="shape shape-3"></div>
      </div>
    </div>
    
    <div class="login-content">
      <div class="login-card fade-in">
        <div class="login-header">
          <div class="logo-section">
            <div class="logo-icon">
              <icon-book />
            </div>
            <h1 class="login-title">高校排课系统</h1>
            <p class="login-subtitle">欢迎使用智能排课管理系统</p>
          </div>
        </div>
        
        <a-form 
          :model="form" 
          ref="formRef" 
          :rules="rules" 
          layout="vertical"
          class="login-form"
          @submit.prevent="handleLogin"
        >
          <a-form-item field="userType" label="用户类型">
            <a-radio-group v-model="form.userType" class="user-type-group">
              <a-radio value="student" class="user-type-radio">
                <div class="radio-content">
                  <icon-user />
                  <span>学生</span>
                </div>
              </a-radio>
              <a-radio value="teacher" class="user-type-radio">
                <div class="radio-content">
                  <icon-user />
                  <span>教师</span>
                </div>
              </a-radio>
              <a-radio value="admin" class="user-type-radio">
                <div class="radio-content">
                  <icon-settings />
                  <span>管理员</span>
                </div>
              </a-radio>
            </a-radio-group>
          </a-form-item>
          
          <a-form-item field="username" label="用户名">
            <a-input 
              v-model="form.username" 
              placeholder="请输入用户名"
              size="large"
              class="login-input"
            >
              <template #prefix>
                <icon-user />
              </template>
            </a-input>
          </a-form-item>
          
          <a-form-item field="password" label="密码">
            <a-input-password 
              v-model="form.password" 
              placeholder="请输入密码"
              size="large"
              class="login-input"
            >
              <template #prefix>
                <icon-lock />
              </template>
            </a-input-password>
          </a-form-item>
          
          <a-form-item>
            <a-button 
              type="primary" 
              long 
              size="large"
              @click="handleLogin"

              :loading="loading"
              class="login-button"
            >
              <template #icon>
                <icon-login />
              </template>
              {{ loading ? '登录中...' : '登录' }}
            </a-button>
          </a-form-item>
        </a-form>
        
        <div class="login-footer">
          <div class="login-tips">
            <div class="tip-item">
              <icon-info-circle />
              <span>学生：使用学号登录</span>
            </div>
            <div class="tip-item">
              <icon-info-circle />
              <span>教师：使用工号登录</span>
            </div>
            <div class="tip-item">
              <icon-info-circle />
              <span>管理员：使用管理员账号登录</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue';
import { Message } from '@arco-design/web-vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/store/authStore';

const router = useRouter();
const authStore = useAuthStore();
const formRef = ref(null);
const loading = ref(false);

const form = reactive({
  userType: 'student',
  username: '',
  password: ''
});

const rules = {
  username: [
    { required: true, message: '请输入用户名' },
    { minLength: 2, message: '用户名至少2个字符' }
  ],
  password: [
    { required: true, message: '请输入密码' },
    { minLength: 6, message: '密码至少6个字符' }
  ]
};

const handleLogin = async () => {
  try {
    // 验证表单
    await formRef.value.validate();
    
    loading.value = true;
    await authStore.login(form.userType, form.username, form.password);
    Message.success('登录成功');
    
    // 根据用户类型跳转
    if (form.userType === 'admin') {
      router.push('/');
    } else if (form.userType === 'teacher') {
      router.push('/');
    } else {
      router.push('/schedule');
    }
  } catch (error) {
    console.error('登录错误:', error);
    Message.error(`登录失败：用户名或密码错误`);
    // 清空密码字段
    form.password = '';
  } finally {
    loading.value = false;
  }
};

</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  width: 100vw;
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--primary-hover) 100%);
  position: relative;
  overflow: hidden;
  margin: 0;
  padding: 0;
}

.login-background {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  z-index: 1;
}

.background-shapes {
  position: relative;
  width: 100%;
  height: 100%;
}

.shape {
  position: absolute;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.1);
  animation: float 6s ease-in-out infinite;
  backdrop-filter: blur(5px);
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.shape-1 {
  width: 100px;
  height: 100px;
  top: 20%;
  left: 10%;
  animation-delay: 0s;
}

.shape-2 {
  width: 150px;
  height: 150px;
  top: 60%;
  right: 15%;
  animation-delay: 2s;
}

.shape-3 {
  width: 80px;
  height: 80px;
  bottom: 20%;
  left: 20%;
  animation-delay: 4s;
}

@keyframes float {
  0%, 100% {
    transform: translateY(0px) rotate(0deg);
    opacity: 0.6;
  }
  50% {
    transform: translateY(-20px) rotate(180deg);
    opacity: 0.8;
  }
}

.login-content {
  position: relative;
  z-index: 2;
  width: 100%;
  max-width: 450px;
  padding: var(--spacing-xl);
  margin: 0 auto;
}

.login-card {
  background: var(--bg-primary);
  border-radius: var(--border-radius-xl);
  box-shadow: var(--shadow-3);
  padding: var(--spacing-xxl);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  animation: fadeIn 0.6s ease-out;
  position: relative;
  overflow: hidden;
}

.login-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(22, 93, 255, 0.05) 0%, rgba(64, 128, 255, 0.05) 100%);
  pointer-events: none;
}

.login-header {
  text-align: center;
  margin-bottom: var(--spacing-xxl);
  position: relative;
  z-index: 1;
}

.logo-section {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: var(--spacing-md);
}

.logo-icon {
  width: 80px;
  height: 80px;
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--primary-hover) 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 36px;
  color: white;
  box-shadow: var(--shadow-2);
  transition: all 0.3s ease;
}

.logo-icon:hover {
  transform: scale(1.05);
  box-shadow: var(--shadow-3);
}

.login-title {
  font-size: 28px;
  font-weight: 600;
  color: var(--text-primary);
  margin: 0;
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--primary-hover) 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.login-subtitle {
  font-size: 16px;
  color: var(--text-secondary);
  margin: 0;
  font-weight: 400;
}

.login-form {
  margin-bottom: var(--spacing-xl);
  position: relative;
  z-index: 1;
}

.login-form .arco-form-item {
  margin-bottom: var(--spacing-lg);
}

.login-form .arco-form-item-label {
  font-weight: 500;
  color: var(--text-primary);
  margin-bottom: var(--spacing-xs);
}

.login-form .arco-form-item-label::before {
  content: '';
  display: inline-block;
  width: 4px;
  height: 4px;
  background: var(--primary-color);
  border-radius: 50%;
  margin-right: var(--spacing-xs);
  vertical-align: middle;
}

.user-type-group {
  display: flex;
  gap: var(--spacing-sm);
  width: 100%;
}

.user-type-radio {
  flex: 1;
}

.user-type-radio .arco-radio {
  width: 100%;
  height: 60px;
  border: 2px solid var(--border-primary);
  border-radius: var(--border-radius-medium);
  transition: all 0.2s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}

.user-type-radio .arco-radio:hover {
  border-color: var(--primary-color);
  background-color: var(--primary-light);
  transform: translateY(-1px);
  box-shadow: var(--shadow-1);
}

.user-type-radio .arco-radio-checked {
  border-color: var(--primary-color);
  background-color: var(--primary-light);
  transform: translateY(-1px);
  box-shadow: var(--shadow-1);
}

.radio-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: var(--spacing-xs);
  font-size: 14px;
  font-weight: 500;
  color: var(--text-primary);
}

.radio-content .arco-icon {
  font-size: 20px;
  color: var(--primary-color);
  transition: all 0.2s ease;
}

.user-type-radio .arco-radio:hover .radio-content .arco-icon {
  transform: scale(1.1);
}

.login-input {
  border-radius: var(--border-radius-medium);
  transition: all 0.2s ease;
}

.login-input .arco-input-prefix {
  color: var(--text-tertiary);
}

.login-input:focus-within {
  border-color: var(--primary-color);
  box-shadow: 0 0 0 2px var(--primary-light);
}

.login-button {
  height: 48px;
  font-size: 16px;
  font-weight: 500;
  border-radius: var(--border-radius-medium);
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--primary-hover) 100%);
  border: none;
  box-shadow: var(--shadow-1);
  transition: all 0.2s ease;
}

.login-button:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow-2);
}

.login-button:active {
  transform: translateY(0);
}

.login-footer {
  text-align: center;
  margin-top: var(--spacing-lg);
}

.login-tips {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
  padding: var(--spacing-md);
  background: var(--bg-secondary);
  border-radius: var(--border-radius-medium);
  border: 1px solid var(--border-primary);
}

.tip-item {
  display: flex;
  align-items: center;
  gap: var(--spacing-xs);
  font-size: 12px;
  color: var(--text-tertiary);
  transition: all 0.2s ease;
}

.tip-item:hover {
  color: var(--text-secondary);
}

.tip-item .arco-icon {
  color: var(--primary-color);
  font-size: 14px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .login-container {
    padding: 16px;
  }
  
  .login-content {
    padding: 16px;
    max-width: 100%;
  }
  
  .login-card {
    padding: 24px;
  }
  
  .user-type-group {
    flex-direction: column;
  }
  
  .user-type-radio .arco-radio {
    height: 50px;
  }
  
  .logo-icon {
    width: 60px;
    height: 60px;
    font-size: 28px;
  }
  
  .login-title {
    font-size: 24px;
  }
  
  .login-subtitle {
    font-size: 14px;
  }
}

@media (max-width: 480px) {
  .login-container {
    padding: 8px;
  }
  
  .login-content {
    padding: 8px;
  }
  
  .login-card {
    padding: 20px;
  }
  
  .login-title {
    font-size: 20px;
  }
  
  .login-subtitle {
    font-size: 12px;
  }
  
  .logo-icon {
    width: 50px;
    height: 50px;
    font-size: 24px;
  }
}

.fade-in {
  animation: fadeIn 0.6s ease-out;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>