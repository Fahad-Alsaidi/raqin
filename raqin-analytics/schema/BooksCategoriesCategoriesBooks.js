cube(`BooksCategoriesCategoriesBooks`, {
  sql: `SELECT * FROM \`raqin-db\`.books_categories__categories_books`,
  
  preAggregations: {
    // Pre-Aggregations definitions go here
    // Learn more here: https://cube.dev/docs/caching/pre-aggregations/getting-started  
  },
  
  joins: {
    Categories: {
      sql: `${CUBE}.category_id = ${Categories}.id`,
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
