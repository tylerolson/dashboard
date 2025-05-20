<script setup lang="ts">
import MainPage from "@/components/MainPage.vue";
import { useStatStore } from "@/stores/statStore.ts";
import { onMounted } from "vue";

const useStats = useStatStore();

onMounted(() => {
  useStats.fetchStats();
  setInterval(function () {
    useStats.fetchStats();
  }, 3000);
});
</script>

<template>
  <div>
    <MainPage></MainPage>
  </div>
</template>
