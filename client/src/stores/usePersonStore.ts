import { ref } from 'vue'
import { defineStore } from 'pinia'

export interface Person {
  id: string
  name: string
  email: string
}

export const usePersonStore = defineStore('persons', () => {
  const persons = ref<Person[]>([])

  async function fetchPersons() {
      const response = await fetch('http://localhost:8080/persons?term=&field=&page=1&page_size=50&sort=name&order=asc');
      const data = await response.json();

      if (response.ok) {
        data.data.persons.forEach((person: Person) => {
          persons.value.push({
            id: person.id,
            name: person.name,
            email: person.email,
          });
        });

        console.log("statusCode:" + response.status);
      } else {
        console.log('Error fetching emails:', response.statusText);
      }
    }

  return { persons, fetchPersons }
})
