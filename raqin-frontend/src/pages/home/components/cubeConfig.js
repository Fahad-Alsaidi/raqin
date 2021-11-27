export var CUBE_URL;
if (process.env.NODE_ENV == 'production') {
    CUBE_URL = `https://${process.env.CUBEJS_URL}/cubejs-api/v1`;
} else {
    CUBE_URL = `http://localhost:4000/cubejs-api/v1`;
}


export const PieQuery = {
    measures: ["Books.count"],
    timeDimensions: [],
    order: {
        "Books.count": "desc",
    },
    limit: 5000,
    filters: [],
    dimensions: ["Books.stage"],
}

export const LineQuery = {
    measures: [
        "Pages.count"
    ],
    timeDimensions: [
        {
            "dimension": "Pages.createdAt",
            "granularity": "day"
        }
    ],
    order: {},
    dimensions: []
}

export const AnotherLineQuery = {
    dimensions: [
        "Pages.stage"
    ],
    timeDimensions: [
        {
            "dimension": "Pages.createdAt",
            "granularity": "day"
        }
    ],
    order: {
        "Pages.count": "desc"
    },
    measures: [
        "Pages.count"
    ],
    filters: [],
    limit: 5000
}