<template>
    <div class="setting-box">
        <div class="setting-title-box">
            <FluentCommunicationShield24Filled size="5em"/>
            <span class="setting-title">IoT Messaging and Domestic Security Base (IMDSB) System</span>
        </div>


    </div>
</template>

<script>
import FluentCommunicationShield24Filled from "@/components/IM/Frame/IMIcon.vue";

export default {
  name: 'IMSetting',
    components: {
        FluentCommunicationShield24Filled,
    },
  data() {
    return {
      notificationEnabled: true
    };
  }
};
</script>

<style scoped>
.setting-box {
    margin: 20px;
}
.setting-title-box {
    display: flex;
    align-items: center;
    flex-direction: row;
    margin-bottom: 20px;
}
.setting-title {
    font-size: 2em;
    font-weight: bold;
    margin-left: 10px;
}
</style>
