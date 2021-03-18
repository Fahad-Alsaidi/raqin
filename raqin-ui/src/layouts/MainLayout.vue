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

    <Drawer :drawer="drawer" />

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script>
import Drawer from '../components/drawer'
export default {
  name: 'MainLayout',
  components:{Drawer},
  data () {
    return {
      drawer: false,
    }
  },
}
</script>
