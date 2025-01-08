<script setup lang="ts">
import { Icon } from '@iconify/vue';
</script>

<template>
  <div class="min-h-full">
    <nav class="bg-sky-800">
      <div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div class="flex h-16 items-center justify-between">
          <div class="flex items-center">
            <div class="shrink-0">
              <a href="/" class="inline-flex items-center text-lg font-bold text-white">
                <Icon icon="ic:baseline-photo" class="mr-1 text-2xl" />
                Papestash
              </a>
            </div>
            <div>
              <div class="ml-10 flex items-baseline space-x-4">
                <a
                  href="#"
                  class="rounded-md bg-sky-900 px-3 py-2 text-sm font-medium text-white"
                  aria-current="page"
                >
                  Wallpapers
                </a>
                <a
                  href="#"
                  class="rounded-md px-3 py-2 text-sm font-medium text-gray-300 hover:bg-sky-700 hover:text-white"
                >
                  Logout
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </nav>
    <main>
      <div class="mx-auto max-w-7xl px-4 py-6 sm:px-6 lg:px-8">
        <slot></slot>
      </div>
    </main>
  </div>
</template>
