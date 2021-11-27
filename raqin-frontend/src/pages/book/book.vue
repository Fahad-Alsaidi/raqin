<template>

  <div class="q-ma-xl">
    <BooksTable />
  </div>
  
</template>

<script>

import BooksTable from "./components/booksTable"

export default {
  name: "project",
  components: {
    BooksTable
  },
  data(){
    return {
    }
  }
};
</script>
