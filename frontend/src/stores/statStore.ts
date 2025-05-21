import { defineStore } from "pinia";
import type { StatsResponse } from "@/responses.ts";
import { ref } from "vue";

export const useStatStore = defineStore("statStore", () => {
  const stats = ref<StatsResponse>();
  let intervalId = 0;

  async function fetchStats() {
    try {
      const response: Response = await fetch("api/stats");
      if (!response.ok) {
        console.error(`HTTP error: status: ${response.status}`);
        return;
      }

      const data: StatsResponse = await response.json();
      console.log("Got stats response:", data);

      stats.value = data;
    } catch (error: unknown) {
      console.error(error);
    }
  }

  function startInterval(timeMs: number = 3000) {
    if (intervalId != 0) {
      return;
    }

    fetchStats();
    intervalId = setInterval(fetchStats, timeMs);
  }

  function stopInterval() {
    if (intervalId != 0) {
      clearInterval(intervalId);
      intervalId = 0;
    }
  }

  return { stats, fetchStats, startInterval, stopInterval };
});
