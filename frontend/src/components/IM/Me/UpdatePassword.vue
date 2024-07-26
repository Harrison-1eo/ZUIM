<template>
    <el-card class="update-password-card">
        <h2>修改密码</h2>
        <el-divider></el-divider>
        <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="120px" class="demo-ruleForm">
            <el-form-item label="原密码" prop="oriPass">
                <el-input type="password" v-model="ruleForm.oriPass" autocomplete="off" show-password></el-input>
            </el-form-item>
            <el-form-item label="密码" prop="pass">
                <el-input type="password" v-model="ruleForm.pass" autocomplete="off" show-password></el-input>
            </el-form-item>
            <el-form-item label="确认密码" prop="checkPass">
                <el-input type="password" v-model="ruleForm.checkPass" autocomplete="off" show-password></el-input>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="submitForm('ruleForm')">提交</el-button>
                <el-button @click="resetForm('ruleForm')">重置</el-button>
            </el-form-item>
        </el-form>
    </el-card>
</template>

<style scoped>
.update-password-card {
    width: 500px;
    margin: 40px auto;
    padding: 20px;
    border-radius: 10px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}
</style>

<script>
import axios_config from "@/utils/axios-config";
import {ElDivider, ElMessage} from "element-plus";

export default {
    name: 'UpdatePassword',
    components: {ElDivider},
    data() {
        return {
            ruleForm: {
                oriPass: '',
                pass: '',
                checkPass: '',
            },
            rules: {
                oriPass: [
                    { validator: this.validateOriPass, trigger: 'blur' }
                ],
                pass: [
                    { validator: this.validatePass, trigger: 'blur' }
                ],
                checkPass: [
                    { validator: this.validateCheckPass, trigger: 'blur' }
                ],
            }
        };
    },

    methods: {
        validateOriPass(rule, value) {
            return new Promise((resolve, reject) => {
                if (value === '') {
                    reject(new Error('请输入原密码'));
                } else {
                    resolve();
                }
            });
        },

        validatePass(rule, value) {
            return new Promise((resolve, reject) => {
                if (value === '') {
                    reject(new Error('请输入密码'));
                } else {
                    resolve();
                }
            });
        },

        validateCheckPass(rule, value) {
            return new Promise((resolve, reject) => {
                if (value === '') {
                    reject(new Error('请再次输入密码'));
                } else if (value !== this.ruleForm.pass) {
                    reject(new Error('两次输入密码不一致!'));
                } else {
                    resolve();
                }
            });
        },
        submitForm(formName) {
            this.$refs[formName].validate().then(valid => {
                if (valid) {
                    axios_config.post('/api/user/update_password', {
                        username: JSON.parse(localStorage.getItem('user')),
                        old_password: this.ruleForm.oriPass,
                        new_password: this.ruleForm.pass
                    }).then(res => {
                        if (res.data.code === 0) {
                            ElMessage.success('修改密码成功');
                            localStorage.removeItem('token');
                            this.$router.push('/login');
                        } else {
                            ElMessage.error(res.data.msg);
                        }
                    }).catch(err => {
                        ElMessage.error('修改密码失败，连接服务器错误');
                    });
                } else {
                    ElMessage.error('请检查输入是否正确');
                }
            }).catch(e => {
                ElMessage.error(e.checkPass[0].message || '验证过程中出错');
            });
        },
        resetForm(formName) {
            this.$refs[formName].resetFields();
        },
        login() {
            this.$router.push('/login');
        }
    }
}
</script>