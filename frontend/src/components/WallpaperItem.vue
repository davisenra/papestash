<script setup lang="ts">
import { Icon } from '@iconify/vue';
</script>

<template>
  <div class="flex cursor-pointer flex-col rounded-md border px-2 pt-2">
    <img class="h-max" src="https://placehold.co/1920x1080" />
    <div class="my-1 flex space-x-1">
      <button>
        <Icon icon="ic:baseline-info" class="text-xl text-neutral-400 hover:text-neutral-500" />
      </button>
      <button>
        <Icon icon="ic:baseline-download" class="text-xl text-neutral-400 hover:text-neutral-500" />
      </button>
    </div>
  </div>
</template>
