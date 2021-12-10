import axios from 'axios'
import { Loading, QSpinnerGears, Notify } from 'quasar'

export default ({ store, Vue }) => {

  let baseURL;
  if (process.env.NODE_ENV == 'production') {
    baseURL = `https://${process.env.API_URL}`;
  } else {
    baseURL = `http://localhost:1337`;
  }

  const instance = axios.create()

  store.commit('auth/setAxiosInstance', instance)
  Vue.prototype.$axios = instance

  // attach request interceptor
  instance.interceptors.request.use(config => {

    // show spinner before request is sent
    Loading.show({
      spinner: QSpinnerGears,
    })

    config.baseURL = baseURL;

    // attach token to request if exits
    if (store && store.getters['auth/getToken']) {
      config.headers = { "Authorization": `Bearer ${store.getters['auth/getToken']}` }
    }

    return config
  }, error => {

    Loading.hide()

    Notify.create({
      type: 'negative',
      color: 'negative',
      textColor: 'white',
      timeout: 3000,
      position: 'bottom',
      message: error.message || 'Oops. Something went wrong!'
    })
    return Promise.reject(error)
  })

  // attach response interceptor
  instance.interceptors.response.use(response => {
    Loading.hide()
    return response
  }, ({ response }) => {
    let errMsg;
    if (Array.isArray(response?.data?.message)) {
      errMsg = response?.data?.error;
    } else errMsg = `${response?.data?.error}: ${response?.data?.message}`
    Notify.create({
      type: 'negative',
      color: 'negative',
      textColor: 'white',
      timeout: 3000,
      position: 'bottom',
      message: errMsg || "something went wrong"
    })
    
    Loading.hide()

    return Promise.reject(response.data)
  })

}