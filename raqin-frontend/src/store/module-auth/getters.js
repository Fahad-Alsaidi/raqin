
export function isLoggedIn(state) {
    return !!state.token;
}

export function getToken(state) {
    return state.token;
}

export function getAxios(state) {
    return state.axiosInstance;
}

export function getUser(state) {
    return state.user
}