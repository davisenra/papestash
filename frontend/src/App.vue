<script setup lang="ts">
import WallpaperItem from './components/WallpaperItem.vue';
import AppLayout from './layouts/AppLayout.vue';
</script>

<template>
  <AppLayout>
    <div class="grid grid-cols-3 gap-3">
      <WallpaperItem />
      <WallpaperItem />
      <WallpaperItem />
      <WallpaperItem />
      <WallpaperItem />
      <WallpaperItem />
      <WallpaperItem />
    </div>
  </AppLayout>
</template>
