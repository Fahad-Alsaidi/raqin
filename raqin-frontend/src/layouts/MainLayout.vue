<template>
  <q-layout :view="isArabic() ? 'lHh Lpr lFf' : 'rHh Lpr lFf'" :style="isArabic() ? 'direction: rtl' : 'direction: ltr'">
    <q-header elevated>
      <q-toolbar>
        <q-btn flat dense round icon="menu" aria-label="Menu" @click="drawer = !drawer" />
        <q-separator dark vertical inset style="margin-right: 0.7rem; margin-left: 0.5rem" />
        <q-toolbar-title>{{ $t("layout.appName") }}</q-toolbar-title>

        <q-select
          class="to-customize-button tw-border-b-2 tw-border-b-white-500 q-mx-sm"
          v-model="lang"
          :options="langOptions"
          rounded
          dense
          borderless
          emit-value
          map-options
          options-dense
          style="min-width: 150px;"
        />
        <q-separator inset dark vertical />
        <q-btn stretch flat :label="$t('layout.logout')" @click="logout" />
      </q-toolbar>
    </q-header>

    <q-drawer
      :side="isArabic() ? 'right' : 'left'"
      v-model="drawer"
      show-if-above
      :mini="!drawer || miniState"
      @click.capture="drawerClick"
      :width="200"
      :breakpoint="500"
      bordered
      persistent
      content-class="bg-navbar fit text-white"
    >
      <q-scroll-area class="fit">
        <q-list padding>
          <div v-for="(m, index) in menu" :key="index">
            <q-item
              :active="$route.path == m.to"
              clickable
              v-ripple
              @click.native="changeRoute(m.to)"
            >
              <q-item-section avatar>
                <q-icon :name="m.icon" />
              </q-item-section>

              <q-item-section>{{ m.name }}</q-item-section>
            </q-item>
          </div>

          <q-separator inset />
        </q-list>
      </q-scroll-area>

      <div
        class="q-mini-drawer-hide absolute"
        :style="isArabic() ? 'top: 15px; right: 185px': 'top: 15px; left: 185px'"
      >
        <q-btn
          dense
          round
          unelevated
          color="primary"
          :icon="isArabic() ? 'chevron_right' : 'chevron_left'"
          @click="miniState = true"
        />
      </div>
    </q-drawer>

    <q-page-container class="bg-content">
      <q-page>
        <q-scroll-area :thumb-style="thumbStyle" :bar-style="barStyle" style="height: 100vh">
          <router-view />
          <!-- place QPageScroller at end of page -->
          <q-page-scroller position="bottom-right" :scroll-offset="150" :offset="[18, 18]">
            <q-btn fab icon="keyboard_arrow_up" color="accent" />
          </q-page-scroller>
        </q-scroll-area>
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script>
export default {
  name: "MainLayout",
  data() {
    return {
      drawer: false,
      miniState: true,
      menu: [
        { name: this.$t("layout.menu.home"), icon: "home", to: "/home" },
        { name: this.$t("layout.menu.books"), icon: "menu_book", to: "/books" },
        {
          name: this.$t("layout.menu.myPages"),
          icon: "description",
          to: "/pages"
        },
        // { name: "أوسمتي", icon: "military_tech", to: "medals" },
        {
          name: this.$t("layout.menu.profile"),
          icon: "account_circle",
          to: "/profile"
        }
      ],
      thumbStyle: {
        right: "4px",
        borderRadius: "5px",
        backgroundColor: "#027be3",
        width: "5px",
        opacity: 0.75
      },

      barStyle: {
        right: "2px",
        borderRadius: "9px",
        backgroundColor: "#027be3",
        width: "9px",
        opacity: 0.2
      },
      lang: this.$i18n.locale,
      langOptions: [
        { value: "ar", label: "عربي" },
        { value: "en-us", label: "English" }
      ]
    };
  },
  methods: {
    drawerClick(e) {
      // if in "mini" state and user
      // click on drawer, we switch it to "normal" mode
      if (this.miniState) {
        this.miniState = false;

        // notice we have registered an event with capture flag;
        // we need to stop further propagation as this click is
        // intended for switching drawer to "normal" mode only
        e.stopPropagation();
      }
    },
    isArabic() {
      return this.$i18n.locale == "ar";
    },
    logout() {
      this.$store
        .dispatch("auth/logout")
        .then(this.$router.push("/login"))
        .catch(err => console.log(err));
    },
    changeRoute(to) {
      if (to != this.$route.path) {
        this.$router.push(to);
      }
    }
  },
  watch: {
    lang(lang) {
      this.$i18n.locale = lang;
      this.$store.commit("setting/setLocale", lang);
      location.reload();
    }
  }
};
</script>
<style>
body {
  overflow: hidden;
}
.to-customize-button .q-field__native,
.q-field__marginal {
  color: rgb(247, 245, 245);
}
</style>
