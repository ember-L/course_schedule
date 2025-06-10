import { defineStore } from 'pinia';
import api from '../api';

export const useCourseStore = defineStore('course', {
  state: () => ({
    courses: [],
    loading: false,
    error: null
  }),
  
  actions: {
    async fetchCourses() {
      this.loading = true;
      try {
        const data = await api.get('/courses');
        this.courses = data;
        this.error = null;
      } catch (err) {
        this.error = err.message || '获取课程失败';
      } finally {
        this.loading = false;
      }
    },
    
    async addCourse(course) {
      this.loading = true;
      try {
        const data = await api.post('/courses', course);
        this.courses.push(data);
        return data;
      } catch (err) {
        this.error = err.message || '添加课程失败';
        throw err;
      } finally {
        this.loading = false;
      }
    },
    
    async updateCourse(course) {
      this.loading = true;
      try {
        const data = await api.put(`/courses/${course.id}`, course);
        const index = this.courses.findIndex(c => c.id === course.id);
        if (index !== -1) {
          this.courses[index] = data;
        }
        return data;
      } catch (err) {
        this.error = err.message || '更新课程失败';
        throw err;
      } finally {
        this.loading = false;
      }
    },
    
    async deleteCourse(id) {
      this.loading = true;
      try {
        await api.delete(`/courses/${id}`);
        this.courses = this.courses.filter(c => c.id !== id);
      } catch (err) {
        this.error = err.message || '删除课程失败';
        throw err;
      } finally {
        this.loading = false;
      }
    }
  }
});