export function setUser(state, user) {
    state.user = user;
}

export function setToken(state, token) {
    state.token = token;
}

export function setAxiosInstance(state, instance) {
    state.axiosInstance = instance;
}

export function clearUser(state){
    state.user = null;
}

export function clearToken(state){
    state.token = '';
}