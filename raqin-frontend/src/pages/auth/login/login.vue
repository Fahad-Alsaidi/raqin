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
            <div class="flex flex-center tw-text-lg">سجل الدخول</div>
            <q-card-section>
              <q-form class="q-gutter-md">
                <q-input
                  dense
                  outlined
                  v-model="user.identifier"
                  type="identifier"
                  label="المعرف"
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
                label="دخول"
                @click="login()"
              />
            </q-card-actions>
            <q-card-section class="text-center q-pa-none">
              <div class>
                ليس لديك حساب؟
                <a href="/register" class="text-primary">اشترك اﻵن</a>
              </div>
            </q-card-section>
          </q-card>
        </div>
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script>
export default {
  name: "login",
  data() {
    return {
      user: {
        identifier: "",
        password: ""
      }
    };
  },
  methods: {
    login() {
      if (this.$store.getters["auth/isLoggedIn"]) {
        this.$q.notify({
          message: "يبدو أنك لم تسجل الخروج بعد, سيتم تحويلك للصفحة الرئيسية",
          color: "primary",
          textColor: "white",
          icon: "cloud_done"
        });
        setTimeout(() => {
          this.$router.push("/home");
        }, 2000);
      } else {
        this.$store
          .dispatch("auth/login", this.user)
          .then(e => {
            this.$router.push("/home");
          })
          .catch(err => console.log(err));
      }
    }
  }
};
</script>

<style scoped></style>
