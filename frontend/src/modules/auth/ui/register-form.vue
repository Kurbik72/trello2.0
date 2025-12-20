<script setup lang="ts">
import AuthorizationForm from './authorization-form.vue'
import Icon from '@/shared/ui-kit/icon/icon.vue'
</script>

<template>
  <authorization-form max-width="400">
    <template #header>
      <div>
        <icon prime-vue-icon="pi pi-user" :width="24" :height="24" :style="{ color: '#635bff' }" />
      </div>
    </template>
  </authorization-form>
</template>

<style scoped></style>
