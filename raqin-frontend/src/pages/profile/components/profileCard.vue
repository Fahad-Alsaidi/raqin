<template>
  <q-card
    flat
    bordered
    class="tw-w-full lg:tw-w-1/3 md:tw-w-1/2 tw-p-6 tw-border-t-2 tw-border-l-2 bg-content"
  >
    <q-card-section>
      <q-input
        dense
        outlined
        :label="$t('profilePage.userName')"
        :disable="!edit"
        v-model="username"
        type="text"
        class="bg-navbar"
      />
      <q-input
        dense
        outlined
        :label="$t('profilePage.email')"
        :disable="!edit"
        v-model="email"
        type="email"
        class="tw-mt-4 bg-navbar"
      />
    </q-card-section>

    <q-card-actions class="tw-flex tw-flex-row">
      <q-btn color="primary" @click="edit = !edit">
        {{
        edit ? $t('profilePage.cancel') : $t('profilePage.edit')
        }}
      </q-btn>

      <q-space />

      <q-btn color="primary" v-if="edit" :loading="loading" @click="updateProfile()">{{$t('profilePage.save')}}</q-btn>
    </q-card-actions>
  </q-card>
</template>

<script>
import notifyMixin from "../../mixins/notifyMixin";

export default {
  mixins: [notifyMixin],
  data: function() {
    return {
      loading: false,
      edit: false,
      username: "",
      email: ""
    };
  },
  mounted() {
    this.getUser();
  },
  methods: {
    getUser() {
      let user = this.$store.getters["auth/getUser"];
      if (user) {
        this.username = user.username;
        this.email = user.email;
      }
    },
    updateProfile() {
      if (this.username && this.email) {
        this.loading = true;
        this.$store
          .dispatch("auth/updateProfile", {
            username: this.username,
            email: this.email
          })
          .then(res => {
            this.loading = false;
            this.$q.notify(this.showSuccess("updated successfully"));
            this.edit = !this.edit;
          })
          .catch(err => {
            this.loading = false;
            this.$q.notify("update failed");
          });
      }
    }
  }
};
</script>
<style scoped></style>