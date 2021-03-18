<template>
  <q-page class="flex flex-center">
    <div class="q-pa-md row items-start q-gutter-md" style="direction: rtl;">
    <q-card class="my-card" flat bordered>
      <viewer :img="[imgPath]" />
      <q-card-section>
        <q-input dense v-model="text">
          <template v-slot:prepend>
            <div class="text-overline text-orange-9">{{ question }}</div>
          </template>
        </q-input>
        <div class="text-h5 q-mt-sm q-mb-xs">{{ book }}</div>
        <div class="text-caption text-grey">
          {{ desription }}
        </div>
      </q-card-section>

      <q-card-actions>
        <q-btn flat color="dark" :label="change" @click="changePic()" />
        <q-btn flat color="primary" :label="solve" @click="SendResponse()" />

        <q-space />

        <q-btn
          color="grey"
          round
          flat
          dense
          :icon="expanded ? 'keyboard_arrow_up' : 'keyboard_arrow_down'"
          @click="expanded = !expanded"
        />
      </q-card-actions>

      <q-slide-transition>
        <div v-show="expanded">
          <q-separator />
          <q-card-section class="text-subitle2">
            {{ lorem }}
          </q-card-section>
        </div>
      </q-slide-transition>
    </q-card>
  </div>
  </q-page>
</template>

<script>
import axios from "axios";
import viewer from "components/viewer";
export default {
  name: "contribute",
  components: { viewer },
  data() {
    return {
      imgIndex: 0,
      text: "",
      expanded: false,
      question: "هل تستطيع حلها ؟",
      book: "من كتاب أنت أقوى",
      desription: "هذه الصورة من كتاب أنت أقوى من الصفحة 94, فهل تستطيع حلها ؟",
      solve: "حل",
      change: "تغيير الصورة",
      lorem:
        "تستخدم ترجمات هذه النصوص في بناء كتب يستطيع القارئ العربي الإستفادة منها وقرائتها بكل حرية."
    };
  },
  computed: {
    imgPath() {
      return "img/" + this.imgIndex + ".jpg";
    }
  },
  methods: {
    changePic() {
      if (this.imgIndex < 24) {
        this.imgIndex = this.imgIndex + 1;
        return;
      } else {
        return (this.imgIndex = 0);
      }
    },
    SendResponse() {
      axios
        .post(
          "localhost:1323/client/profile",
          { answer: this.text },
          {
            headers: {
              "Access-Control-Allow-Origin": "*"
            }
          }
        )
        .then(res => {
          console.log(res);
        })
        .catch(err => {
          console.log(err);
        });
    }
  }
};
</script>

<style lang="sass" scoped>
.my-card
  width: 100%
  max-width: 650px
</style>
