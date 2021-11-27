<template>
  <VuePdfApp theme="light" class="tw-pt-4" :key="currentPage" :page-number="currentPage" :config="config" :pdf="currentBook" />
</template>

<script>
import VuePdfApp from "vue-pdf-app";
import "vue-pdf-app/dist/icons/main.css";
// alternatives to above lib:
// pspdfkit, vue-pdfs, vue-pdf, pdfvuer
// vue-pdfjs, vue-pdf-renderer

export default {
  name: "pdf-viewer",
  props: ["pdfFile", "pageNumber"],
  components: {
    VuePdfApp
  },
  data() {
    return {
      config: {
        sidebar: false,
        secondaryToolbar: false,
        toolbar: {
          toolbarViewerLeft: {
            findbar: true,
            previous: false,
            next: false,
            pageNumber: true
          },
          toolbarViewerRight: {
            presentationMode: true,
            openFile: true,
            print: false,
            download: true,
            viewBookmark: false
          },
          toolbarViewerMiddle: {
            zoomOut: true,
            zoomIn: true,
            scaleSelectContainer: false
          }
        }
      }
    };
  },
  mounted() {},
  computed: {
    currentBook: function() {
      return this.pdfFile
        ? `https://${process.env.API_URL}${this.pdfFile.file.url}`
        : null;
    },
    currentPage: function() {
      return this.pageNumber || 1;
    }
  }
};
</script>

<style>
.pdf-app {
  background-color: rgba(255, 255, 255, 0) !important;
}

.pdfViewer {
  background-color: rgba(108,117,125,0.3) !important;
}
</style>
