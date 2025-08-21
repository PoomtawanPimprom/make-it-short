<template>
  <nav class="bg-white shadow sticky top-0 z-50">
    <div class="container mx-auto px-6 py-4 flex justify-between items-center">
      <!-- Brand -->
      <RouterLink  to="/" class="text-2xl font-bold text-blue-600">make it short</RouterLink>

      <!-- Links -->
      <div class="space-x-6 hidden md:flex">
        <RouterLink to="/" class="text-gray-700 hover:text-blue-600">Home</RouterLink>
        <RouterLink to="/shortURL" class="text-gray-700 hover:text-blue-600">Shorten</RouterLink>
      </div>

      <!-- Mobile Menu -->
      <button
        @click="toggleMenu"
        class="md:hidden text-gray-700 hover:text-blue-600 focus:outline-none"
      >
        ☰
      </button>
    </div>

    <!-- Mobile Dropdown -->
    <div v-if="isOpen" class="md:hidden bg-gray-50 px-6 py-4 space-y-2">
      <RouterLink to="/" class="block text-gray-700 hover:text-blue-600">Home</RouterLink>
      <RouterLink to="/shortURL" class="block text-gray-700 hover:text-blue-600">Shorten</RouterLink>
    </div>
  </nav>
</template>

<script setup>
import { ref } from 'vue'
import {RouterLink} from 'vue-router'
const isOpen = ref(false)
const toggleMenu = () => {
  isOpen.value = !isOpen.value
}
</script>

