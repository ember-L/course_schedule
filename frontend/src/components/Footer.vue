<template>
  <a-layout-footer class="main-footer">
    <div class="footer-content">
      <div class="footer-left">
        <span>高校排课系统 &copy; {{ new Date().getFullYear() }}</span>
      </div>
      <div class="footer-right">
        <span>智能排课 · 高效管理</span>
      </div>
    </div>
  </a-layout-footer>
</template>

<script setup>
</script>

<style scoped>
.main-footer {
  background: #fff;
  border-top: 1px solid #e5e6eb;
  padding: 0;
  height: 60px;
}
.footer-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 100%;
  padding: 0 32px;
  max-width: 1400px;
  margin: 0 auto;
  color: #4e5969;
  font-size: 14px;
}
.footer-left {
  display: flex;
  align-items: center;
}
.footer-right {
  display: flex;
  align-items: center;
}
@media (max-width: 768px) {
  .footer-content {
    flex-direction: column;
    gap: 8px;
    padding: 16px;
    text-align: center;
  }
}
</style> 