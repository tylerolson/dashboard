<script setup lang="ts">
import { ref, onMounted } from "vue";
import type { StatsResponse } from "@/responses.ts";
import StatsItem from "@/components/StatsItem.vue";

const stats = ref<StatsResponse>();

const width = window.innerWidth;
const height = window.innerHeight;

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
  <div class="h-full">
    <h3>StatsBox {{ width }} x {{ height }}</h3>
    <div
      v-if="stats != undefined"
      class="grid grid-cols-4 gap-5 px-5 md:grid-cols-8 lg:grid-cols-12"
    >
      <StatsItem title="CPU Info" :stat="stats.cpuInfo"></StatsItem>
      <StatsItem title="Disk" :stat="stats.diskInfo"></StatsItem>
      <StatsItem title="Host Info" :stat="stats.hostInfo"></StatsItem>
      <StatsItem title="Memory" :stat="stats.memInfo"></StatsItem>
    </div>
    <div v-else>
      <p>loading...</p>
    </div>
  </div>
</template>

<style scoped></style>
