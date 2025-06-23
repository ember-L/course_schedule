# 高校排课系统 RESTful API 文档

---

## 一、后端RESTful API文档

### 1. 认证与用户

#### 1.1 登录
- **POST** `/admin/login`（管理员）
- **POST** `/teacher/login`（教师）
- **POST** `/student/login`（学生）

**请求体：**
```json
{
  "username": "admin",
  "password": "123456"
}
```
**返回：**
```json
{
  "token": "jwt-token",
  "user": { "id": 1, "name": "管理员" },
  "expiresAt": "2024-06-08T12:00:00Z"
}
```

---

### 2. 课程管理

#### 2.1 获取课程列表
- **GET** `/courses`
- **权限**：所有用户

**返回：**
```json
[
  { "id": 1, "code": "CS101", "name": "计算机基础", "credits": 3, "description": "..." }
]
```

#### 2.2 新增课程
- **POST** `/courses`
- **权限**：管理员

**请求体：**
```json
{ "code": "CS102", "name": "数据结构", "credits": 4, "description": "..." }
```

---

### 3. 教师管理

#### 3.1 获取教师列表
- **GET** `/teachers`
- **权限**：管理员

#### 3.2 新增教师
- **POST** `/teachers`
- **权限**：管理员

---

### 4. 学生管理

#### 4.1 获取学生列表
- **GET** `/students`
- **权限**：管理员

#### 4.2 新增学生
- **POST** `/students`
- **权限**：管理员

---

### 5. 教室管理

#### 5.1 获取教室列表
- **GET** `/classrooms`
- **权限**：所有用户

---

### 6. 时间段管理

#### 6.1 获取时间段列表
- **GET** `/timeSlots`
- **权限**：所有用户

---

### 7. 排课管理（Section）

#### 7.1 获取排课列表
- **GET** `/sections`
- **权限**：管理员、教师

#### 7.2 新增排课
- **POST** `/sections`
- **权限**：管理员、教师

**请求体：**
```json
{
  "courseId": 1,
  "teacherId": 2,
  "classroomId": 3,
  "timeSlotId": 4,
  "semester": "2024春",
  "maxStudents": 50
}
```

#### 7.3 自动排课
- **POST** `/sections/auto-schedule`
- **权限**：管理员

**请求体：**
```json
{
  "courseId": 1,
  "teacherId": 2,
  "semester": "2024春",
  "maxStudents": 50
}
```
**返回：**
```json
{
  "id": 10,
  "courseId": 1,
  "teacherId": 2,
  "classroomId": 3,
  "timeSlotId": 4,
  "semester": "2024春",
  "maxStudents": 50
}
```

---

### 8. 选课管理（Enrollment）

#### 8.1 学生选课
- **POST** `/enrollments/enroll`
- **权限**：学生

**请求体：**
```json
{ "sectionId": 10 }
```

#### 8.2 退课
- **DELETE** `/enrollments/drop/:id`
- **权限**：学生

---

### 9. 课表查询

#### 9.1 获取我的课表
- **GET** `/schedule/my`
- **权限**：所有认证用户

#### 9.2 获取今日课表
- **GET** `/schedule/today`
- **权限**：所有认证用户

#### 9.3 管理员查任意学生/教师课表
- **GET** `/schedule/student/:studentId`
- **GET** `/schedule/teacher/:teacherId`
- **权限**：管理员


---

## 二、前端API调用示例

### 1. 登录
```js
await api.post('/admin/login', { username, password });
```

### 2. 获取课程列表
```js
await api.get('/courses');
```

### 3. 自动排课
```js
await api.post('/sections/auto-schedule', {
  courseId: 1,
  teacherId: 2,
  semester: '2024春',
  maxStudents: 50
});
```

### 4. 学生选课
```js
await api.post('/enrollments/enroll', { sectionId: 10 });
```

### 5. 获取我的课表
```js
await api.get('/schedule/my');
```

---

## 三、权限说明

- **管理员**：可管理所有数据，排课、自动排课、查所有课表、查日志等。
- **教师**：可查自己课程、排课、查自己课表。
- **学生**：可查自己课表、选课、退课。

---

## 四、返回通用格式

- 成功：
  ```json
  { "code": 0, "data": { ... }, "message": "操作成功" }
  ```
- 失败：
  ```json
  { "code": 1, "error": "错误信息" }
  ```

---

如需Swagger/OpenAPI文档、Markdown/HTML格式、接口字段详细说明或前端API模块注释，请联系开发者。 