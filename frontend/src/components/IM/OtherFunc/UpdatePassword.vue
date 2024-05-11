<template>
  <el-card class="change-password-card" shadow="hover">
    <div class="form-container">
      <el-form ref="passwordForm" :model="passwordData" @submit.prevent="changePassword">
        <el-form-item label="Current Password" prop="currentPassword" required>
          <el-input
              type="password"
              v-model="passwordData.currentPassword"
              autocomplete="current-password"
              placeholder="Enter your current password"></el-input>
        </el-form-item>
        <el-form-item label="New Password" prop="newPassword" required>
          <el-input
              type="password"
              v-model="passwordData.newPassword"
              autocomplete="new-password"
              placeholder="Enter your new password"></el-input>
        </el-form-item>
        <el-form-item label="Confirm New Password" prop="confirmNewPassword" required>
          <el-input
              type="password"
              v-model="passwordData.confirmNewPassword"
              autocomplete="new-password"
              placeholder="Confirm your new password"
              @input="validateNewPassword"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" native-type="submit">Change Password</el-button>
        </el-form-item>
      </el-form>
    </div>
  </el-card>
</template>

<script setup>
import { ref } from 'vue';
import { ElCard, ElForm, ElFormItem, ElInput, ElButton } from 'element-plus';
import axios from 'axios';

const passwordData = ref({
  currentPassword: '',
  newPassword: '',
  confirmNewPassword: ''
});

const passwordForm = ref(null);

const validateNewPassword = () => {
  if (passwordData.value.newPassword !== passwordData.value.confirmNewPassword) {
    alert('The new passwords do not match.');
    return false;
  }
  return true;
};

const changePassword = async () => {
  if (passwordForm.value.validate()) {
    if (!validateNewPassword()) {
      return;
    }
    try {
      const response = await axios.post('/api/change-password', {
        currentPassword: passwordData.value.currentPassword,
        newPassword: passwordData.value.newPassword
      });
      if (response.data.success) {
        alert('Password successfully changed.');
        passwordData.value.currentPassword = '';
        passwordData.value.newPassword = '';
        passwordData.value.confirmNewPassword = '';
      } else {
        alert(response.data.error);
      }
    } catch (error) {
      console.error('Error changing password:', error);
      alert('Failed to change password.');
    }
  }
};
</script>

<style scoped>
.change-password-card {
  max-width: 400px;
  margin: 50px auto;
  padding: 20px;
}

.form-container {
  padding: 20px;
}
</style>
