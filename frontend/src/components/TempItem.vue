<script setup lang="ts">
import GridItem from "@/components/GridItem.vue";
import ProgressBar from "@/components/ProgressBar.vue";
import type { TemperatureStat } from "@/responses.ts";
import { camelToTitleCase } from "@/utils";

const { stat } = defineProps<{
  stat: TemperatureStat;
}>();
</script>

<template>
  <GridItem :title="camelToTitleCase(stat.sensorKey)">
    <ProgressBar :progress="stat.temperature" unit="°C"></ProgressBar>
  </GridItem>
</template>
