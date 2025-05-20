import { defineStore } from "pinia";
import type { StatsResponse } from "@/responses.ts";
import { ref } from "vue";

export const useStatStore = defineStore("statStore", () => {
  const stats = ref<StatsResponse>();

  async function fetchStats() {
    try {
      const response: Response = await fetch("api/stats");
      if (!response.ok) {
        throw new Error(`HTTP error: status: ${response.status}`);
      }

      const data: StatsResponse = await response.json();
      console.log("Got stats response:", data);

      stats.value = data;
    } catch (error: unknown) {
      console.error(error);
    }
  }

  return { stats, fetchStats };
});
