<script setup lang="ts">
import { ref, onMounted } from "vue";
import type { StatsResponse } from "@/responses.ts";
import StatBox from "@/components/StatBox.vue";

const stats = ref<StatsResponse>();

async function fetchStats() {
  try {
    const response = await fetch("api/stats");
    if (!response.ok) {
      throw new Error(`HTTP error: status: ${response.status}`);
    }

    const data: StatsResponse = await response.json();
    console.log("Got stats response:", data);

    stats.value = data;
  } catch (error: unknown) {
    console.error(error);
  }
}

onMounted(() => {
  fetchStats();
});
</script>

<template>
  <div>
    <h3>StatsBox</h3>
    <div v-if="stats != undefined">
      <StatBox title="CPU Info" :stat="stats.cpuInfo"></StatBox>
      <StatBox title="Disk" :stat="stats.diskInfo"></StatBox>
      <StatBox title="Host Info" :stat="stats.hostInfo"></StatBox>
      <StatBox title="Memory" :stat="stats.memInfo"></StatBox>
    </div>
    <div v-else>
      <p>loading...</p>
    </div>
  </div>
</template>

<style scoped></style>
