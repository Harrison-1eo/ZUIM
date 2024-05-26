<template>
    <div class="im-home">
        <div class="title-box">
            <div class="title-box-content">
                <h1>{{ greeting }}，</h1>
            </div>
            <div class="user-avatar">
                <el-avatar :size="60" :src="getAvatar(userInfo.avatar)" />
            </div>
            <div class="title-box-content">
                <h1>{{ userInfo.username }}。</h1>
            </div>
        </div>

        <el-divider />

        <div class="row-box">
            <StatisticItem class="pointer-emphasis" :title="'在线终端'" :value="statsInfo.online_user_count" :icon="Avatar" />
            <StatisticItem class="pointer-emphasis" :title="'警告消息'" :value="2" :icon="WarningFilled" />
            <StatisticItem class="pointer-emphasis" :title="'错误消息'" :value="0" :icon="CircleCloseFilled" />
        </div>
        <div class="row-box">
            <StatisticItem class="pointer-emphasis" :title="'终端总数'" :value="statsInfo.friend_count" :icon="Grid" />
            <StatisticItem class="pointer-emphasis" :title="'房间总数'" :value="statsInfo.room_count" :icon="PhoneFilled" />
            <StatisticItem class="pointer-emphasis" :title="'未读消息'" :value="statsInfo.unread_message_count" :icon="Comment" />
        </div>

        <el-divider />

        <div class="home-statistic-box">
            <div class="home-statistic-left-box">
                <h1>具体统计信息：</h1>
                <div class="table-box">
                    <h3>房间消息数量统计</h3>
                    <table class="table">
                        <thead>
                        <tr>
                            <th>序号</th>
                            <th>房间名称</th>
                            <th>消息数量</th>
                        </tr>
                        </thead>
                        <tbody>
                        <tr v-for="item in tableData" :key="item.no">
                            <td>{{ item.no }}</td>
                            <td>{{ item.name }}</td>
                            <td>{{ item.number }}</td>
                        </tr>
                        </tbody>
                    </table>
<!--                    <el-table :data="tableData" width="450" stripe>-->
<!--                        <el-table-column type="index" label="序号" />-->
<!--                        <el-table-column prop="name" label="房间名称" />-->
<!--                        <el-table-column prop="number" label="消息数量" />-->
<!--                    </el-table>-->

                </div>

            </div>

            <div class="chart-box">
                <!--            <div class="chart-title">-->
                <!--                <h2>消息统计</h2>-->
                <!--            </div>-->
                <div class="chart-chart pointer-emphasis">
                    <div ref="chart" style="width: 600px; height: 400px;"></div>
                </div>
            </div>
        </div>




    </div>
</template>

<script>
import axios_config from "@/utils/axios-config";
import StatisticItem from "@/components/IM/Home/StatisticItem.vue";
import {Avatar, CircleCloseFilled, Grid, WarningFilled, Comment, PhoneFilled} from "@element-plus/icons";
import {ElMessage} from "element-plus";
import * as echarts from 'echarts';
import {backendBaseUrl} from "@/utils/base-url-setting";

