<script setup lang="ts">
import WallpaperItem from './components/WallpaperItem.vue';
import AppLayout from './layouts/AppLayout.vue';
</script>

<template>
  <AppLayout>
    <div class="flex space-x-3">
      <div class="flex items-center space-x-2">
        <label for="perPage">Results per page</label>
        <select name="perPage" class="bg-white px-2 py-1.5">
          <option value="25">25</option>
          <option value="50">50</option>
          <option value="75">75</option>
          <option value="100">100</option>
        </select>
      </div>
      <div class="flex items-center space-x-2">
        <label for="ratio">Ratio</label>
        <select name="ratio" class="bg-white px-2 py-1.5">
          <option value=""></option>
          <option value="16:9">16:9</option>
          <option value="16:10">16:10</option>
          <option value="4:3">4:3</option>
          <option value="3:3">3:3</option>
          <option value="9:16">9:16</option>
        </select>
      </div>
      <div class="flex items-center space-x-2">
        <label>Resolution</label>
        <input name="height" class="w-16 bg-white px-2 py-1.5" />
        <span>x</span>
        <input name="width" class="w-16 bg-white px-2 py-1.5" />
      </div>
    </div>
    <div class="my-3 grid grid-cols-3 gap-3">
      <WallpaperItem />
      <WallpaperItem />
      <WallpaperItem />
      <WallpaperItem />
      <WallpaperItem />
      <WallpaperItem />
      <WallpaperItem />
    </div>
    <div class="mb-6 flex space-x-2">
      <button class="rounded-md bg-white px-3 py-1.5">Previous</button>
      <button class="rounded-md bg-white px-3 py-1.5">1</button>
      <button class="rounded-md bg-white px-3 py-1.5">2</button>
      <button class="rounded-md bg-white px-3 py-1.5">3</button>
      <button class="rounded-md bg-white px-3 py-1.5">Next</button>
    </div>
  </AppLayout>
</template>
