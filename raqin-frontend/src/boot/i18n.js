import Vue from 'vue'
import VueI18n from 'vue-i18n'
import messages from 'src/i18n'

Vue.use(VueI18n)

let i18n;

export default ({ app, store }) => {
  
  i18n = new VueI18n({
    locale: store.getters['setting/getLocale'] || 'ar',
    fallbackLocale: 'en-us',
    messages
  })
  // Set i18n instance on app
  app.i18n = i18n
}

export { i18n }
