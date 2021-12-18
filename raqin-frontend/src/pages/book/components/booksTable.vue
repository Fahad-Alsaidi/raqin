<template>
  <div>
    <q-table
      :title="$t('booksPage.table.title')"
      class="bg-content"
      title-class="text-primary"
      :data="books"
      :columns="columns"
      row-key="name"
      :rows-per-page-label="$t('booksPage.table.rowsPerPage')"
      :pagination-label="
        (firstRowIndex, endRowIndex, totalRowsNumber) =>
          `${firstRowIndex} - ${endRowIndex} ${$t('helperWords.of')} ${totalRowsNumber}`
      "
    >
      <!-- table header  -->
      <template v-slot:header="props">
        <q-tr :props="props">
          <q-th
            v-for="col in props.cols"
            :key="col.name"
            :props="props"
            class="text-white bg-primary"
          >{{ col.label }}</q-th>
        </q-tr>
      </template>

      <!-- table body  -->
      <template v-slot:body="props">
        <q-tr :props="props">
          <q-td key="name" :props="props">{{ props.row.name }}</q-td>
          <q-td key="categories" :props="props">
            <q-chip
              v-for="(category, index) in props.row.categories"
              :key="index"
              color="primary"
            >{{ category.name }}</q-chip>
          </q-td>
          <q-td key="authors" :props="props">
            <q-chip v-for="(author, index) in props.row.authors" color="primary" :key="index">
              {{
              author.name
              }}
            </q-chip>
          </q-td>
          <q-td key="numberOfPages" :props="props">{{ props.row.numberOfPages }}</q-td>
          <q-td
            key="completion"
            :props="props"
          >{{ bookCompletion(props.row.pages, props.row.numberOfPages) }}%</q-td>
          <q-td key="stage" :props="props">{{ Stage(props.row.stage) }}</q-td>
          <q-td key="action" :props="props">
            <q-btn
              stretch
              color="primary"
              :label="$t('booksPage.table.contributeBtn')"
              @click="storeAndOpenCurrentBook(props.row)"
            />
          </q-td>
        </q-tr>
      </template>
    </q-table>
  </div>
</template>

<script>
import tableMixin from "pages/mixins/tableMixins";

export default {
  name: "books-table",
  mixins: [tableMixin],
  data() {
    return {
      books: [],
      columns: [
        {
          name: "name",
          required: true,
          label: this.$t("booksPage.table.columns.first"),
          align: "right",
          field: row => row.name,
          format: val => `${val}`
        },
        {
          name: "categories",
          label: this.$t("booksPage.table.columns.second"),
          field: "categories"
        },
        {
          name: "authors",
          label: this.$t("booksPage.table.columns.third"),
          field: "authors"
        },
        {
          name: "numberOfPages",
          label: this.$t("booksPage.table.columns.fourth"),
          field: "numberOfPages"
        },
        {
          name: "completion",
          label: this.$t("booksPage.table.columns.fifth"),
          field: "completion"
        },
        {
          name: "stage",
          label: this.$t("booksPage.table.columns.sixth"),
          field: "stage"
        },
        {
          name: "action",
          label: this.$t("booksPage.table.columns.seventh"),
          field: "action"
        }
      ]
    };
  },
  mounted() {
    this.getBooks();
  },
  methods: {
    bookCompletion(pages, total) {
      let donePages = pages.filter(page => page.stage === "done");
      let distinctPages = [...new Set(donePages.map(x => x.pageNumber))];

      return Math.floor((distinctPages.length / total) * 100);
    },
    storeAndOpenCurrentBook(book) {
      this.$store.commit("editor/setCurrentBook", book);
      this.$store.commit("editor/setCurrentPage", null);
      this.$router.push("/contribute");
    }
  }
};
</script>

<style>
</style>