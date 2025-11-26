import { defineStore } from 'pinia'
import { ref } from 'vue'
import { apiClient, httpService } from '../service/http-service'
import { useFetch } from '@vueuse/core'

interface DashboardResponse {
  id: string
  title: string
  isFavorite: boolean
  linkToBoard: string
}

interface DashboardRequest {
  user_id: string
}
export const useBoardsStore = defineStore('boards', () => {
  const boards = ref<DashboardResponse[]>([])

  const saveBoard = async (board: DashboardResponse) => {
    const response = await httpService<DashboardResponse[]>('/api/url', {
      
    })
    boards.value = boardResponse.value ?? []
  }
  return { boards }
})
