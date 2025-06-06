<script setup lang="ts">
import GridItem from "@/components/GridItem.vue";
import ProgressBar from "@/components/ProgressBar.vue";
import type { TemperatureStat } from "@/responses.ts";

const { stat } = defineProps<{
  stat: TemperatureStat;
}>();
</script>

<template>
  <GridItem :title="stat.sensorKey">
    <p class="text-lg">
      <span v-if="'usedGbs' in stat && 'totalGbs' in stat">
        {{ stat.usedGbs }} GB of {{ stat.totalGbs }} GB used
      </span>
      <span v-else></span>
    </p>
    <ProgressBar :progress="stat.temperature" unit="°C"></ProgressBar>
  </GridItem>
</template>
