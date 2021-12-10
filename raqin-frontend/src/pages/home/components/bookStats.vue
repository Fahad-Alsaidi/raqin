<template>
  <div ref="piechart" id="pieChart"></div>
</template>
<script>
import cubejs from "@cubejs-client/core";
import ChartMixin from "../../mixins/chartMixin";
import { CUBE_URL, PieQuery } from "./cubeConfig";
const cubejsApi = cubejs({
  apiUrl: CUBE_URL
});

export default {
  name: "ChartRenderer",
  mixins: [ChartMixin],
  data() {
    return {
      loading: true,
      cubejsApi,
      model: false,
      pie_chart: null,
      resultSet: null,
      chartType: "pie",
      option: {
        color: ["#716F81", "#B97A95", "#F6AE99"],
        title: {
          text: this.$t('homePage.bookChart.title'),
          left: "center"
        },
        tooltip: {
          trigger: "item",
          formatter: "{a} <br/>{b} : {c} ({d}%)"
        },
        legend: {
          orient: "vertical",
          left: "left",
          data: []
        },
        series: [
          {
            name: this.$t('homePage.bookChart.seriesName'),
            type: "pie",
            radius: "55%",
            center: ["50%", "60%"],
            data: [],
            emphasis: {
              itemStyle: {
                shadowBlur: 10,
                shadowOffsetX: 0,
                shadowColor: "rgba(0, 0, 0, 0.5)"
              }
            }
          }
        ]
      }
    };
  },
  mounted() {
    this.queryCube();
  },
  methods: {
    drawChart() {
      let pieChart = document.getElementById("pieChart");
      echarts.dispose(pieChart);
      let theme = this.model ? "dark" : "light";
      this.pie_chart = echarts.init(pieChart, theme);
      let chartData = this.data(this.resultSet);
      this.option.series[0].data = chartData;
      this.option.legend.data = chartData;
      this.pie_chart.setOption(this.option);
    },
    queryCube() {
      cubejsApi
        .load(PieQuery)
        .then(resp => {
          this.resultSet = resp;
          try {
            this.drawChart();
          } catch (error) {
            this.loading = false
          }
        })
        .catch(err => {
          this.loading = false
          console.log(err);
        });
    }
  }
};
</script>
