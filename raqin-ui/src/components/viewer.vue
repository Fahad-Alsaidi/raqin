<template>
  <div id="app">
    <vue-previewer :style="style" :images="img" mode="image" :options="options">
        <template v-slot:footer="{image}">{{image.name}}</template>
      </vue-previewer>
  </div>
</template>

<script>
import VuePreviewer from 'vue-image-previewer'

export default {
  props:['img'],
  components: {
    VuePreviewer
  },
  data() {
      return {
        options: {
          defaultWidth: '98%', 
          defaultHeight: '100px',
          // more previewer options here
        },
        style: {
          width: '100%',
          // more styles here
        },
        imgs: []
      }
    },
  methods: {
    inited(viewer) {
      this.$viewer = viewer;
    },
    show() {
      this.$viewer.show();
    },
  }
};
</script>

<style></style>
