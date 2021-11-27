import axios from 'axios'

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

  instance.interceptors.request.use(function (config) {
    
    config.baseURL= baseURL;
  
    if (store && store.getters['auth/getToken']) {
      config.headers = {"Authorization": `Bearer ${store.getters['auth/getToken']}`}
    }

    return config
  }, function (error) {
    // if (error.response && error.response.status === 401) {
    //   // not authorized, redirect to login page
    //   router.app.$root.$emit('auth:token-expired')
    // }
    return Promise.reject(error)
  })

}