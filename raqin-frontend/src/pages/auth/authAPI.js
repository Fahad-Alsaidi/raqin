// import Store to get axios instance
import { Store } from 'src/store';
import { updateProfile } from 'src/store/module-auth/actions';

const resource = '/auth/local';

export default {
    
    register(payload) {
        return Store.getters['auth/getAxios'].post(`${resource}/register`, payload);
    },
    login(payload) {
        return Store.getters['auth/getAxios'].post(`${resource}`, payload);
    },
    myProfile(){
        return Store.getters['auth/getAxios'].get(`/users/me`);
    },
    updateProfile(payload){
        return Store.getters['auth/getAxios'].put(`/users-permissions/users/me`, payload);
    }
    
};