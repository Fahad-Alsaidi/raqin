import Vue from 'vue'
import Vuex from 'vuex'
import SecureLS from "secure-ls";
import VuexPersist from "vuex-persist";

import auth from './module-auth'
import editor from './module-editor'
import setting from './setting'

const secureLocalStorage = new SecureLS({ isCompression: false, encodingType: "rc4" });
const vuexPersist = new VuexPersist({
  key: "raqin",
  saveState: (key, state) => secureLocalStorage.set(key, state),
  restoreState: key => secureLocalStorage.get(key)
});


Vue.use(Vuex)

/*
 * If not building with SSR mode, you can
 * directly export the Store instantiation;
 *
 * The function below can be async too; either use
 * async/await or return a Promise which resolves
 * with the Store instance.
 */

const Store = new Vuex.Store({
  plugins: [vuexPersist.plugin],
  modules: {
    auth,
    editor,
    setting
  },

  // enable strict mode (adds overhead!)
  // for dev mode only
  strict: process.env.DEBUGGING
})

export default function (/* { ssrContext } */){
  return Store
}

export { Store }