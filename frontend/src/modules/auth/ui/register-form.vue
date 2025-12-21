<script setup lang="ts">
import AuthorizationForm from './authorization-form.vue'
import Icon from '@/shared/ui-kit/icon/icon.vue'
</script>

<template>
  <authorization-form max-width="400">
    <template #header>
      <div class="header-icon">
        <icon is-prime-vue-icon name="user" :width="48" :height="48" :style="{ color: '#ffff' }" />
      </div>
      <span class="header-title">Create an account</span>
    </template>
  </authorization-form>
</template>

<style scoped lang="scss">
.header {
  &-icon {
    border: 1px solid transparent;
    border-radius: 8px;
    background:
      linear-gradient(90deg, #635bff, #928aff) border-box,
      #fff padding-box;
    padding: 12px;
  }
  &-title {
    font-family: 'Inter';
    font-size: 18px;
    font-style: normal;
    font-weight: 400;
    color: #1f2937;
  }
}
</style>
