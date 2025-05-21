<script setup lang="ts">
import MainPage from "@/components/MainPage.vue";
import { useStatStore } from "@/stores/statStore.ts";
import { onMounted, onUnmounted } from "vue";

const useStats = useStatStore();

onMounted(() => {
  useStats.startInterval();
});

onUnmounted(() => {
  useStats.stopInterval();
});
</script>

<template>
  <div>
    <MainPage></MainPage>
  </div>
</template>
