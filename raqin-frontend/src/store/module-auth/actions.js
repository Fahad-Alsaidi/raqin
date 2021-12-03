import authAPI from "pages/auth/authAPI";

export function login({commit}, user) {
  return new Promise((resolve, reject) => {
  authAPI
    .login(user)
    .then(res => {
      let token = res.data.jwt
      commit('setUser', res.data.user);
      commit('setToken', token);
      resolve(res);
    })
    .catch(err => {
      reject(err);
    });
  });
}

export function register({commit}, user) {
  return new Promise((resolve, reject) => {
  authAPI
    .register(user)
    .then(res => {
      let token = res.data.jwt
      commit('setUser', res.data.user);
      commit('setToken', token);
      resolve(res);
    })
    .catch(err => {
      reject(err);
    });
  });
}

export function logout({commit}) {
  return new Promise((resolve, reject) => {
      commit('clearUser');
      commit('clearToken');
      resolve();
  });
}

export function myProfile({commit}) {
  return new Promise((resolve, reject) => {
    authAPI
    .myProfile()
    .then(res => {
      commit('setUser', res.data);
      resolve(res.data);
    })
    .catch(err => {
      commit('clearUser');
      commit('clearToken');
      reject(err);
    });
  });
}

export function updateProfile({commit, state} , payload) {
  return new Promise((resolve, reject) => {
    authAPI
    .updateProfile(payload)
    .then(res => {
      commit('setUser', res.data);
      resolve(res.data);
    })
    .catch(err => {
      reject(err);
    });
  });
}