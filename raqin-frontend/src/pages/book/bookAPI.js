// import Store to get axios instance
import { Store } from 'src/store';

const resource = '/books';

export default {
    
    get() {
        return Store.getters['auth/getAxios'].get(`${resource}`);
    },
    getBook(id) {
        return Store.getters['auth/getAxios'].get(`${resource}/${id}`);
    },
    // create(payload) {
    //     return Store.getters['auth/getAxios'].post(`${resource}`, payload);
    // },
    // update(payload, id) {
    //     return Store.getters['auth/getAxios'].put(`${resource}/${id}`, payload);
    // },
    // delete(id) {
    //     return Store.getters['auth/getAxios'].delete(`${resource}/${id}`)
    // },

};