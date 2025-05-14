<script setup lang="ts">
import { ref, onMounted } from "vue";
import type { StatsResponse } from "@/responses.ts";

const stats = ref<StatsResponse>();

async function fetchStats() {
  try {
    const response = await fetch("api/stats");
    if (!response.ok) {
      throw new Error(`HTTP error: status: ${response.status}`);
    }

    const data: StatsResponse = await response.json();
    console.log(data);

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
      <pre>CPU Info: {{ stats.cpu_info }}</pre>
      <pre>CPU Percent: {{ stats.cpu_percent }}</pre>
      <pre>Disk: {{ stats.disk }}</pre>
      <pre>Host Info: {{ stats.host_info }}</pre>
      <pre>Memory: {{ stats.memory }}</pre>
    </div>
    <div v-else>
      <p>loading...</p>
    </div>
  </div>
</template>

<style scoped></style>
