<template>
  <q-page class="flex flex-center">
    <div class="q-pa-md row items-start q-gutter-md" style="direction: rtl;">

    <q-card class="my-card q-gutter-col-md" flat bordered>
      <q-card-section>

        <div class="text-h5 q-mt-sm q-mb-md"> ارفع كتابا جديدا للرقن </div>

        <q-file
          v-model="files"
          label="ارفع كتاب"
          square
          flat
          counter
          outlined
          use-chips
          
          clearable
          accept=".csv,.txt,.xls,.xlsx,.doc,.docx,.pdf,.dbf,.zip,.rar,.7z,.jpg,.png,.gif"
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
            <div class="text-overline text-orange-9 q-ml-md">إسم المؤلف</div>
          </template>
        </q-input>

      </q-card-section>

      <q-card-actions>
        <q-btn flat color="primary" label="رفع الكتاب" />
      </q-card-actions>

    </q-card>

  </div>
  </q-page>
</template>

<script>

export default {
  name: "project",
  data() {
    return {
      files: null,
      book: null,
      author: null
    };
  },
  methods: {
    onSubmit (evt) {
      this.loading = true // add loading state to submit button
      const formData = new FormData()

      if (this.files && this.files.length > 0) {
        for (let i = 0; i < this.files.length; i++) {
          formData.append('files[' + i + ']', this.files[i])
        }
      }
      for (const [key, value] of Object.entries(this.form)) {
        formData.append(key, value)
      }

      this.$axios.get('/sanctum/csrf-cookie').then(response => {
        this.$axios({
          method: 'post',
          url: '/api/request',
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
              message: 'You can upload up to 10 files.'
            })

            break
        }
      }
    }
  }
};
</script>

<style lang="sass" scoped>
.my-card
  width: 100%
  width: 500px
  max-width: 650px
</style>
