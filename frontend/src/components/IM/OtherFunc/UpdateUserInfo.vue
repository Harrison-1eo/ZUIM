<template>
  <el-card class="update-user-info-card" shadow="hover">
    <div class="form-container">
      <el-form :model="userInfo" ref="userInfoForm" label-position="top">
        <el-form-item label="Username">
          <el-input v-model="userInfo.username"></el-input>
        </el-form-item>
        <el-form-item label="Email">
          <el-input v-model="userInfo.email"></el-input>
        </el-form-item>
        <el-form-item label="Nickname">
          <el-input v-model="userInfo.nike_name"></el-input>
        </el-form-item>
        <el-form-item label="Country">
          <el-input v-model="userInfo.country"></el-input>
        </el-form-item>
        <el-form-item label="Province">
          <el-input v-model="userInfo.province"></el-input>
        </el-form-item>
        <el-form-item label="City">
          <el-input v-model="userInfo.city"></el-input>
        </el-form-item>
        <el-form-item label="Avatar">
          <el-upload
              class="avatar-uploader"
              action="path-to-api-to-upload-avatar"
              :show-file-list="false"
              :on-success="handleAvatarSuccess"
              :before-upload="beforeAvatarUpload">
            <img v-if="userInfo.avatar" :src="userInfo.avatar" class="avatar">
            <el-button size="small">Click to upload new avatar</el-button>
          </el-upload>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="updateUserInfo">Update Information</el-button>
        </el-form-item>
      </el-form>
    </div>
  </el-card>
</template>

<script setup>
import { ref } from 'vue';
import { ElCard, ElForm, ElFormItem, ElInput, ElButton, ElUpload } from 'element-plus';
import axios from 'axios';

const userInfo = ref({
  username: '',
  nike_name: '',
  avatar: '',
  email: '',
  country: '',
  province: '',
  city: ''
});

const getUserInfo = async () => {
  try {
    const response = await axios.get('/api/user/info');
    if (response.data.code === 0) {
      userInfo.value = response.data.data;
    } else {
      console.error('Failed to fetch user info:', response.data.msg);
    }
  } catch (error) {
    console.error('Error fetching user information:', error);
  }
};

const updateUserInfo = async () => {
  // Update user info code here
};

const handleAvatarSuccess = (response, file) => {
  userInfo.value.avatar = URL.createObjectURL(file.raw);
};

const beforeAvatarUpload = (file) => {
  const isJPG = file.type === 'image/jpeg' || file.type === 'image/png';
  if (!isJPG) {
    this.$message.error('Avatar picture must be in JPG or PNG format!');
    return false;
  }
  const isLt2M = file.size / 1024 / 1024 < 2;
  if (!isLt2M) {
    this.$message.error('Avatar size must not exceed 2MB!');
    return false;
  }
  return true;
};

getUserInfo();
</script>

<style scoped>
.update-user-info-card {
  max-width: 600px;
  margin: 50px auto;
}

.form-container {
  padding: 20px;
}

.avatar {
  width: 100px;
  height: 100px;
  border-radius: 50%;
}
</style>
