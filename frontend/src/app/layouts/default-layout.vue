<script setup lang="ts">
import Header from '@/widgets/header/ui/header.vue'
</script>

<template>
  <div class="default-layout">
    <Header />
    <slot class="default-layout-content" />
  </div>
</template>
