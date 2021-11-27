// import Store to get axios instance
import { Store } from 'src/store';

const resource = '/pages';

export default {
    
    get() {
        return Store.getters['auth/getAxios'].get(`${resource}`);
    },
    getByPageNumberAndBookID(pageNumber, bookID) {
        return Store.getters['auth/getAxios'].get(`${resource}?pageNumber=${pageNumber}&book.id=${bookID}`);
    },
    getByUserID(userID) {
        return Store.getters['auth/getAxios'].get(`${resource}?user=${userID}`);
    },
    getBookCompletion(bookID) {
        return Store.getters['auth/getAxios'].get(`${resource}/count-done-of-book/${bookID}`);
    },
    getPage(id) {
        return Store.getters['auth/getAxios'].get(`${resource}/${id}`);
    },
    create(payload) {
        return Store.getters['auth/getAxios'].post(`${resource}`, payload);
    },
    update(payload, id) {
        return Store.getters['auth/getAxios'].put(`${resource}/${id}`, payload);
    },
    delete(id) {
        return Store.getters['auth/getAxios'].delete(`${resource}/${id}`)
    },

};