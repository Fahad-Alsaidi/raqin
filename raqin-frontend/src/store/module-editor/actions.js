import booksAPI from "pages/book/bookAPI";
import pagesAPI from "src/pages/my-pages/pageAPI";

export function getBooks({ commit, state }) {
  return new Promise((resolve, reject) => {
    booksAPI.get()
      .then(res => {
        // commit('setBooks', res.data);
        resolve(res.data);
      })
      .catch(err => {
        reject(err);
      });
  });
}

export function pagesByUserID({ commit, state }) {
  return new Promise((resolve, reject) => {
    pagesAPI.getByUserID(this.getters['auth/getUser'].id)
      .then(res => {
        resolve(res.data);
      })
      .catch(err => {
        reject(err);
      });
  });
}

export function createPage({ commit, state }, page) {
  return new Promise((resolve, reject) => {
    pagesAPI.create(page)
      .then(res => {
        resolve(res);
      })
      .catch(err => {
        reject(err);
      });
  });
}

export function updatePage({ commit, state }, { page, pageID }) {
  return new Promise((resolve, reject) => {
    pagesAPI.update(page, pageID)
      .then(res => {
        resolve(res);
      })
      .catch(err => {
        reject(err);
      });
  });
}

export function donePagesCountByBookId({ commit, state }, bookID) {
  return new Promise((resolve, reject) => {
    pagesAPI.countDonePagesByBookID(bookID)
      .then(res => {
        resolve(res.data);
      })
      .catch(err => {
        reject(err);
      });
  });
}