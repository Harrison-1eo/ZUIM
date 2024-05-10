<script setup>

import { ElMessage, ElMessageBox, ElNotification } from 'element-plus'

// 消息
const openMsg = () => {
  ElMessage({
    type: 'success', // success | warning | info | error
    message: 'dengruicode.com',
    showClose: true
  })
}

// 确认框
const openConfirm = () => {
  ElMessageBox.confirm('确认删除?', '标题', {
    type: 'warning',
    confirmButtonText: '确认',
    cancelButtonText: '取消'
  }).then(() => {
    console.log('确认')
  }).catch(() => {
    console.log('取消')
  })
}

// 通知
const openNotifiy = () => {
  ElNotification({
    title: '标题',
    message: '邓瑞编程',
    duration: 1500 // 展示时间 [单位:毫秒]
  })
}

// 通知2
const openNotifiy2 = () => {
  ElNotification({
    type: 'success', // success | warning | info | error
    title: '标题',
    message: 'dengruicode.com',
    duration: 1500,
    position: 'bottom-right'
  })
}

import { ref, reactive } from 'vue'

//默认选中的标签名称
const selectedName = ref("2")

//选中标签触发的回调
const tabClick = (tab, event) => {
  console.log("tab", tab.props, "event", event)
}

const tab = reactive({
  arr: [
    { name: "1", title: '邓瑞', content: '内容1' },
    { name: "2", title: '邓瑞编程', content: '内容2' },
    { name: "3", title: 'dengruicode.com', content: '内容3' },
  ]
})

//添加
const tabAdd = () => {
  let index = tab.arr.length
  index++

  tab.arr.push({
    name: index,
    title: '新选项卡' + index,
    content: '内容' + index
  })
}

//移除
const tabRemove = (name) => {
  console.log("name:", name)

  const index = tab.arr.findIndex((value) => {
    return value.name === name
  })

  tab.arr.splice(index, 1) //移除元素

  const name2 = ref('')
  const password = ref('')
  const content = ref('邓瑞编程')
  const url = ref('dengruicode.com')
  const url2 = ref('dengruicode')
  const email = ref('123456')

  //const selected = ref('')
  const selected = ref('2') //选中的下拉框

}

</script>

