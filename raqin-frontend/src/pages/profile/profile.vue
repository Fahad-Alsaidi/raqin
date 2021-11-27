<template>
  <q-page class="flex flex-center">
    <ProfileCard />
  </q-page>
</template>

<script>
import ProfileCard from "./components/profileCard.vue";
export default {
  name: "PageIndex",
  components:{
    ProfileCard
  }
};
</script>

<style lang="sass" scoped>
</style>
