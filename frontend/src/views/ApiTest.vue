<template>
  <div class="api-test-container">
    <a-card title="API连接测试">
      <a-space direction="vertical" :size="16" fill>
        <a-button type="primary" @click="runTest" :loading="loading">
          测试API连接
        </a-button>
        
        <a-alert v-if="result" :type="result.success ? 'success' : 'error'" :title="result.message" />
        
        <a-collapse v-if="result && result.success">
          <a-collapse-item header="测试结果详情" key="1">
            <pre>{{ JSON.stringify(result.data, null, 2) }}</pre>
          </a-collapse-item>
        </a-collapse>
        
        <a-collapse v-if="result && !result.success">
          <a-collapse-item header="错误详情" key="1">
            <pre>{{ JSON.stringify(result.error, null, 2) }}</pre>
          </a-collapse-item>
        </a-collapse>
      </a-space>
    </a-card>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { testApiConnection } from '../utils/apiTest';

const loading = ref(false);
const result = ref(null);

const runTest = async () => {
  loading.value = true;
  try {
    result.value = await testApiConnection();
  } catch (error) {
    result.value = {
      success: false,
      message: '测试执行出错',
      error: error.message
    };
  } finally {
    loading.value = false;
  }
};
</script>

<style scoped>
.api-test-container {
  max-width: 800px;
  margin: 40px auto;
}

pre {
  background-color: #f5f5f5;
  padding: 10px;
  border-radius: 4px;
  overflow-x: auto;
}
</style>