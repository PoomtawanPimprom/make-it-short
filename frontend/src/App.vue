

<script setup lang="ts">
import Navbar from './components/Navbar.vue';
import { RouterView } from 'vue-router';
import Footer from './components/Footer.vue';
</script>

<template>
  <Navbar/>
  <main>
    <RouterView/>
  </main>
  <Footer/>
</template>


