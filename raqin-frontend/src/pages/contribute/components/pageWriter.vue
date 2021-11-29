<template>
  <div v-if="book">
    <div
      class="
        q-my-md
        tw-w-full
        tw-flex
        tw-flex-row
        tw-flex-wrap
        tw-justify-between
        tw-justify-items-start
        tw-content-start
      "
    >
      <q-input
        filled
        v-model.number="pageNumber"
        type="number"
        label="رقم الصفحة"
        class="tw-w-full xs:tw-w-1/3 tw-pb-2"
        :min="1"
        :max="MaxPagesNo"
        error-message="رقم الصفحة خطأ"
        :error="!isValidPageNumber()"
        dense
      />

      <q-select
        v-if="pages.length > 0"
        filled
        v-model="page"
        input-debounce="0"
        label="صاحب الصفحة"
        :options="pages"
        :option-label="(page) => pageByWhichUser(page)"
        class="tw-w-full xs:tw-w-1/3 tw-pb-2"
        :style="colorizeSelectBox()"
        behavior="menu"
        dense
      />
    </div>

    <div
      v-if="pages.length > 0"
      class="
        tw-my-2
        tw-w-full
        tw-flex
        tw-flex-row
        tw-flex-wrap
        tw-justify-center
        tw-justify-items-start
        tw-content-start
      "
    >
      <div
        class="
          text-orange
          tw-w-full
          tw-flex
          tw-flex-row
          tw-flex-wrap
          tw-justify-center
          tw-items-center
          tw-content-ceter
        "
        v-if="pages.length > 0 && !userHasPage()"
      >
        <div>هناك من يعمل على هذه الصفحة</div>

        <q-icon name="warning" class="tw-mx-2" />
        <q-tooltip>
          يوجد نسخة/نسخ من هذه الصفحة لـ أشخاص يعملون عليها, يفضل ترك الصفحات
          التي يتم العمل عليها من قبل المتطوعين اﻵخرين إن لم تكن مهملة من قبلهم
        </q-tooltip>
      </div>

      <div
        class="
          text-green
          tw-w-full
          tw-flex
          tw-flex-row
          tw-flex-wrap
          tw-justify-center
          tw-items-center
          tw-content-start
        "
        v-if="pages.length > 0 && userHasPage()"
      >
        <div>أكمل العمل على صفحتك</div>

        <q-icon name="check_circle" class="tw-mx-2" />
        <q-tooltip>
          يوجد لديك نسخة لهذه الصفحة, يمكنك إكمال العمل عليها
        </q-tooltip>
      </div>
    </div>

    <q-card flat bordered class="text-white">
      <q-card-actions flat class="bg-primary">
        <q-btn label="حفظ" @click="updateOrCreatePage()" />
        <q-space />
        <a> {{` كتاب: ${book.name} `}} </a>
      </q-card-actions>

      <q-card-section class="tw-w-full no-padding">
        <q-input
          v-model="pageText"
          class="tw-w-full bg-navbar tw-px-2"
          rounded
          type="textarea"
        />
      </q-card-section>
    </q-card>
  </div>
</template>

<script>
import notifyMixin from "../../mixins/notifyMixin";
import pageAPI from "../../my-pages/pageAPI";
export default {
  name: "page-writer",
  mixins: [notifyMixin],
  data() {
    return {
      book: null,
      page: null,
      pages: [],
      currentUserPageID: 0,
      pageNumber: 1,
      pageText: "",
      MaxPagesNo: 0,
    };
  },
  mounted() {
    this.getCurrentBookAndPageIfExist();
    this.checkPagesOfCurrentPageNumber(this.pageNumber);
  },
  methods: {
    colorizeSelectBox() {
      if (this.userHasPage()) {
        return "border-bottom: 2px solid #4caf50";
      } else if (!this.userHasPage()) {
        return "border-bottom: 2px solid #ff9800";
      } else "";
    },
    updateOrCreatePage() {
      if (
        this.page &&
        this.page.id == this.currentUserPageID &&
        this.page.pageNumber == this.pageNumber
      ) {
        this.updatePage();
      } else this.createPage();
    },
    updatePage() {
      let page = {
        text: this.pageText,
      };
      this.$store
        .dispatch("editor/updatePage", {
          page: page,
          pageID: this.currentUserPageID,
        })
        .then((res) => {
          console.log(res);
          this.$q.notify(this.pageTextSaved);
        })
        .catch((err) => {
          console.log(err);
          this.$q.notify(this.networkError);
        });
    },
    createPage() {
      let page = {
        text: this.pageText,
        pageNumber: this.pageNumber,
        book: this.book.id,
      };

      this.$store
        .dispatch("editor/createPage", page)
        .then((res) => {
          console.log(res);
          this.$q.notify(this.pageTextSaved);
        })
        .catch((err) => {
          console.log(err);
          this.$q.notify(this.networkError);
        });
    },
    getCurrentBookAndPageIfExist() {
      let book = this.$store.getters["editor/getCurrentBook"];
      if (book != null && this.book == null) {
        this.book = book;
        this.MaxPagesNo = book.numberOfPages
      }

      let page = this.$store.getters["editor/getCurrentPage"];
      if (page != null) {
        this.page = page;
        this.pageNumber = page.pageNumber;
        this.currentUserPageID = page.id;
        this.pageText = page.text;
      }
    },
    userHasPage() {
      if (this.getUserOwnPageIfExist()) {
        return true;
      } else return false;
    },
    getUserOwnPageIfExist() {
      let currentUser = this.$store.getters["auth/getUser"];
      let userOwnPage = this.pages.find(
        (page) => page.user.id == currentUser.id
      );
      if (userOwnPage) {
        return userOwnPage;
      } else return null;
    },
    changeToUserOwnPage() {
      let userOwnPage = this.getUserOwnPageIfExist();

      if (userOwnPage) {
        this.page = userOwnPage;
        this.currentUserPageID = userOwnPage.id;
        this.pageText = userOwnPage.text;
      } else {
        this.page = null;
        this.currentUserPageID = 0;
        this.pageText = "";
      }
    },
    checkPagesOfCurrentPageNumber(val) {
      pageAPI
        .getByPageNumberAndBookID(val, this.book.id)
        .then((res) => {
          this.pages = res.data;
          this.changeToUserOwnPage();
        })
        .catch((err) => {
          console.log(err);
        });
    },
    pageByWhichUser(page) {
      if (page) {
        if (
          page.user &&
          page.user.id == this.$store.getters["auth/getUser"].id
        ) {
          return "أنا";
        } else {
          return page.user.username;
        }
      } else {
        return "-";
      }
    },
    isValidPageNumber(){
      if (this.pageNumber > this.MaxPagesNo || this.pageNumber < 1) {
        return false
      } else return true
    }
  },
  watch: {
    book: function (val) {
      this.$emit("book", val);
    },
    pageNumber: function (newVal, oldVal) {
      if (newVal !== oldVal && newVal > 0 && newVal < this.MaxPagesNo + 1) {
        this.checkPagesOfCurrentPageNumber(newVal);
        this.$emit("pageNumber", newVal);
      }
    },
    page: function (newPage) {
      let userPage = this.getUserOwnPageIfExist();
      if (newPage && userPage && newPage.id == userPage.id) {
        this.currentUserPageID = newPage.id;
        this.pageText = newPage.text;
      } else if (newPage) {
        this.currentUserPageID = 0;
        this.pageText = newPage.text;
      }
    },
  },
};
</script>

<style></style>
