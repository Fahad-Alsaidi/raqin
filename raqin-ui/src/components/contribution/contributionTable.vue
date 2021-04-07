<template>
  <div>
      <q-table
      title="مساهمات الرقن"
      title-class="text-primary"
      :data="data"
      :columns="columns"
      row-key="name"
    >

     <!-- table header  -->
      <template v-slot:header="props">
        <q-tr :props="props">
          <q-th
            v-for="col in props.cols"
            :key="col.name"
            :props="props"
            class="text-white bg-primary"
          >
            {{ col.label }}
          </q-th>
        </q-tr>
      </template>

      <!-- table body  -->
       <template v-slot:body="props">
        <q-tr :props="props">
          <q-td key="name" :props="props">
            {{ props.row.name }}
          </q-td>
          <q-td key="book" :props="props">
              {{ props.row.book }}
          </q-td>
          <q-td key="completed_lines" :props="props">
              {{ props.row.completed_lines }}
          </q-td>
          <q-td key="completed_pages" :props="props">
              {{ props.row.completed_pages }}
          </q-td>
          <q-td key="completion" :props="props">
              {{ props.row.completion }}
          </q-td>
          <q-td key="action">
              <q-btn>
                  سطر
              </q-btn>
              <q-btn>
                  صفحة
              </q-btn>
          </q-td>
        </q-tr>
      </template>
    </q-table>
  </div>
</template>

<script>
export default {
    name: "contribtion-table",
    data () {
        return {
          columns: [
            {
              name: 'name',
              required: true,
              label: 'المشروع',
              align: 'right',
              field: row => row.name,
              format: val => `${val}`,
            },
            { name: 'book', label: 'الكتاب', field: 'book' },
            { name: 'completed_lines', label: 'السطور المكتملة', field: 'completed_lines' },
            { name: 'completed_pages', label: 'الصفحات المكتملة', field: 'completed_pages' },
            { name: 'completion', label: 'النسبة المنجزة', field: 'completion' },
            { name: 'action', label: 'ارقن', field: 'action' },
          ],
           data: [
        {
          name: 'رقن الكتاب الفلاني',
          book: 'ارقن معنا',
          completed_lines: '20 من 12351',
          completed_pages: '13 من 276',
          completion: '9%',
        },
           ]
        }
    }
}
</script>

<style>

</style>