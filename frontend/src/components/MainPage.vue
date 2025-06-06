<script setup lang="ts">
import StatItem from "@/components/StatItem.vue";
import HostInfoItem from "@/components/HostInfoItem.vue";
import { useStatStore } from "@/stores/statStore.ts";
import { storeToRefs } from "pinia";
import TempItem from "./TempItem.vue";

const statStore = useStatStore();
const { stats } = storeToRefs(statStore);
</script>

<template>
  <div>
    <div
      v-if="stats != undefined"
      class="grid grid-cols-4 gap-5 p-5 md:grid-cols-8 lg:grid-cols-12"
    >
      <StatItem title="Cpu" :stat="stats.cpuStat"></StatItem>
      <StatItem title="Disk" :stat="stats.diskStat"></StatItem>
      <StatItem title="Memory" :stat="stats.memStat"></StatItem>
      <TempItem v-for="temp in stats.tempStat" :stat="temp" :key="temp.sensorKey"></TempItem>
    </div>
    <div
      v-if="stats != undefined"
      class="grid grid-cols-4 gap-5 px-5 md:grid-cols-8 lg:grid-cols-12"
    >
      <HostInfoItem :stat="stats.hostInfo"></HostInfoItem>
    </div>
    <div v-else>
      <p>loading...</p>
    </div>
  </div>
</template>
