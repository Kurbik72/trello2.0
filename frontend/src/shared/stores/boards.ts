import { defineStore } from 'pinia'
import { ref } from 'vue'
import { httpService } from '../service/http-service'
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

  const saveBoard = async (board: DashboardRequest) => {
    const { data: boardResponse } = await httpService.post<DashboardResponse[]>('/api/boards', {
      body: JSON.stringify(board),
    })
    boards.value = boardResponse.value ?? []
  }
  return { boards }
})
