import state from './state'
import * as getters from './getters'
import * as mutations from './mutations'

export default {
  namespaced: true,
  getters,
  mutations,
  state
}
