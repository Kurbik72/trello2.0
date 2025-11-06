import { ref } from 'vue'

export const useToggle = () => {
  const isActive = ref(false)

  const toggle = () => {
    isActive.value = !isActive.value
  }
  const setActive = () => {
    isActive.value = true
  }
  const setInactive = () => {
    isActive.value = false
  }
  return {
    isActive,
    toggle,
    setActive,
    setInactive,
  }
}
