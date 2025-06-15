import api from '@/api/index';

// 课程相关API
export const courseApi = {
  getAll: () => api.get('/courses'),
  getById: (id) => api.get(`/courses/${id}`),
  create: (data) => api.post('/courses', data),
  update: (id, data) => api.put(`/courses/${id}`, data),
  delete: (id) => api.delete(`/courses/${id}`)
};

// 教师相关API
export const teacherApi = {
  getAll: () => api.get('/teachers'),
  getById: (id) => api.get(`/teachers/${id}`),
  create: (data) => api.post('/teachers', data),
  update: (id, data) => api.put(`/teachers/${id}`, data),
  delete: (id) => api.delete(`/teachers/${id}`)
};

// 教室相关API
export const classroomApi = {
  getAll: () => api.get('/classrooms'),
  getById: (id) => api.get(`/classrooms/${id}`),
  create: (data) => api.post('/classrooms', data),
  update: (id, data) => api.put(`/classrooms/${id}`, data),
  delete: (id) => api.delete(`/classrooms/${id}`)
};

// 时间段相关API
export const timeSlotApi = {
  getAll: () => api.get('/time-slots'),
  getById: (id) => api.get(`/time-slots/${id}`),
  create: (data) => api.post('/time-slots', data),
  update: (id, data) => api.put(`/time-slots/${id}`, data),
  delete: (id) => api.delete(`/time-slots/${id}`)
};

// 排课相关API
export const sectionApi = {
  getAll: () => api.get('/sections'),
  getById: (id) => api.get(`/sections/${id}`),
  create: (data) => api.post('/sections', data),
  update: (id, data) => api.put(`/sections/${id}`, data),
  delete: (id) => api.delete(`/sections/${id}`),
  checkConflict: (data) => api.post('/sections/check-conflict', data)
};

// 学生相关API
export const studentApi = {
  getAll: () => api.get('/students'),
  getById: (id) => api.get(`/students/${id}`),
  create: (data) => api.post('/students', data),
  update: (id, data) => api.put(`/students/${id}`, data),
  delete: (id) => api.delete(`/students/${id}`)
};

// 选课相关API
export const enrollmentApi = {
  getAll: () => api.get('/enrollments'),
  getByStudent: (studentId) => api.get(`/enrollments/student/${studentId}`),
  create: (data) => api.post('/enrollments', data),
  delete: (id) => api.delete(`/enrollments/${id}`),
  updateGrade: (id, grade) => api.put(`/enrollments/${id}/grade`, { grade })
};