<script setup lang="ts">
import { computed } from 'vue'
import { useToggle } from '@/shared/composables/use-toggle'
import IconButton from '@/shared/ui-kit/option-button/icon-button.vue'
const { isActive, toggle } = useToggle()

const filterCardClasses = computed(
  () => ` ${isActive.value ? 'pi pi-filter-fill icon-active' : 'pi pi-filter'}`,
)
</script>

<template>
  <icon-button :icon="filterCardClasses" :isActive="isActive" @toggle="toggle" />
</template>
