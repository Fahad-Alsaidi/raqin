<template>
  <div class="tw-h-screen">
    <div
      class="tw-h-full md:tw-flex md:tw-flex-row md:tw-flex-wrap"
    >
      <PageWriter
      class="tw-w-full md:tw-w-1/2 tw-px-4 tw-py-2"
        @book="passBook"
        @pageNumber="passPageNumber"
      />

      <BookPDFViewer
      class="tw-w-full md:tw-w-1/2 tw-px-4 tw-py-2"
        :pdfFile="book"
        :pageNumber="pageNumber"
      />
    </div>
  </div>
</template>

<script>
import PageWriter from "./components/pageWriter.vue";
import BookPDFViewer from "./components/PDFViewer.vue";

export default {
  name: "contribute",
  components: { BookPDFViewer, PageWriter },
  data() {
    return {
      book: null,
      pageNumber: null
    };
  },
  methods: {
    passBook(book) {
      this.book = book;
    },
    passPageNumber(pageNumber) {
      this.pageNumber = pageNumber;
    }
  }
};
</script>

<style scoped></style>
