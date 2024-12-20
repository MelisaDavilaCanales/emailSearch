<template>
  <main class="table-container rounded-md mt-2">
    <!-- <div class="loading-container text-center">
      <p class="text-lg font-semibold mb-4 text-sky-200">Cargando datos...</p>
    </div> -->

    <div class="table-wrapper h-full flex items-center">
      <table ref="dataTable"
        class="min-w-full h-[80%]  table-auto border-collapse bg-white rounded-lg overflow-hidden text-sm">
        <thead class="bg-gray-100 text-gray-700">
          <tr>
            <th class="px-4 py-2 text-left">#</th>
            <th class="px-4 py-2 text-left cursor-pointer">Id ↕</th>
            <th class="px-4 py-2 text-left cursor-pointer">Field ↕</th>
            <th class="px-4 py-2 text-left cursor-pointer">Field ↕</th>
          </tr>
        </thead>
        <tbody class="text-gray-600">
          <tr class="border-t hover:bg-gray-50">
            <td class="px-4 py-2">{{ 1 }}</td>
            <td class="px-4 py-2">{{ "Example" }}</td>
            <td class="px-4 py-2">{{ "Example" }}</td>
            <td class="px-4 py-2">{{ "Example" }}</td>
          </tr>
          <tr class="border-t hover:bg-gray-50">
            <td class="px-4 py-2">{{ 1 }}</td>
            <td class="px-4 py-2">{{ "Example" }}</td>
            <td class="px-4 py-2">{{ "Example" }}</td>
            <td class="px-4 py-2">{{ "Example" }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </main>
</template>
