<template>
    <div class="statistic-box">
        <div class="statistic-info">
            <div class="statistic-title">
                {{ title }}
            </div>
            <div class="statistic-value">
                {{ value }}
            </div>
        </div>
        <div class="statistic-icon-box">
            <div class="statistic-icon-bg">
                <div class="statistic-icon">
                    <el-icon size="30px" color="#ffffff">
                        <component :is="icon" />
                    </el-icon>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'StatisticItem',
    props: {
        title: {
            type: String,
            default: 'Title'
        },
        value: {
            type: Number,
            default: 0
        },
        icon: {
            type: String,
            default: 'Avatar'
        }
    }
}

</script>

<style>
.statistic-box {
    width: 100%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 10px 20px;
    border-radius: 10px;
    box-shadow: 0 4px 8px rgba(0,0,0,0.1);
    background-color: #ffffff;
    margin: 10px 30px;
}

.statistic-info {
    margin: 5px 0 0 0;
    display: flex;
    flex-direction: column;
}

.statistic-title {
    font-size: 16px;
    color: #8b8a8a;
}

.statistic-value {
    font-size: 40px;
    font-weight: bold;
    color: #333;
}

.statistic-icon-box {
    padding: 10px;
    justify-content: center;
    align-items: center;
}

.statistic-icon-bg {
    width: 65px;
    height: 65px;
    border-radius: 15px;
    background: linear-gradient(322deg,#1f6ff5 7.93%,#9a7ffc 92.22%);
}

.statistic-icon {
    /* 图标上下左右居中 */
    height: 100%;
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
}

</style>