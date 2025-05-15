<script setup lang="ts">
import type { CpuInfo, DiskInfo, HostInfo, MemInfo } from "@/responses.ts";

const { stat } = defineProps<{
  title: string;
  stat: CpuInfo | MemInfo | HostInfo | DiskInfo;
}>();

function camelToTitleCase(input: string) {
  const result = input.replace(/([A-Z])/g, " $1");
  return result.charAt(0).toUpperCase() + result.slice(1);
}
</script>

<template>
  <div>
    <h6>{{ title }}</h6>
    <div v-for="(value, key) in stat" :key="key">
      {{ camelToTitleCase(key) }} - {{ value }}
      <span></span>
    </div>
  </div>
</template>

<style scoped></style>
