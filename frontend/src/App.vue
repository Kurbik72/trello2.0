<script setup lang="ts">
import DefaultLayout from './app/layouts/default-layout.vue'
import { useRoute } from 'vue-router'
const route = useRoute()
</script>

<template>
  <default-layout v-if="!route.fullPath.startsWith('/auth')" />
  <router-view />
</template>
