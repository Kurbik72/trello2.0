<script setup lang="ts"></script>

<template>
  <div class="auth-layout">
    <slot />
  </div>
</template>

<style scoped>
.auth-layout {
  min-height: 100vh;
  width: 100%;
  background: radial-gradient(
    1200px circle at 50% 40%,
    #ffffff 0%,
    #f3f4ff 45%,
    #ecebff 70%,
    #e6e4ff 100%
  );
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