export default {
    name: 'IMHome',
    computed: {
        PhoneFilled() {
            return PhoneFilled
        },
        Comment() {
            return Comment
        },
        Grid() {
            return Grid
        },
        CircleCloseFilled() {
            return CircleCloseFilled
        },
        Avatar() {
            return Avatar
        },
        WarningFilled() {
            return WarningFilled
        },
    },
    components: {StatisticItem},
    data() {
        return {
            statsInfo: {},
            userInfo: {},
            message: '欢迎来到首页',
            greeting: '',
            tableData: [
                {
                    no: '1',
                    name: '房间1',
                    number: 100,
                },
            ],
        };
    },
    created() {
        this.getStats();
        this.getUserInfo();
        this.greeting = this.getGreeting();
    },
    methods: {
        getStats() {
            axios_config.get('/api/user/stats').then(response => {
                this.statsInfo = response.data.data;
                console.log(response);
                const messageCount = this.statsInfo.message_count;

                const sortedMessageCount = Object.entries(messageCount).sort((a, b) => new Date(b[0]) - new Date(a[0]));
                const sortedRoomMessageCount = Object.entries(this.statsInfo.room_message_count).sort((a, b) => b[1] - a[1]);
                console.log(sortedMessageCount);
                console.log(sortedRoomMessageCount);
                this.statsInfo.message_count = sortedMessageCount;
                this.statsInfo.room_message_count = sortedRoomMessageCount;

                this.renderChart();
                const tableData = [];
                for (const [key, value] of Object.entries(this.statsInfo.room_message_count)) {
                    tableData.push({
                        no: key,
                        name: value[0],
                        number: value[1],
                    });
                }
                this.tableData = tableData;
            }).catch(error => {
                console.error('Error fetching stats:', error);
                ElMessage.error('获取统计信息失败');
            });
        },
        getUserInfo() {
            axios_config.get('/api/user/my').then(response => {
                this.userInfo = response.data.data;
            }).catch(error => {
                console.error('Error fetching user info:', error);
                ElMessage.error('获取用户信息失败');
            });
        },
        getGreeting() {
            const now = new Date();
            const hours = now.getHours();
            let greeting;

            if (hours >= 6 && hours < 12) {
                greeting = "早上好";
            } else if (hours >= 12 && hours < 14) {
                greeting = "中午好";
            } else if (hours >= 14 && hours < 18) {
                greeting = "下午好";
            } else {
                greeting = "晚上好";
            }

            return greeting;
        },
        getAvatar(url) {
            // console.log(url);
            if (url) {
                return backendBaseUrl + url;
                // return this.avatarUrl(url);
            }
            return backendBaseUrl + '/static/avatars/nopic.png';
        },
        renderChart() {
            const sortedMessageCount = this.statsInfo.message_count.reverse();

            const dates = sortedMessageCount.map(item => item[0]);
            const counts = sortedMessageCount.map(item => item[1]);

            const chart = echarts.init(this.$refs.chart);

            const option = {
                title:{
                    text:'发送信息数量统计图',
                    x:'center',
                    y:'top',
                    // textAlign:'center'
                },
                tooltip: {
                    trigger: 'axis'
                },
                xAxis: {
                    type: 'category',
                    data: dates
                },
                yAxis: {
                    type: 'value'
                },
                series: [
                    {
                        data: counts,
                        type: 'line',
                        smooth: true
                    }
                ]
            };
            chart.setOption(option);
        },
    }
};
</script>

<style scoped>
h1 {
    font-size: 3em;
    font-weight: bolder;
    margin: 0;
}

.im-home {
    margin: 20px 60px;
}

.title-box {
    display: flex;
    align-items: center;
    margin: 40px 0 20px 0 ;
}

.user-avatar {
    margin: 0 10px;
}

.row-box {
    display: flex;
    justify-content: space-between;
    margin: 20px 0;
}

.home-statistic-box {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
}

.table-box {
    width: 100%;
    margin: 20px;
    padding: 20px;
    background: #ffffff;
    border-radius: 10px;
    display: flex;
    flex-direction: column;
    align-items: center;
}

.table {
    width: 100%;
    text-align: center;
}

.chart-box {
    margin: 20px;
    display: flex;
    flex-direction: column;
    align-items: center;
}

.chart-title {
    margin: 20px 0;
}

.chart-chart {
    display: flex;
    justify-content: center;
    background: #ffffff;
    border-radius: 10px;
    box-shadow: 0 4px 8px rgba(0,0,0,0.1);
    padding: 20px 0 0 0;
}

.pointer-emphasis {
    border: 1px solid #ffffff;
    transition: border 0.5s linear;
}

.pointer-emphasis:hover {
    border: 1px solid #4955f5;
}

</style>
