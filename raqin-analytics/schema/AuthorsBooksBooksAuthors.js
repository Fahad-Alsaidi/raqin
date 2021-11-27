cube(`AuthorsBooksBooksAuthors`, {
  sql: `SELECT * FROM \`raqin-db\`.authors_books__books_authors`,
  
  preAggregations: {
    // Pre-Aggregations definitions go here
    // Learn more here: https://cube.dev/docs/caching/pre-aggregations/getting-started  
  },
  
  joins: {
    Authors: {
      sql: `${CUBE}.author_id = ${Authors}.id`,
      relationship: `belongsTo`
    },
    
    Books: {
      sql: `${CUBE}.book_id = ${Books}.id`,
      relationship: `belongsTo`
    }
  },
  
  measures: {
    count: {
      type: `count`,
      drillMembers: [id]
    }
  },
  
  dimensions: {
    id: {
      sql: `id`,
      type: `number`,
      primaryKey: true
    }
  },
  
  dataSource: `default`
});
