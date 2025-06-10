<template>
  <div class="change-password-container">
    <a-card title="修改密码" :style="{ width: '500px' }">
      <a-form :model="form" ref="formRef" :rules="rules" layout="vertical">
        <a-form-item field="oldPassword" label="原密码">
          <a-input-password v-model="form.oldPassword" placeholder="请输入原密码" />
        </a-form-item>
        <a-form-item field="newPassword" label="新密码">
          <a-input-password v-model="form.newPassword" placeholder="请输入新密码" />
        </a-form-item>
        <a-form-item field="confirmPassword" label="确认新密码">
          <a-input-password v-model="form.confirmPassword" placeholder="请再次输入新密码" />
        </a-form-item>
        <a-form-item>
          <a-space>
            <a-button type="primary" @click="handleSubmit" :loading="loading">确认修改</a-button>
            <a-button @click="$router.back()">返回</a-button>
          </a-space>
        </a-form-item>
      </a-form>
    </a-card>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue';
import { Message } from '@arco-design/web-vue';
import { useRouter } from 'vue-router';
import { authApi } from '../api/auth';

const router = useRouter();
const formRef = ref(null);
const loading = ref(false);

const form = reactive({
  oldPassword: '',
  newPassword: '',
  confirmPassword: ''
});

const validateConfirmPassword = (value, cb) => {
  if (value !== form.newPassword) {
    return cb('两次输入的密码不一致');
  }
  return cb();
};

const rules = {
  oldPassword: [{ required: true, message: '请输入原密码' }],
  newPassword: [
    { required: true, message: '请输入新密码' },
    { minLength: 6, message: '密码长度不能小于6位' }
  ],
  confirmPassword: [
    { required: true, message: '请确认新密码' },
    { validator: validateConfirmPassword, message: '两次输入的密码不一致' }
  ]
};

const handleSubmit = async () => {
  formRef.value.validate(async (valid) => {
    if (valid) {
      loading.value = true;
      try {
        await authApi.changePassword(form.oldPassword, form.newPassword);
        Message.success('密码修改成功，请重新登录');
        // 退出登录并跳转到登录页
        router.push('/login');
      } catch (error) {
        Message.error('密码修改失败：' + (error.response?.data?.message || '原密码错误'));
      } finally {
        loading.value = false;
      }
    }
  });
};
</script>

<style scoped>
.change-password-container {
  display: flex;
  justify-content: center;
  padding: 40px;
}
</style>