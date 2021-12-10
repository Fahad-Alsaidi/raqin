import { QueryRenderer } from "@cubejs-client/vue";

export default {
    components: {
        QueryRenderer,
      },

    methods: {
        data(resultSet) {
            if (['line', 'area'].includes(this.chartType)) {
                return this.series(resultSet);
            }

            if (this.chartType === 'pie') {
                return this.pairs(resultSet);
            }

            if (this.chartType === 'bar') {
                return this.seriesPairs(resultSet);
            }
        },

        series(resultSet) {
            if (!resultSet) {
                return [];
            }

            const seriesNames = resultSet.seriesNames();
            const pivot = resultSet.chartPivot();
            const series = [];
            seriesNames.forEach(e => {
                const data = pivot.map(p => [p.x, p[e.key]]);
                series.push({
                    name: e.title,
                    data
                });
            });
            return series;
        },

        pairs(resultSet) {
            let res = resultSet.series()[0].series.map(item => {return {value: item.value, name: this.Stage(item.x)}});
            return res;
        },

        seriesPairs(resultSet) {
            return resultSet.series().map(seriesItem => ({
                name: seriesItem.title,
                data: seriesItem.series.map(item => [item.x, item.value])
            }));
        },

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
    },

    computed: {
        componentType() {
          if (this.chartType === 'table' || this.chartType === 'number') {
            return null;
          }
    
          return [this.chartType === 'bar' ? 'column' : this.chartType, '-chart'].join('');
        },
    
        isStacked() {
          return this.chartType === 'area';
        }
    
      },
}