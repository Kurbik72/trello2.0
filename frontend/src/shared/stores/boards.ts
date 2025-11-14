import { defineStore } from 'pinia'
import { ref } from 'vue'

interface Dashboard {
  id: string
  title: string
  isFavorite: boolean
  linkToBoard: string
  user_id: string
}

export const useBoardsStore = defineStore('boards', () => {
  const boards = ref<Dashboard[]>([])

  return { boards }
})
