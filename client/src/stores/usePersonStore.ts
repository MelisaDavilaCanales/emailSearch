import { computed, ref, watch } from 'vue';
import { defineStore } from 'pinia';

export interface Person {
  id: string;
  name: string;
  email: string;
}

export const usePersonStore = defineStore('persons', () => {
  const personList = ref<Person[]>([]);

  const selectedPersonEmail = ref<string>('');

  const totalPage = ref<number>(0);
  const pageNumber = ref<number>(1);
  const pageSize = ref<number>(0);
  const sortField = ref<string>('name');
  const sortOrder = ref<string>( 'asc');

  const searchParam = ref<string>('')

  const baseUrl = import.meta.env.VITE_API_BASE_URL;

  const query = computed(() => {
    return (
      '?' +
      'page=' +
      pageNumber.value +
      '&page_size=' +
      pageSize.value +
      searchParam.value +
      '&sort=' +
      sortField.value +
      '&order=' +
      sortOrder.value
    );
  });

  const personSearchURL = computed(() => {
    return baseUrl + '/persons' + query.value;
  });

  async function fetchPersons() {
    console.log('fetching persons:', personSearchURL.value);
    const response = await fetch(personSearchURL.value);

    personList.value = [];

    if (response.ok) {
      const data = await response.json();

      if (data && data.data &&  data.data.persons !== null) {
        data.data.persons.forEach((person: Person) => {
          personList.value.push({
            id: person.id,
            name: person.name,
            email: person.email,
          });
        });

        totalPage.value = data.data.total_pages || 0;
        pageNumber.value = data.data.page || 1;
        pageSize.value = data.data.page_size || 0;
        console.log('statusCode:' + response.status);
        console.log('data:', data);
      }

    }
  }


  function setPersonPageNumber(page: number) {
    pageNumber.value = page;
  }

  function setPersonPageSize(size: number) {
    pageSize.value = size;
  }

  function setPersonSortOrder(order: string) {
    sortOrder.value = order;
  }

  function setPersonSortField(field: string) {
    sortField.value = field;
  }

  function setSelectedPersonEmail(email: string) {
    selectedPersonEmail.value = email;
    localStorage.setItem('selectedPersonEmail', email);
  }


  function sortPersonsByField(field: string) {
    sortField.value = field;

    if (sortOrder.value == 'asc') {
      sortOrder.value = 'desc';
      localStorage.setItem('sortOrder', 'desc');
      return;
    }

    sortOrder.value = 'asc';
    localStorage.setItem('sortOrder', 'asc');
  }

  function setPersonSearchParams(field: string, term: string) {
    pageNumber.value = 1
    searchParam.value =  '&field=' + field + '&term=' + term
  }

  function setNextPage() {
    if (pageNumber.value < totalPage.value) {
      pageNumber.value++
    }
  }

  function setPreviousPage() {
    if (pageNumber.value > 1) {
      pageNumber.value--
    }
  }

  watch(query, fetchPersons);
  watch(query, () => {
    console.log('query:', query.value);
  });

  return {
    personList,
    selectedPersonEmail,

    sortOrder,
    sortField,

    pageNumber,
    pageSize,
    totalPage,

    setNextPage,
    setPreviousPage,

    fetchPersons,

    setPersonPageNumber,
    setPersonPageSize,
    setPersonSortOrder,
    setPersonSortField,

    sortPersonsByField,
    setSelectedPersonEmail,

    searchParam,
    setPersonSearchParams,
  };
});
