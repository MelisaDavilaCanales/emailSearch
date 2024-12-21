import { computed, ref, watch } from 'vue'
import { defineStore } from 'pinia'

export interface Person {
  id: string
  name: string
  email: string
}

export const usePersonStore = defineStore('persons', () => {
  const persons = ref<Person[]>([])

  const pageNumber = ref<number>(1)
  const pageSize = ref<number>(40)
  const tatalPages = ref<number>(0)
  const searchTerm = ref<string>('')
  const searchField = ref<string>('_all')
  const sortField= ref<string>('name')
  const sortOrder = ref<string>('asc')

  const baseUrl = import.meta.env.VITE_API_BASE_URL

  const personSearchURL = computed(() => {
    return baseUrl + '/persons' + query.value
  })

  const query = computed(() => {
    return (
      '?' +
      'page=' +
      pageNumber.value +
      '&page_size=' +
      pageSize.value +
      '&term=' +
      searchTerm.value +
      '&field=' +
      searchField.value +
      '&sort=' +
      sortField.value +
      '&order=' +
      sortOrder.value
    )
  })

  async function fetchPersons() {
      const response = await fetch(personSearchURL.value);
      const data = await response.json();

      persons.value = []

      if (response.ok) {
        data.data.persons.forEach((person: Person) => {
          persons.value.push({
            id: person.id,
            name: person.name,
            email: person.email,
          });
        });

        pageNumber.value = data.data.page
        pageSize.value = data.data.page_size
        tatalPages.value = data.data.total_pages

        console.log("statusCode:" + response.status);
      } else {
        console.log('Error fetching emails:', response.statusText);
      }
    }

    function setPersonPageNumber(page: number) {
      pageNumber.value = page
    }

    function setPersonPageSize(size: number) {
      pageSize.value = size
    }

    function setPersonSearchTerm(term: string) {
      searchTerm.value = term
    }

    function setPersonSearchField(field: string) {
      searchField.value = field
    }

    function setPersonSortOrder(order: string) {
      sortOrder.value = order
    }

    function setPersonSortField(field: string) {
      sortField.value = field
    }

    function sortPersonsByField(field: string) {
      setPersonSortField(field)

      if (sortOrder.value == 'asc') {
        sortOrder.value = 'desc'
        return
      }

      sortOrder.value = 'asc'
    }

      watch(personSearchURL, fetchPersons)


    return {
      persons,

      pageNumber,
      pageSize,
      tatalPages,

      fetchPersons,

      setPersonPageNumber,
      setPersonPageSize,
      setPersonSearchTerm,
      setPersonSearchField,
      setPersonSortOrder,
      setPersonSortField,

      sortPersonsByField,
    }


})
