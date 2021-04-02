<template>
  <div>
      <q-table
      title="المشاريع والكتب"
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
          <q-td key="admins" :props="props">
              {{ props.row.admins }}
          </q-td>
          <q-td key="notes" :props="props">
              {{ props.row.notes }}
          </q-td>
          <q-td key="book" :props="props">
              {{ props.row.book }}
          </q-td>
          <q-td key="category" :props="props">
              {{ props.row.category }}
          </q-td>
          <q-td key="author" :props="props">
              {{ props.row.author }}
          </q-td>
          <q-td key="completion" :props="props">
              {{ props.row.completion }}
          </q-td>
        </q-tr>
      </template>
    </q-table>
  </div>
</template>

<script>
export default {
    name: "projects-table",
    data () {
        return {
          columns: [
            {
              name: 'name',
              required: true,
              label: 'اسم المشروع',
              align: 'right',
              field: row => row.name,
              format: val => `${val}`,
            },
            { name: 'admins', align: 'center', label: 'إدارة المشروع', field: 'admins' },
            { name: 'notes', label: 'ملاحظات', field: 'notes' },
            { name: 'book', label: 'اسم الكتاب', field: 'book' },
            { name: 'category', label: 'تصنيف الكتاب', field: 'category' },
            { name: 'author', label: 'تأليف', field: 'author' },
            { name: 'completion', label: 'نسبة اﻹنجاز', field: 'completion' },
          ],
           data: [
        {
          name: 'رقن الكتاب الفلاني',
          admins: 'فلان علان',
          notes: 'لا يوجد ملاحظات',
          book: 'ارقن معنا',
          category: 'ادب',
          author: 'ابو ليلى',
          completion: '20%',
        },
           ]
        }
    }
}
</script>

<style>

</style>