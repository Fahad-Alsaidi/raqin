import { date } from 'quasar'

export default {
  methods: {
    Stage(value) {
      switch (value) {
        case "rev":
          return this.$t('bookStages.rev');
        case "done":
          return this.$t('bookStages.done');
        default:
          return this.$t('bookStages.init');
      }
    },
    getDateFormat: function (value) {
      return date.formatDate(value, 'YYYY-MM-DD')
    },
    getBooks: function () {
      this.$store.dispatch("editor/getBooks")
        .then(res => {
          this.books = res;
        })
        .catch(err => {
          console.log(err);
        })
    },
    BookWithID: function (books, id) {
      let book = books.find(e => e.id === id);
      return book ? book : { id: 0, name: "-" }
    }
  },
}