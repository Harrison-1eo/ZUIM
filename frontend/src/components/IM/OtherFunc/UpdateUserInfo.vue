<template>
  <el-upload
      class="avatar-uploader"
      action="https://jsonplaceholder.typicode.com/posts/"
      :show-file-list="false"
      :on-success="handleAvatarSuccess"
      :before-upload="beforeAvatarUpload">
    <img v-if="imageUrl" :src="imageUrl" class="avatar" alt="+">
    <i v-else class="el-icon-plus avatar-uploader-icon"></i>
  </el-upload>
</template>

<script setup>
import { ref } from 'vue';
import { ElCard, ElForm, ElFormItem, ElInput, ElButton, ElUpload } from 'element-plus';
import axios from "@/axios-config";

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
  try {
    const response = await axios.post('/api/user/update', {
      user_id: userInfo.value.user_id,
      username: userInfo.value.username,
      nike_name: userInfo.value.nike_name,
      avatar: userInfo.value.avatar,
      sexuality: userInfo.value.sexuality,
      year: userInfo.value.year,
      month: userInfo.value.month,
      day: userInfo.value.day,
      country: userInfo.value.country,
      province: userInfo.value.province,
      city: userInfo.value.city,
      email: userInfo.value.email
    });
    if (response.data.code === 0) {
      this.$message.success('User information updated successfully!');
    } else {
      this.$message.error('Failed to update user information: ' + response.data.msg);
    }
  } catch (error) {
    console.error('Error updating user information:', error);
    this.$message.error('Failed to update user information. Please try again later.');
  }
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
.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}
.avatar {
  width: 178px;
  height: 178px;
  display: block;
}
</style>

<script>
export default {
  data() {
    return {
      imageUrl: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png'
    };
  },
  methods: {
    handleAvatarSuccess(res, file) {
      this.imageUrl = URL.createObjectURL(file.raw);
    },
    beforeAvatarUpload(file) {
      const isJPG = file.type === 'image/jpeg';
      const isLt2M = file.size / 1024 / 1024 < 2;

      if (!isJPG) {
        this.$message.error('上传头像图片只能是 JPG 格式!');
      }
      if (!isLt2M) {
        this.$message.error('上传头像图片大小不能超过 2MB!');
      }
      return isJPG && isLt2M;
    }
  }
}
</script>