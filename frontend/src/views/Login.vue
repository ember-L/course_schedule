<template>
  <div class="login-container">
    <a-card class="login-card">
      <template #title>
        <div class="login-title">高校排课系统</div>
      </template>
      <a-form :model="form" ref="formRef" :rules="rules" layout="vertical">
        <a-form-item field="userType" label="用户类型">
          <a-radio-group v-model="form.userType">
            <a-radio value="student">学生</a-radio>
            <a-radio value="teacher">教师</a-radio>
            <a-radio value="admin">管理员</a-radio>
          </a-radio-group>
        </a-form-item>
        <a-form-item field="username" label="用户名">
          <a-input v-model="form.username" placeholder="请输入用户名" />
        </a-form-item>
        <a-form-item field="password" label="密码">
          <a-input-password v-model="form.password" placeholder="请输入密码" />
        </a-form-item>
        <a-form-item>
          <a-button type="primary" long @click="handleLogin" :loading="loading">登录</a-button>
        </a-form-item>
      </a-form>
    </a-card>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue';
import { Message } from '@arco-design/web-vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '../store/authStore';

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
  username: [{ required: true, message: '请输入用户名' }],
  password: [{ required: true, message: '请输入密码' }]
};

const handleLogin = async () => {
  formRef.value.validate(async (valid) => {
    if (valid) {
      loading.value = true;
      try {
        await authStore.login(form.userType, form.username, form.password);
        Message.success('登录成功');
        
        // 根据用户类型跳转到不同页面
        if (form.userType === 'admin') {
          router.push('/');
        } else if (form.userType === 'teacher') {
          router.push('/sections');
        } else {
          router.push('/schedule');
        }
      } catch (error) {
        Message.error('登录失败：' + (error.message || '用户名或密码错误'));
      } finally {
        loading.value = false;
      }
    }
  });
};
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f5f5f5;
}

.login-card {
  width: 400px;
}

.login-title {
  font-size: 24px;
  text-align: center;
  margin-bottom: 20px;
}

.login-tips {
  margin-top: 20px;
  font-size: 12px;
  color: #999;
  text-align: center;
}
</style>