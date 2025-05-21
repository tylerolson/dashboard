<script setup lang="ts">
import GridItem from "@/components/GridItem.vue";
import { camelToTitleCase, convertSeconds } from "@/utils.ts";
import type { HostInfo } from "@/responses.ts";

const { stat } = defineProps<{
  stat: HostInfo;
}>();
</script>

<template>
  <GridItem title="Host" :stat="stat">
    <div v-for="(value, key) in stat" :key="key">
      <p v-if="key == 'uptime'">
        {{ camelToTitleCase(key) }}
        <span class="float-right">{{ convertSeconds(Number(value)) }}</span>
      </p>

      <p v-else-if="key == 'bootTime'">
        {{ camelToTitleCase(key) }}
        <span class="float-right">{{ new Date(Number(value) * 1000).toLocaleString() }}</span>
      </p>
      <p v-else>
        {{ camelToTitleCase(key) }} <span class="float-right">{{ value }}</span>
      </p>
    </div>
  </GridItem>
</template>
