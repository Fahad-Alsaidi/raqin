<template>
  <div ref="linechart" id="lineChart"></div>
</template>
<script>
import cubejs from "@cubejs-client/core";
import ChartMixin from "../../mixins/chartMixin";
import { CUBE_URL, LineQuery } from "./cubeConfig";
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
      line_chart: null,
      resultSet: null,
      chartType: "line",
      option: {
        color: ["#716F81", "#B97A95", "#F6AE99"],
        xAxis: {
          type: "category",
          data: []
        },
        yAxis: {
          type: "value",
          minInterval: 1
        },
        title: {
          text: "الصفحات في كل يوم",
          left: "center"
        },
        tooltip: {
          trigger: "item",
          formatter: "{a} <br/>{b} : {c}"
        },
        legend: {
          orient: "vertical",
          left: "left",
          data: []
        },
        series: [
          {
            name: "الصفحات",
            type: "line",
            smooth: true,
            data: []
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
      let lineChart = document.getElementById("lineChart");
      echarts.dispose(lineChart);
      let theme = this.model ? "dark" : "light";
      this.line_chart = echarts.init(lineChart, theme);
      let chartData = this.data(this.resultSet)[0]["data"];
      this.option.series[0].data = chartData.map(function(item) {
        return item[1];
      });
      this.option.xAxis.data = chartData.map(function(item) {
        return new Date(item[0]).toDateString();
      });
      this.line_chart.setOption(this.option);
    },
    queryCube() {
      cubejsApi
        .load(LineQuery)
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
