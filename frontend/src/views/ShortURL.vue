<template>
  <div class="flex flex-col items-center justify-center h-dvh bg-gray-100">
    <div class="max-w-md mx-auto  mt-20 p-6 bg-white shadow rounded">
      <h1 class="text-2xl font-bold mb-4">Shorten Your URL</h1>
      <form @submit.prevent="shortenURL">
        <input
          type="url"
          v-model="longUrl"
          placeholder="Enter long URL"
          class="w-full p-2 border rounded mb-4"
          required
        />
        <button type="submit" class="w-full bg-blue-500 text-white py-2 rounded hover:bg-blue-600">
          Shorten
        </button>
      </form>

      <div v-if="shortUrl" class="mt-4 p-2 bg-green-100 rounded">
        <p class="font-medium">Short URL:</p>
        <a :href="shortUrl" target="_blank" class="text-blue-600 underline">{{ shortUrl }}</a>
      </div>

      <div v-if="error" class="mt-4 p-2 bg-red-100 rounded text-red-700">
        {{ error }}
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import axios from 'axios'

const longUrl = ref('')
const shortUrl = ref('')
const error = ref('')

const shortenURL = async () => {
  error.value = ''
  shortUrl.value = ''
  try {
    const res = await axios.post('/api/shorten', { url: longUrl.value })
    shortUrl.value = window.location.origin + '/' + res.data.shortID
    longUrl.value = ''
  } catch (err) {
    error.value = err.response?.data?.message || 'Something went wrong'
  }
}
</script>
