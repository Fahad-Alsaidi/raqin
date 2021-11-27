<template>
    <q-card class="my-card q-gutter-col-md" flat bordered style="direction: rtl;">
      <q-card-section>

        <div class="text-h5 q-mt-sm q-mb-md"> ابدأ مشروع رقن جديد </div>

        <q-file
          v-model="file"
          label="اختر كتاب"
          square
          flat
          counter
          outlined
          use-chips
          
          clearable
          accept=".pdf"
          max-files="1"
          max-file-size="5120000"
          @rejected="onRejected"
        >
          <template v-slot:prepend>
            <q-icon name="attach_file" />
          </template>
        </q-file>
        
        <q-input outlined v-model="book" class="q-my-sm">
          <template v-slot:prepend>
            <div class="text-overline text-orange-9 q-ml-md">إسم الكتاب</div>
          </template>
        </q-input>

        <q-input outlined v-model="author" class="q-my-sm">
          <template v-slot:prepend>
            <div class="text-overline text-orange-9 q-ml-md">التصنيف</div>
          </template>
        </q-input>

        <q-input outlined v-model="author" class="q-my-sm">
          <template v-slot:prepend>
            <div class="text-overline text-orange-9 q-ml-md">المؤلفون</div>
          </template>
        </q-input>

        <q-input outlined v-model="author" class="q-my-sm">
          <template v-slot:prepend>
            <div class="text-overline text-orange-9 q-ml-md">ملاحظات</div>
          </template>
        </q-input>

      </q-card-section>

      <q-card-actions>
        <q-btn flat color="primary" label="رفع الكتاب" @click="onSubmit" />
      </q-card-actions>

    </q-card>
</template>

<script>
export default {
    name: "book-project",
    data() {
    return {
      file: null,
      book: null,
      author: null
    };
  },
  methods: {
    onSubmit (evt) {
      this.loading = true // add loading state to submit button
      const formData = new FormData()
      formData.append("name", book)
      formData.append("author", author)
      formData.append("file", this.file)

      this.$axios({
          method: 'post',
          url: '/admin/upload',
          data: formData,
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        })
          .then((response) => {
            this.loading = false

            this.$q.notify({
              color: 'positive',
              position: 'top',
              message: 'Request sent! We\'ll contact you soon.',
              icon: 'done'
            })
          })
          .catch(() => {
            this.loading = false

            this.$q.notify({
              color: 'negative',
              position: 'top',
              message: 'An error occurred',
              icon: 'report_problem'
            })
          })
    },
              
    onRejected (entries) {
      if (entries.length > 0) {
        switch (entries[0].failedPropValidation) {
          case 'max-file-size':
            this.$q.notify({
              position: 'top',
              type: 'negative',
              message: 'File exceeds 5MB.'
            })

            break

          case 'max-files':
            this.$q.notify({
              position: 'top',
              type: 'negative',
              message: 'You can upload only 1 file.'
            })

            break
        }
      }
    }
  }
}
</script>

<style lang="sass" scoped>
.my-card
  width: 100%
  width: 500px
  max-width: 650px
</style>
