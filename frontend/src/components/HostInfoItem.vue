<script setup lang="ts">
import GridItem from "@/components/GridItem.vue";
import { camelToTitleCase } from "@/utils.ts";
import type { HostInfo } from "@/responses.ts";

const { stat } = defineProps<{
  stat: HostInfo;
}>();
</script>

<template>
  <GridItem title="Host" :stat="stat">
    <div v-for="(value, key) in stat" :key="key">
      <p>
        {{ camelToTitleCase(key) }} <span class="float-right">{{ value }}</span>
      </p>
    </div>
  </GridItem>
</template>
