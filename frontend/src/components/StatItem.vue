<script setup lang="ts">
import GridItem from "@/components/GridItem.vue";
import ProgressBar from "@/components/ProgressBar.vue";
import type { CpuStat, DiskStat, MemStat } from "@/responses.ts";

const { stat } = defineProps<{
  title: string;
  stat: CpuStat | DiskStat | MemStat;
}>();
</script>

<template>
  <GridItem :title :stat="stat">
    <p class="text-lg">
      <span v-if="'usedGbs' in stat && 'totalGbs' in stat">
        {{ stat.usedGbs }} GB of {{ stat.totalGbs }} GB used
      </span>
      <span v-else> Cpu Usage </span>
    </p>
    <ProgressBar :progress="stat.usedPercent"></ProgressBar>
  </GridItem>
</template>
