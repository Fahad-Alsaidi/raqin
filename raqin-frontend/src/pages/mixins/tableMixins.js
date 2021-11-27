import { date } from 'quasar'

export default {
  data(){
    return {
      books: [],
    }
  },
  mounted(){
    this.getBooks();
  },
  methods: {
    Stage: function (value) {
      switch (value) {
        case "rev":
          return "مراجعة";
        case "done":
          return "منتهي";
        default:
          return "مبدئي";
      }
    },
    getDateFormat: function(value){
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
    BookWithID: function(id){
      let book = this.books.find(e => e.id === id);
      return book ? book : {id: 0, name: "-"}
    }
  },
}