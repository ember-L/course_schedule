<template>
  <div class="timetable-container">
    <a-table
      :columns="columns"
      :data="tableData"
      :bordered="true"
      :pagination="false"
      :scroll="{ x: '100%' }"
    >
      <template #cell="{ record, column }">
        <div v-if="column.dataIndex !== 'time'" class="cell-content">
          <template v-if="record[column.dataIndex]">
            <div 
              v-for="course in record[column.dataIndex]" 
              :key="course.id"
              class="course-item"
            >
              <div class="course-name">{{ course.name }}</div>
              <div class="course-info">{{ course.teacher }}</div>
              <div class="course-info">{{ course.location }}</div>
            </div>
          </template>
        </div>
      </template>
    </a-table>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue';

const props = defineProps({
  sections: {
    type: Array,
    default: () => []
  }
});

// 时间段
const timeSlots = [
  { id: 1, name: '第1-2节', startTime: '08:00', endTime: '09:40' },
  { id: 2, name: '第3-4节', startTime: '10:00', endTime: '11:40' },
  { id: 3, name: '第5-6节', startTime: '14:00', endTime: '15:40' },
  { id: 4, name: '第7-8节', startTime: '16:00', endTime: '17:40' },
  { id: 5, name: '第9-10节', startTime: '19:00', endTime: '20:40' }
];

// 表格列定义
const columns = [
  {
    title: '时间/星期',
    dataIndex: 'time',
    width: 100,
    fixed: 'left'
  },
  {
    title: '周一',
    dataIndex: '1'
  },
  {
    title: '周二',
    dataIndex: '2'
  },
  {
    title: '周三',
    dataIndex: '3'
  },
  {
    title: '周四',
    dataIndex: '4'
  },
  {
    title: '周五',
    dataIndex: '5'
  },
  {
    title: '周六',
    dataIndex: '6'
  },
  {
    title: '周日',
    dataIndex: '0'
  }
];

// 生成表格数据
const tableData = computed(() => {
  const data = timeSlots.map(slot => {
    const row = {
      time: `${slot.name}\n${slot.startTime}-${slot.endTime}`
    };
    
    // 初始化每天的数据为空数组
    for (let i = 0; i < 7; i++) {
      row[i.toString()] = [];
    }
    
    // 填充课程数据
    props.sections.forEach(section => {
      const timeSlot = section.timeSlot;
      // 简化处理，根据开始时间匹配时间段
      if (timeSlot.startTime === slot.startTime) {
        const dayKey = timeSlot.dayOfWeek.toString();
        row[dayKey].push({
          id: section.id,
          name: section.course.name,
          teacher: section.teacher.name,
          location: `${section.classroom.building}-${section.classroom.room}`
        });
      }
    });
    
    return row;
  });
  
  return data;
});

// 监听sections变化，重新计算表格数据
watch(() => props.sections, () => {
  // tableData会自动重新计算
}, { deep: true });
</script>

<style scoped>
.timetable-container {
  margin-top: 20px;
}

.cell-content {
  min-height: 80px;
}

.course-item {
  background-color: #e6f7ff;
  border-radius: 4px;
  padding: 4px 8px;
  margin-bottom: 4px;
}

.course-name {
  font-weight: bold;
  font-size: 14px;
}

.course-info {
  font-size: 12px;
  color: #666;
}
</style>