export default {
  helperWords:{
    of: 'of'
  },
  layout: {
    appName: "raqin",
    logout: "logout",
    menu: {
      home: 'Home',
      books: 'Books',
      myPages: 'My Pages',
      profile: 'Profile'
    }
  },
  bookStages: {
    init: 'Initial',
    rev: 'Revision',
    done: 'Done'
  },
  homePage: {
    noCharts: 'Sorry, no data!',
    bookChart: {
      title: 'Books category of completion',
      seriesName: 'Category of completion'
    },
    pageChart: {
      title: 'Proof-read pages by days',
      seriesName: 'Page count'
    }
  },
  booksPage: {
    table: {
      title: 'Books',
      rowsPerPage: 'rows in page:',
      contributeBtn: 'Contribute',
      columns: {
        first: 'Book Name',
        second: 'Category',
        third: 'Authors',
        fourth: 'Number of Pages',
        fifth: 'Completion',
        sixth: 'Stage',
        seventh: ''
      }
    }
  },
  myPagesPage: {
    table: {
      title: 'My page contribution',
      rowsPerPage: 'rows in page:',
      contributeBtn: 'Contribute',
      columns: {
        first: 'Book',
        second: 'Text',
        third: 'Editor',
        fourth: 'Page Number',
        fifth: 'Stage',
        sixth: ''
      }
    }
  },
  profilePage: {
    userName: 'Name',
    email: 'Email',
    edit: 'Edit',
    cancel: 'Cancel',
    save: 'Save'
  }
}