<template>
  <div>
    <p>
      If Element Plus is successfully added to this project, you'll see an
      <code v-text="'<el-button>'"></code>
      below
    </p>
    <el-button type="primary">el-button</el-button>
  </div>
  <HelloWorld msg="Welcome to Your Vue.js App"/>
  <h3>按钮</h3>
  <el-button>默认按钮</el-button>
  <el-button type="primary">主要按钮</el-button>
  <el-button type="success">成功按钮</el-button>
  <el-button type="info">信息按钮</el-button>
  <el-button type="warning">警告按钮</el-button>
  <el-button type="danger">危险按钮</el-button>

  <hr>
  <h3>按钮属性</h3>
  <el-button plain>朴素按钮</el-button>
  <el-button round>圆角按钮</el-button>
  <el-button circle>圆</el-button>
  <el-button disabled>禁用按钮</el-button>
  <el-button loading>加载中</el-button>

  <hr>
  <h3>尺寸</h3>
  <el-button size="large">大型按钮</el-button>
  <el-button>默认按钮</el-button>
  <el-button size="small">小型按钮</el-button>

  <h3>图标</h3>
  <el-icon><Plus /></el-icon>
  <el-icon><Edit /></el-icon>
  <el-icon><Delete /></el-icon>
  <el-icon class="is-loading"><Loading /></el-icon>

  <hr>
  <h3>属性</h3>
  <el-icon size="30" color="red"><Search /></el-icon>

  <hr>
  <h3>按钮</h3>
  <el-button type="primary">
    <el-icon><Search /></el-icon>
    <span> 搜索 </span>
  </el-button>

  <el-button type="primary">
    <el-icon><Search /></el-icon>
  </el-button>

  <el-button type="primary" circle>
    <el-icon><Search /></el-icon>
  </el-button>

  <hr>
  <h3>按钮组</h3>
  <el-button-group>
    <el-button type="primary">
      <el-icon><Plus /></el-icon>
    </el-button>

    <el-button type="primary">
      <el-icon><Edit /></el-icon>
    </el-button>

    <el-button type="primary">
      <el-icon><Delete /></el-icon>
    </el-button>
  </el-button-group>

  <hr>
  <h3>消息提示</h3>
  <el-button @click="openMsg">消息</el-button>
  <el-button @click="openConfirm">确认框</el-button>
  <el-button @click="openNotifiy">通知</el-button>
  <el-button @click="openNotifiy2">通知2</el-button>

  <hr>
  <h3>标签页</h3>
  <el-tabs v-model="selectedName" @tab-click="tabClick">
    <el-tab-pane label="邓瑞" name="1">内容1</el-tab-pane>
    <el-tab-pane label="邓瑞编程" name="2">内容2</el-tab-pane>
    <el-tab-pane label="dengruicode.com" name="3">内容3</el-tab-pane>
  </el-tabs>

  <h3>卡片风格</h3>
  <el-tabs v-model="selectedName" @tab-click="tabClick" type="card">
    <el-tab-pane label="邓瑞" name="a">内容1</el-tab-pane>
    <el-tab-pane label="邓瑞编程" name="b">内容2</el-tab-pane>
    <el-tab-pane label="dengruicode.com" name="b">内容3</el-tab-pane>
  </el-tabs>

  <h3>带有边框的卡片风格</h3>
  <el-tabs v-model="selectedName" @tab-click="tabClick" type="border-card">
    <el-tab-pane label="邓瑞" name="A">内容1</el-tab-pane>
    <el-tab-pane label="邓瑞编程" name="B">内容2</el-tab-pane>
    <el-tab-pane label="dengruicode.com" name="C">内容3</el-tab-pane>
  </el-tabs>

  <h3>动态添加</h3>
  <el-button @click="tabAdd">添加</el-button>

  <el-tabs v-model="selectedName" @tab-remove="tabRemove" closable type="card">
    <el-tab-pane v-for="(value, key) in tab.arr" :key="value.name" :label="value.title" :name="value.name">
      {{ value.content }}
    </el-tab-pane>
  </el-tabs>

  <hr>
  <h3>表单</h3>
  <div style="width: 300px;">
    <!-- clearable 可一键清空 -->
    <h3>输入框</h3>
    <el-input v-model="name2" clearable placeholder="请输入用户名" />

    <!-- show-password 可切换显示隐藏密码 -->
    <h3>密码框</h3>
    <el-input v-model="password" show-password placeholder="请输入密码" />

    <h3>文本域</h3>
    <el-input type="textarea" v-model="content" rows="2" />

    <h3>输入内容长度限制 - 输入框</h3>
    <el-input v-model="name2" maxlength="10" show-word-limit />

    <h3>输入内容长度限制 - 文本域</h3>
    <el-input type="textarea" v-model="content" maxlength="20" rows="3" show-word-limit />

    <h3>尺寸</h3>
    大 <el-input size="large" />
    默认 <el-input />
    小 <el-input size="small" />

    <h3>前置</h3>
    <el-input v-model="url">
      <template #prepend>https://</template>
    </el-input>

    <h3>后置</h3>
    <el-input v-model="email">
      <template #append>@qq.com</template>
    </el-input>

    <h3>前置后置</h3>
    <el-input v-model="url2">
      <template #prepend>https://</template>
      <template #append>.com</template>
    </el-input>

    <h3>前置后置扩展 - 搜索</h3>
    <el-input placeholder="请输入课程名称">
      <template #prepend>
        <el-select v-model="selected" placeholder="请选择" style="width: 100px;">
          <el-option label="前端" value="1" />
          <el-option label="后端" value="2" />
          <el-option label="服务端" value="3" />
        </el-select>
      </template>
      <template #append>
        <el-button>
          <el-icon><Search /></el-icon>
        </el-button>
      </template>
    </el-input>
  </div>

</template>

<script>
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
