<script setup lang="ts">
import AuthorizationForm from './authorization-form.vue'
import Icon from '@/shared/ui-kit/icon/icon.vue'
</script>

<template>
  <authorization-form max-width="400">
    <template #header>
      <div>
        <icon
          is-prime-vue-icon
          :width="48"
          :height="48"
          :style="{ color: '#635bff' }"
          name="user"
        />
        <p>sfsdf</p>
      </div>
    </template>
  </authorization-form>
</template>

<style scoped></style>
