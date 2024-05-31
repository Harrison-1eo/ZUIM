<template>
    <div class="im-home">
        <div class="welcome-box">
            <HomeWelcome :userInfo="userInfo" />
        </div>

        <el-divider />

        <div class="row-box">
            <StatisticItem :title="'在线终端'" :value="statsInfo.online_user_count" :icon="Avatar" :iconColor="'#40c9c6'" />
            <StatisticItem :title="'终端总数'" :value="statsInfo.friend_count" :icon="Grid" :iconColor="'#0fbe7d'" />
            <StatisticItem :title="'参与房间'" :value="statsInfo.room_count" :icon="PhoneFilled" :iconColor="'#36a3f7'" />
            <StatisticItem :title="'未读消息'" :value="statsInfo.unread_message_count" :icon="Comment" :iconColor="'#f4516c'" />
        </div>

        <div class="line-chart-box">
            <div class="line-chart-chart pointer-emphasis">
                <div ref="lineChart" style="width: 100%; height: 300px;"></div>
            </div>
        </div>

        <div class="pie-chart-info-box">
            <div class="pie-chart-chart pointer-emphasis">
                <div ref="pieChart" style="width: 400px; height: 400px;"></div>
            </div>
            <InfoBox />
        </div>

    </div>
</template>

<script>
import axios_config from "@/utils/axios-config";
import StatisticItem from "@/components/IM/Home/StatisticItem.vue";
import HomeWelcome from "@/components/IM/Home/HomeWelcome.vue";
import InfoBox from "@/components/IM/Home/InfoBox.vue";
import { Avatar, CircleCloseFilled, Grid, WarningFilled, Comment, PhoneFilled } from "@element-plus/icons";
import { ElMessage } from "element-plus";
import * as echarts from 'echarts';

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
    components: { StatisticItem, InfoBox, HomeWelcome },
    data() {
        return {
            statsInfo: {},
            userInfo: {},
            message: '欢迎来到首页',
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
    },
    methods: {
        getStats() {
            axios_config.get('/api/user/stats').then(response => {
                this.statsInfo = response.data.data;

                const messageCount = this.statsInfo.message_count;

                const sortedMessageCount = Object.entries(messageCount).sort((a, b) => new Date(b[0]) - new Date(a[0]));
                const sortedRoomMessageCount = Object.entries(this.statsInfo.room_message_count).sort((a, b) => b[1] - a[1]);
                this.statsInfo.message_count = sortedMessageCount;
                this.statsInfo.room_message_count = sortedRoomMessageCount;

                const tableData = [];
                for (const [key, value] of Object.entries(this.statsInfo.room_message_count)) {
                    tableData.push({
                        no: key,
                        name: value[0],
                        number: value[1],
                    });
                }
                this.tableData = tableData;

                this.renderLineChart();
                this.renderPieChar();
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
        renderLineChart() {
            const sortedMessageCount = this.statsInfo.message_count.reverse();

            const dates = sortedMessageCount.map(item => item[0]);
            const counts = sortedMessageCount.map(item => item[1]);

            const chart = echarts.init(this.$refs.lineChart);

            const option = {
                title: {
                    text: '终端流量统计图',
                    x: 'center',
                    y: 'top',
                    // textAlign:'center'
                },
                tooltip: {
                    trigger: 'axis'
                },
                xAxis: {
                    name: '时间',
                    type: 'category',
                    data: dates
                },
                yAxis: {
                    name: '消息流量(KB)',
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
        renderPieChar() {
            const top4 = this.statsInfo.room_message_count.slice(0, 4);
            const otherTotal = this.statsInfo.room_message_count.slice(4).reduce((sum, item) => sum + item[1], 0);

            const data = top4.map(item => ({ value: item[1], name: item[0] }));
            if (otherTotal > 0) {
                data.push({ value: otherTotal, name: '其他' });
            }

            const chart = echarts.init(this.$refs.pieChart);
            const option = {
                title: {
                    text: '房间消息数量统计图',
                    x: 'center',
                    y: 'top',
                    // textAlign:'center'
                },
                tooltip: {
                    trigger: 'item',
                    formatter: '{a} <br/>{b} : {c} ({d}%)'
                },
                legend: {
                    left: 'center',
                    bottom: '40',
                    data: data.map(item => item.name)
                },
                series: [
                    {
                        name: '房间消息数量',
                        type: 'pie',
                        roseType: 'radius',
                        radius: [15, 95],
                        center: ['50%', '38%'],
                        data: data,
                        animationEasing: 'cubicInOut',
                        animationDuration: 2600
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

.welcome-box {
    margin: 40px 0 20px 30px;
}

.row-box {
    display: flex;
    justify-content: space-between;
    margin: 20px 0;
}

.line-chart-box {
    margin: 30px;
}

.pie-chart-info-box {
    display: flex;
    justify-content: space-between;
    flex-direction: row;
}

.line-chart-chart,
.pie-chart-chart {
    display: flex;
    justify-content: center;
    background: #ffffff;
    border-radius: 10px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    padding: 20px 0 0 0;
}

.pie-chart-chart {
    margin: 15px 30px 20px 30px;
    padding: 20px 20px 0 20px;
}

.pointer-emphasis {
    border: 1px solid #ffffff;
    transition: border 0.5s linear;
}

.pointer-emphasis:hover {
    border: 1px solid #4955f5;
}
</style>
