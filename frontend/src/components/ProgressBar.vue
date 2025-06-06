<script setup lang="ts">
const { progress } = defineProps<{
  progress: number;
  unit: string;
}>();

function getBarColor(progress: number) {
  if (progress <= 50) return "bg-green-500 text-green-100";
  if (progress <= 70) return "bg-yellow-500 text-yellow-50";
  if (progress <= 85) return "bg-orange-400 text-orange-100";
  return "bg-red-600 text-red-100";
}
</script>

<template>
  <div class="mt-1 mb-4 h-4 w-full overflow-hidden rounded-full bg-zinc-700">
    <div
      class="h-full bg-green-500 p-0.5 text-center text-xs leading-none font-medium transition-all duration-500 ease-out"
      :class="getBarColor(progress)"
      :style="{ width: progress + '%' }"
    >
      {{ progress }}{{ unit }}
    </div>
  </div>
</template>
