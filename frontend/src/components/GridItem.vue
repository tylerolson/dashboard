<script setup lang="ts">
const { title } = defineProps<{
  title: string;
}>();
</script>

<template>
  <div class="col-span-4 rounded-xl border border-zinc-700 bg-zinc-800 p-4 shadow">
    <p class="pb-2 text-zinc-200">{{ title }}</p>
    <slot />
  </div>
</template>
