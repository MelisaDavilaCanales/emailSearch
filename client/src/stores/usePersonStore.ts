import { computed, ref, watch } from 'vue';
import { defineStore } from 'pinia';

export interface Person {
  id: string;
  name: string;
  email: string;
}

export const usePersonStore = defineStore('persons', () => {
  const persons = ref<Person[]>([]);
  const selectedPersonEmail = ref<string>(localStorage.getItem('selectedPersonEmail') || '');

  const pageNumber = ref<number>(parseInt(localStorage.getItem('pageNumber') || '1', 10));
  const pageSize = ref<number>(parseInt(localStorage.getItem('pageSize') || '40', 10));
  const searchTerm = ref<string>(localStorage.getItem('searchTerm') || '');
  const searchField = ref<string>(localStorage.getItem('searchField') || '_all');
  const sortField = ref<string>(localStorage.getItem('sortField') || 'name');
  const sortOrder = ref<string>(localStorage.getItem('sortOrder') || 'asc');

  const baseUrl = import.meta.env.VITE_API_BASE_URL;

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
    );
  });

  const personSearchURL = computed(() => {
    return baseUrl + '/persons' + query.value;
  });

  async function fetchPersons() {
    const response = await fetch(personSearchURL.value);
    const data = await response.json();

    persons.value = [];

    if (response.ok) {
      data.data.persons.forEach((person: Person) => {
        persons.value.push({
          id: person.id,
          name: person.name,
          email: person.email,
        });
      });

      pageNumber.value = data.data.page;
      pageSize.value = data.data.page_size;
      localStorage.setItem('pageNumber', String(pageNumber.value));
      localStorage.setItem('pageSize', String(pageSize.value));

      console.log('statusCode:' + response.status);
    } else {
      console.log('Error fetching emails:', response.statusText);
    }
  }

  function setPersonPageNumber(page: number) {
    pageNumber.value = page;
  }

  function setPersonPageSize(size: number) {
    pageSize.value = size;
  }

  function setPersonSearchTerm(term: string) {
    searchTerm.value = term;
  }

  function setPersonSearchField(field: string) {
    searchField.value = field;
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
    setPersonSortField(field);

    if (sortOrder.value == 'asc') {
      sortOrder.value = 'desc';
      localStorage.setItem('sortOrder', 'desc');
      return;
    }

    sortOrder.value = 'asc';
    localStorage.setItem('sortOrder', 'asc');
  }

  watch([pageNumber, pageSize, searchTerm, searchField, sortField, sortOrder], () => {
    localStorage.setItem('pageNumber', String(pageNumber.value));
    localStorage.setItem('pageSize', String(pageSize.value));
    localStorage.setItem('searchTerm', searchTerm.value);
    localStorage.setItem('searchField', searchField.value);
    localStorage.setItem('sortField', sortField.value);
    localStorage.setItem('sortOrder', sortOrder.value);
  });

  return {
    persons,
    selectedPersonEmail,

    pageNumber,
    pageSize,

    fetchPersons,

    setPersonPageNumber,
    setPersonPageSize,
    setPersonSearchTerm,
    setPersonSearchField,
    setPersonSortOrder,
    setPersonSortField,

    sortPersonsByField,
    setSelectedPersonEmail,
  };
});
