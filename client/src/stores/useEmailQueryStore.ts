import { computed, ref } from 'vue'
import { defineStore } from 'pinia'

export const useEmailQueryStore = defineStore('emailQuery', () => {
    const pageNumber = ref<number>(0)
    const pageSize = ref<number>(10)
    const searchTerm = ref<string>('')
    const searchField = ref<string>('_all')
    const sortOrder = ref<string>('asc')
    const sortField= ref<string>('name')

    const baseUrl = import.meta.env.VITE_API_BASE_URL;

    const emailSearchURL = computed(() => {
      return baseUrl + "/persons" + query.value
    })

    // http://localhost:8080/emails?page=1&page_size=10&term=charles&field=from&sort=to&order=desc
    const query = computed(() => {
        return  "?" + "page=" +  pageNumber.value + "&page_size=" + pageSize.value + "&term=" + searchTerm.value + "&field=" + searchField.value + "&sort=" + sortOrder.value + "&order=" + sortField.value
    })

    function setPageNumber(page: number) {
        pageNumber.value = page
    }

    function setPageSize(size: number) {
        pageSize.value = size
    }

    function setSearchTerm(term: string) {
        searchTerm.value = term
    }

    function setSearchField(field: string) {
        searchField.value = field
    }

    function setSortOrder(order: string) {
        sortOrder.value = order
    }

    function setSortField(field: string) {
        sortField.value = field
    }

    return { setPageNumber, setPageSize, setSearchTerm, setSearchField, setSortOrder, setSortField, emailSearchURL}

})
