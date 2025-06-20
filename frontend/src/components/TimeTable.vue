<template>
  <div class="timetable-container">
    <a-table
      :columns="processedColumns"
      :data="tableData"
      :bordered="true"
      :pagination="false"
      :scroll="{ x: '100%' }"
      row-key="time"
    >
      <!-- 时间列自定义渲染 -->
      <template #time="{ record }">
        <div class="time-cell">
          <div class="time-name">{{ record.time.split('\n')[0] }}</div>
          <div class="time-range">{{ record.time.split('\n')[1] }}</div>
        </div>
      </template>
      
      <!-- 动态生成星期列 -->
      <template v-for="day in days" :key="day.key" #[day.dataIndex]="{ record }">
        <div class="day-cell">
          <div v-for="course in record[day.dataIndex]" :key="course.id" class="course-item">
            <div class="course-name">{{ course.name }}</div>
            <div class="course-info">{{ course.teacher }}</div>
            <div class="course-info">{{ course.location }}</div>
          </div>
          <div v-if="!record[day.dataIndex] || record[day.dataIndex].length === 0" class="empty-cell">
            -
          </div>
        </div>
      </template>
    </a-table>
  </div>
</template>

<script setup>
import { computed } from 'vue';

const props = defineProps({
  sections: {
    type: Array,
    default: () => []
  }
});

// 时间段定义
const timeSlots = [
  { id: 1, name: '第1-2节', startTime: '08:00', endTime: '09:40' },
  { id: 2, name: '第3-4节', startTime: '10:00', endTime: '11:40' },
  { id: 3, name: '第5-6节', startTime: '14:00', endTime: '15:40' },
  { id: 4, name: '第7-8节', startTime: '16:00', endTime: '17:40' },
  { id: 5, name: '第9-10节', startTime: '19:00', endTime: '20:40' }
];

// 星期定义 (0=周日,1=周一,...,6=周六)
const days = [
  { title: '周一', dataIndex: '1', key: 'mon' },
  { title: '周二', dataIndex: '2', key: 'tue' },
  { title: '周三', dataIndex: '3', key: 'wed' },
  { title: '周四', dataIndex: '4', key: 'thu' },
  { title: '周五', dataIndex: '5', key: 'fri' },
  { title: '周六', dataIndex: '6', key: 'sat' },
  { title: '周日', dataIndex: '0', key: 'sun' }
];

// 处理列定义
const processedColumns = computed(() => {
  return [
    {
      title: '时间/星期',
      dataIndex: 'time',
      width: 100,
      fixed: 'left',
      slotName: 'time'
    },
    ...days.map(day => ({
      title: day.title,
      dataIndex: day.dataIndex,
      slotName: day.dataIndex
    }))
  ];
});

// 生成表格数据
const tableData = computed(() => {
  return timeSlots.map(slot => {
    const row = {
      time: `${slot.name}\n${slot.startTime}-${slot.endTime}`
    };
    
    // 初始化所有天数为空数组
    days.forEach(day => {
      row[day.dataIndex] = [];
    });
    
    // 填充课程数据
    props.sections.forEach(section => {
      const timeSlot = section.timeSlot;
      
      // 检查是否属于当前时间段
      if (timeSlot.startTime === slot.startTime) {
        const dayKey = timeSlot.dayOfWeek.toString();
        
        // 确保dayKey在有效范围内
        if (Object.prototype.hasOwnProperty.call(row, dayKey)) {
          row[dayKey].push({
            id: section.id,
            name: section.course.name,
            teacher: section.teacher.name,
            location: `${section.classroom.building}${section.classroom.room}`
          });
        }
      }
    });
    
    return row;
  });
});
</script>

<style scoped>
.timetable-container {
  margin-top: 20px;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial;
}

.time-cell {
  font-weight: 500;
  line-height: 1.4;
  padding: 8px 4px;
  background-color: var(--color-fill-2);
}

.time-name {
  font-weight: 600;
  font-size: 13px;
}

.time-range {
  font-size: 12px;
  color: var(--color-text-2);
}

.day-cell {
  min-height: 90px;
  padding: 4px;
}

.course-item {
  background-color: var(--color-primary-light-1);
  border-radius: 4px;
  padding: 6px 8px;
  margin-bottom: 6px;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
  border-left: 3px solid var(--color-primary);
}

.course-name {
  font-weight: 600;
  font-size: 13px;
  margin-bottom: 2px;
  color: var(--color-text-1);
}

.course-info {
  font-size: 12px;
  color: var(--color-text-2);
  line-height: 1.4;
}

.empty-cell {
  color: var(--color-text-3);
  text-align: center;
  padding: 20px 0;
  font-size: 12px;
}

/* 表头样式 */
:deep(.arco-table-th) {
  background-color: var(--color-fill-2) !important;
  font-weight: 600;
  text-align: center;
}

/* 时间列样式 */
:deep(.arco-table-col-fixed-left) {
  background-color: var(--color-fill-2);
  z-index: 1;
}
</style>