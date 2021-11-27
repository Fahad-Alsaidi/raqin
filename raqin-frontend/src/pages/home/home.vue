<template>
  <q-page class="tw-h-screen">
    <div v-if="chartExists" class="tw-h-full tw-py-8 md:tw-flex md:tw-flex-row md:tw-flex-wrap">
      <BookStats ref="bookChart" class="tw-h-full tw-w-full lg:tw-w-1/2 tw-px-4 tw-my-8" />
      <PageStats ref="pageChart" class="tw-h-full tw-w-full lg:tw-w-1/2 tw-px-4 tw-my-8" />
    </div>
    <div
      v-else
      class="tw-flex tw-flex-col tw-h-full tw-w-full md:tw-w-1/2 tw-justify-center tw-mx-auto"
    >
      <img class src="~assets/undraw_warning_cyit.svg" alt />
      <p class="tw-text-center tw-pb-8 tw-pb-16">لا يوجد بيانات</p>
    </div>
  </q-page>
</template>

<script>
import BookStats from "./components/bookStats.vue";
import PageStats from "./components/pageStats.vue";
export default {
  name: "Home",
  components: {
    BookStats,
    PageStats
  },
  data() {
    return {
      chartExists: true,
    };
  },
  mounted() {
    if (this.isChartsExist) {
      this.chartExists = true;
    } else this.chartExists = false;
  },
  computed: {
    isChartsExist: function() {
      let bc = this.$refs?.bookChart?.$refs;
      let pc = this.$refs?.pageChart?.$refs;

      if (bc || pc) {
        return Object.entries(bc).length > 0 || Object.entries(pc).length > 0
          ? true
          : false;
      } else {
        return false;
      }
    }
  },
};
</script>

<style lang="sass" scoped>
</style>
