<template>
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
        color="accent"
        icon="chevron_right"
        @click="miniState = true"
      />
    </div>
  </q-drawer>
</template>

<script>
export default {
  name: "Drawer",
  props: ["drawer"],
  data() {
    return {
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
        { name: "بدء مشروع", icon: "volunteer_activism", to: "project" },
        { name: "المساهمة", icon: "military_tech", to: "contribute" },
        { name: "التفضيلات", icon: "settings_suggest", to: "preferences" },
        { name: "الملف الشخصي", icon: "account_circle", to: "profile" }
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
    }
  }
};
</script>

<style></style>
