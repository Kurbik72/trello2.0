<script setup lang="ts">
import Logo from '@/shared/ui-kit/logo/logo.vue'
import Search from './search.vue'
</script>

<template>
  <header class="header">
    <div>
      <logo />
    </div>
    <div>
      <search placeholder="Search" />
    </div>
  </header>
</template>

<style scoped>
.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px 24px;
  background: #fff;
}
</style>
