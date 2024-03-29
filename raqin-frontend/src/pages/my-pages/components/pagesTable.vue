<template>
  <div>
    <q-table
      :title="$t('myPagesPage.table.title')"
      title-class="text-primary"
      class="bg-content"
      :data="data"
      :columns="columns"
      row-key="name"
      :rows-per-page-label="$t('myPagesPage.table.rowsPerPage')"
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
          <q-td key="name" :props="props">{{ props.row.book ? props.row.book.name : "-" }}</q-td>
          <q-td
            key="text"
            :props="props"
            style="max-width: 20ch; overflow: hidden;"
          >{{ props.row.text }}</q-td>
          <q-td key="author" :props="props">{{ props.row.user ? props.row.user.username : "-" }}</q-td>
          <q-td key="pageNumber" :props="props">{{ props.row.pageNumber }}</q-td>
          <q-td key="stage" :props="props">{{ Stage(props.row.stage) }}</q-td>
          <q-td key="action" :props="props" v-if="belogsToUser(props.row.user)">
            <q-btn
              stretch
              color="primary"
              :label="$t('myPagesPage.table.contributeBtn')"
              @click="storeAndOpenCurrentBookAndPage(props.row)"
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
  name: "pages-table",
  mixins: [tableMixin],
  data() {
    return {
      columns: [
        {
          name: "name",
          required: true,
          label: this.$t("myPagesPage.table.columns.first"),
          align: "right",
          field: row => row.name,
          format: val => `${val}`
        },
        {
          name: "text",
          label: this.$t("myPagesPage.table.columns.second"),
          field: "text"
        },
        {
          name: "author",
          label: this.$t("myPagesPage.table.columns.third"),
          field: "author"
        },
        {
          name: "pageNumber",
          label: this.$t("myPagesPage.table.columns.fourth"),
          field: "pageNumber"
        },
        {
          name: "stage",
          label: this.$t("myPagesPage.table.columns.fifth"),
          field: "stage"
        },
        {
          name: "action",
          label: this.$t("myPagesPage.table.columns.sixth"),
          field: "action"
        }
      ],
      data: []
    };
  },
  mounted() {
    this.getPages();
  },
  methods: {
    getPages() {
      this.$store
        .dispatch("editor/pagesByUserID")
        .then(res => {
          this.data = res;
        })
        .catch(err => {
          console.log(err);
        });
    },
    storeAndOpenCurrentBookAndPage(page) {
      this.$store.commit("editor/setCurrentBook", page.book);
      this.$store.commit("editor/setCurrentPage", page);
      this.$router.push("/contribute");
    },
    belogsToUser(user) {
      if (user && this.$store.getters["auth/getUser"]) {
        return user.id == this.$store.getters["auth/getUser"].id;
      } else return false;
    }
  }
};
</script>

<style></style>
