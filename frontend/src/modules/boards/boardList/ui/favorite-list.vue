<script setup lang="ts"></script>

<template>
  <div class="favorite-list">
    <h1 class="favorite-list-title">
      <i class="pi pi-star-fill" :style="{ color: '#EAB308' }" />Favorite List
    </h1>
  </div>
</template>

<style scoped>
.favorite-list-title {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #374151;
  font-family: 'Inter';
  font-size: 20px;
  font-style: normal;
  font-weight: 500;
  line-height: 28px;
}
</style>
