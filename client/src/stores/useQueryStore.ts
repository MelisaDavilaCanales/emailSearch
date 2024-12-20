import { computed, ref } from 'vue'
import { defineStore } from 'pinia'



export const useQueryStore = defineStore('query', () => {
    
    const page = ref<string>('')
    const pageSize = ref<string>('')
    const term = ref<string>('')
    const id = ref<string | null >('')
    const typeSearch = ref<string>('')
    const field = ref<string>('')
    const sort= ref<string>('')

    const url = computed(() => {
        return baseUrl.value + query.value
    })

    const baseUrl = computed(() => {
        return 'http://localhost:8000/api/' + typeSearch.value + "/" + id.value
    })

    const query = computed(() => {
        return  "?" + page.value + "&" + pageSize.value + "&" + term.value + "&" + id.value + "&" + field.value + "&" + sort.value    
    })
    
    function setPage(value: string) {
        page.value = "page=" + value
    }

    function setPageSize(value: string) {
        pageSize.value = "pageSize=" + value
    }

    function setTerm(value: string) {
        term.value = "term=" + value
    }

    function setId(value: string) {
        id.value = value
    }

    function setTypeSearch(value: string) {
        typeSearch.value = value
    }

    function setField(value: string) {
        field.value =  "field=" + value
    }

    function setSort(value: string) {
        sort.value =  "sort=" + value
    }

    return { setPage, setPageSize, setTerm, setId, setTypeSearch, setField, setSort, url, query }
})
