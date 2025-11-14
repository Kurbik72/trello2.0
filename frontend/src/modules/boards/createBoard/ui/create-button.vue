<script setup lang="ts">
import Button from '@/shared/ui-kit/button/button.vue'

const emit = defineEmits(['toggle'])
</script>

<template>
  <div>
    <Button icon="pi pi-plus" label="Create Board" theme="purple" @click="emit('toggle')" />
  </div>
</template>

<style scoped></style>
