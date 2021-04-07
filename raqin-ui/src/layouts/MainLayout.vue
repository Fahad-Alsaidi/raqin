<template>
  <q-layout view="lHh Lpr lFf" style="direction: rtl">
    <q-header elevated>
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
           @click="drawer = !drawer" 
        />
        <q-separator dark vertical inset style="margin-right: 1rem" />
        <q-toolbar-title>
          راقن
        </q-toolbar-title>

      <q-separator inset dark vertical />
      <q-btn stretch flat label="الخروج" />
      </q-toolbar>
    </q-header>

    <q-drawer
    side="right"
    v-model="drawer"
    show-if-above
    :mini="!drawer || miniState"
    @click.capture="drawerClick"
    :width="200"
    :breakpoint="500"
    bordered persistent
    content-class="bg-grey-3"
  >
    <q-scroll-area class="fit">
      <q-list padding>
        
        <div v-for="(m, index) in a_menu" :key="index">
          <q-item
          :active="$route.path == m.to"
          clickable
          v-ripple
          :to="m.to"
        >
          <q-item-section avatar>
            <q-icon :name="m.icon" />
          </q-item-section>

          <q-item-section>
            {{ m.name }}
          </q-item-section>
        </q-item>
        </div>

        <q-separator inset />
      </q-list>
    </q-scroll-area>

    <div class="q-mini-drawer-hide absolute" style="top: 15px; right:185px">
      <q-btn
        dense
        round
        unelevated
        color="primary"
        icon="chevron_right"
        @click="miniState = true"
      />
    </div>
  </q-drawer>

    <q-page-container>
      <q-page padding>
      <router-view />
          <!-- place QPageScroller at end of page -->
          <q-page-scroller position="bottom-right" :scroll-offset="150" :offset="[18, 18]">
            <q-btn fab icon="keyboard_arrow_up" color="accent" />
          </q-page-scroller>
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script>

export default {
  name: 'MainLayout',
  data () {
    return {
      drawer: false,
      miniState: true,
      c_menu: [
        { name: "الرئيسية", icon: "home", to: "/" },
        {
          name: "مشاركاتي",
          icon: "volunteer_activism",
          to: "my-contributions"
        },
        { name: "أوسمتي", icon: "military_tech", to: "medals" },
        { name: "التفضيلات", icon: "settings_suggest", to: "preferences" },
        { name: "الملف الشخصي", icon: "account_circle", to: "profile" }
      ],
      a_menu: [
        { name: "الرئيسية", icon: "home", to: "home" },
        { name: "المشاريع", icon: "volunteer_activism", to: "project" },
        { name: "المساهمات", icon: "military_tech", to: "contribution" },
        { name: "التفضيلات", icon: "settings_suggest", to: "preferences" },
        { name: "الملف الشخصي", icon: "account_circle", to: "profile" }
      ]
    }
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
    }
  }
}
</script>
