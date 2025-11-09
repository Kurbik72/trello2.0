<script setup lang="ts">
import Icon from '@/shared/ui-kit/icon/icon.vue'
import { iconIds } from '@/shared/constants/IconIds'
import { computed } from 'vue'

interface LogoProps {
  withTitle?: boolean
}

const props = withDefaults(defineProps<LogoProps>(), {
  withTitle: true,
})
const isTitleVisible = computed(() => props.withTitle)
</script>

<template>
  <div class="logo">
    <Icon :width="24" :height="24" :iconId="iconIds.logo" />
    <h1 class="logo-text" v-if="isTitleVisible">Trello Clone</h1>
  </div>
</template>

<style scoped>
.logo {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 16px;
}
.logo-text {
  color: #172b4d;
  font-family: 'Inter';
  font-size: 18px;
  font-style: normal;
  font-weight: 700;
  line-height: 28px;
  letter-spacing: -0.27px;
}
@media (max-width: 500px) {
  .logo-text {
    display: none;
  }
}
</style>
