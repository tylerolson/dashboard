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
  <div class="col-span-4 rounded-xl border border-zinc-700 bg-zinc-900 p-4 shadow">
    <p class="pb-2 text-2xl font-bold">{{ title }}</p>
    <div v-for="(value, key) in stat" :key="key">
      <span class="font-bold">{{ camelToTitleCase(key) }}</span> - {{ value }}
    </div>
  </div>
</template>

<style scoped></style>
