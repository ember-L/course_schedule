# Schedule.vue 错误修复说明

## 问题描述

在 Schedule.vue 第285行出现错误：
```
TypeError: Cannot read properties of null (reading 'id')
```

## 错误原因

在访问 `currentUser.value.id` 时，`currentUser.value` 为 `null`，导致无法读取 `id` 属性。

## 根本原因分析

1. **异步加载问题**: `loadCurrentUser` 函数可能失败，但 `loadUserSchedule` 仍然被调用
2. **缺少空值检查**: 在访问对象属性前没有进行适当的空值检查
3. **错误处理不完善**: 没有正确处理API调用失败的情况

## 修复方案

### 1. 改进 `loadCurrentUser` 函数

```javascript
const loadCurrentUser = async () => {
  loading.value = true;
  try {
    const userId = authStore.currentUser?.id;
    if (!userId) {
      Message.error('用户信息获取失败');
      return false; // 返回false表示加载失败
    }

    let userData = null;
    if (isStudent.value) {
      userData = await studentApi.getById(userId);
    } else if (isTeacher.value) {
      userData = await teacherApi.getById(userId);
    } else {
      Message.error('未知的用户类型');
      return false;
    }
    
    if (!userData) {
      Message.error('用户信息不存在');
      return false;
    }
    
    currentUser.value = userData;
    return true; // 返回true表示加载成功
  } catch (error) {
    Message.error('获取用户信息失败');
    console.error(error);
    return false; // 返回false表示加载失败
  } finally {
    loading.value = false;
  }
};
```

### 2. 改进 `initializeData` 函数

```javascript
const initializeData = async () => {
  if (isAdmin.value) {
    await loadAllUsers();
  } else {
    const userLoaded = await loadCurrentUser();
    if (userLoaded) {
      await loadUserSchedule();
    }
  }
};
```

### 3. 在 `loadUserSchedule` 中添加空值检查

```javascript
} else if (isStudent.value) {
  // 学生查看自己的课表
  if (!currentUser.value?.id) {
    Message.error('用户信息不完整，请重新登录');
    return;
  }
  const enrollments = await enrollmentApi.getByStudent(currentUser.value.id);
  // ...
} else if (isTeacher.value) {
  // 教师查看自己的课表
  if (!currentUser.value?.id) {
    Message.error('用户信息不完整，请重新登录');
    return;
  }
  const allSections = await sectionApi.getAll();
  sectionsData = allSections.filter(section => section.teacherId === currentUser.value.id);
}
```

### 4. 改进导出功能的空值处理

```javascript
const fileName = isAdmin.value 
  ? `课表_${selectedUser.value}_${selectedSemester.value}.pdf`
  : `课表_${currentUser.value?.name || '未知用户'}_${selectedSemester.value}.pdf`;
```

### 5. 优化模板条件渲染

```vue
<!-- 空状态 -->
<div v-else-if="!currentUser && !selectedUser && !loading" class="empty-container">
  <a-empty description="请选择用户查看课表" />
</div>

<!-- 课表内容 -->
<div v-else-if="currentUser || selectedUser">
  <!-- 内容 -->
</div>
```

## 修复的关键点

### 1. 返回值控制
- `loadCurrentUser` 函数现在返回布尔值，表示加载是否成功
- 只有在用户信息加载成功后才调用 `loadUserSchedule`

### 2. 空值检查
- 使用可选链操作符 `?.` 进行安全的属性访问
- 在关键位置添加 `if (!currentUser.value?.id)` 检查

### 3. 错误处理
- 为每个API调用添加适当的错误处理
- 提供用户友好的错误消息
- 在错误情况下设置默认值

### 4. 数据验证
- 验证API返回的数据是否有效
- 检查用户类型是否合法
- 确保必要的数据字段存在

## 预防措施

### 1. 代码审查
- 在访问对象属性前始终检查对象是否存在
- 使用可选链操作符 `?.` 和空值合并操作符 `??`

### 2. 错误边界
- 为异步操作添加适当的错误处理
- 提供降级方案和用户友好的错误消息

### 3. 测试覆盖
- 测试各种错误情况
- 验证错误处理逻辑的正确性

## 测试建议

### 1. 正常流程测试
- [ ] 学生用户登录后正常显示课表
- [ ] 教师用户登录后正常显示课表
- [ ] 管理员用户可以选择查看任意用户课表

### 2. 错误情况测试
- [ ] 用户信息获取失败时的处理
- [ ] API调用失败时的处理
- [ ] 网络异常时的处理
- [ ] 数据不完整时的处理

### 3. 边界情况测试
- [ ] 用户未登录时的处理
- [ ] 用户类型不匹配时的处理
- [ ] 空数据时的处理

## 总结

通过以上修复，Schedule.vue 组件现在具备了完善的错误处理机制：

1. **防止空值访问**: 在访问对象属性前进行适当的检查
2. **异步操作控制**: 确保数据加载的顺序和依赖关系
3. **用户友好**: 提供清晰的错误消息和降级方案
4. **代码健壮性**: 增强了代码的容错能力和稳定性

这些修复确保了组件在各种情况下都能正常工作，提供了更好的用户体验。 