<script setup lang="ts">
import Logo from '@/shared/ui-kit/logo/logo.vue'
import Search from './search.vue'
import FiltersCard from '@/widgets/filters-card/ui/filters-card.vue'
</script>

<template>
  <header class="header">
    <logo />
    <div class="header-options">
      <search />
      <filters-card />
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
  border-bottom: 1px solid #e5e7eb;
}
.header-options {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 16px;
}
</style>
