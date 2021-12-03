<template>
  <q-layout view="lHh Lpr lFf" style="direction: rtl">
    <q-page-container>
      <q-page>
        <div class="bg-content window-height window-width row justify-center items-center">
          <q-card
            square
            bordered
            class="tw-w-full lg:tw-w-1/3 md:tw-w-1/2 tw-p-6 tw-border-t-2 tw-border-l-2 bg-content"
          >
            <div class="flex flex-center tw-text-lg">تسجيل حساب</div>

            <q-card-section>
              <q-form class="q-gutter-md">
                <q-input
                  dense
                  outlined
                  v-model="user.username"
                  type="username"
                  label="المعرف"
                  class="bg-navbar"
                />
                <q-input
                  dense
                  outlined
                  v-model="user.email"
                  type="email"
                  label="البريد اﻹلكتروني"
                  class="bg-navbar"
                />
                <q-input
                  dense
                  outlined
                  v-model="user.password"
                  type="password"
                  label="الرمز"
                  class="bg-navbar"
                />
              </q-form>
            </q-card-section>
            <q-card-actions class="q-px-md flex flex-center">
              <q-btn
                unelevated
                color="primary"
                size="md"
                class="tw-w-1/2"
                label="تسجيل"
                @click="register()"
                :loading="loading"
              />
            </q-card-actions>
          </q-card>
        </div>
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script>
export default {
  name: "register",
  data() {
    return {
      loading: false,
      user: {
        email: "",
        username: "",
        password: ""
      }
    };
  },
  methods: {
    register() {
      if (this.$store.getters["auth/isLoggedIn"]) {
        this.$q.notify({
          message: "يبدو أنك لم تسجل الخروج بعد, سجل الخروج أولا",
          color: "primary",
          textColor: "white",
          icon: "cloud_done"
        });
      } else {
        this.loading = true;
        this.$store
          .dispatch("auth/register", this.user)
          .then(e => {
            this.$router.push("/home");
          })
          .catch(err => {
            this.loading = false;
            console.log(err);
          });
      }
    }
  }
};
</script>

<style scoped></style>
