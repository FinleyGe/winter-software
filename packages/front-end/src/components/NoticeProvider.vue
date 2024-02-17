<script setup lang="ts">
import { NoticeKey, Notice } from "@components/Notice";
import { provide, reactive} from "vue";
import { Check, X, ExclamationMark} from "@vicons/tabler";

const duration = "2s";

const noticeData = reactive<{
  type: string;
  message: string;
  state: "in" | "out";
}[]>([]);

function notice (notice: Notice){
  noticeData.push({
    ...notice,
    state: "in"
  });

  setTimeout(() => {
    noticeData[0].state = "out";
    noticeData.shift();
  }, 2000);
}

provide(NoticeKey, notice);

</script>
<template>
  <slot />
  <div
    v-show="noticeData.length > 0"
    class="notice"
  >
    <div
      v-for="item in noticeData"
      :key="item.message"
      class="notice-item"
      :class="[item.type]"
    >
      <div class="icon">
        <Check v-if="item.type == 'success'" />
        <X v-if="item.type == 'error'" />
        <ExclamationMark v-if="item.type == 'info'" />
      </div>
      {{ item.message }}
    </div>
  </div>
</template>

<style scoped lang="scss">
.notice {
  position: fixed;
  width: 300px;
  top: 20px;
  right: 20px;
  margin: auto;
  z-index: 1000;
  .notice-item {
    padding: 10px 20px;
    margin-bottom: 10px;
    border-radius: 5px;
    color: #fff;
    display: flex;
    flex-direction: row;
    &.success {
      background-color: #67c23a;
    }
    &.warning {
      background-color: #e6a23c;
    }
    &.error {
      background-color: #f56c6c;
    }
    &.info {
      background-color: #909399;
    }
    animation: slideIn v-bind("duration");
    transition: all 0.5s;
  }
}
.icon {
  width: 24px;
  margin-right: 10px;
}
@keyframes slideIn {
  0% {
    transform: translateX(100%);
    opacity: 0;
  }
  10% {
    transform: translateX(0%);
    opacity: 1;
  }

  90% {
    transform: translateX(0%);
    opacity: 1;
  }

  100%{
    transform: translateX(100%);
    opacity: 0;
  }
}

</style>
