<template>
  <div>

    <div class="row justify-content-end my-3">
      <div class="col-12">
        <date-picker
          v-model="dateRange"
          @input="updateData"
          class="float-right"
          confirm
          range
          lang="en"
        ></date-picker>
      </div>
    </div>

    <!--Stats cards-->
    <div class="row">
      <div class="col-md-6 col-xl-3" v-for="stats in cardsData" :key="stats.title">
        <stats-card>
          <div class="icon-big text-center" :class="`icon-${stats.type}`" slot="header">
            <i :class="stats.icon"></i>
          </div>
          <div class="numbers" slot="content">
            <p>{{ stats.title }}</p>
            {{ stats.value }} {{ stats.unit }}
          </div>
          <div class="stats" slot="footer">
            <i class="ti-reload"></i> {{ stats.updated | moment("from") }}
          </div>
        </stats-card>
      </div>
    </div>

    <!--Charts-->
    <div class="row">
      <div class="col-12">
        <chart-card
          v-if="!isFetchingUsersBehavior"
          title="Users behavior"
          sub-title="24 Hours performance"
          :chart-data="usersBehavior"
          :chart-options="usersChart.options"
        >
          <span slot="footer">
            <i class="ti-reload"></i> {{ new Date() | moment("from") }}
          </span>
          <div slot="legend">
            <i class="fa fa-circle text-info"></i> Open
            <i class="fa fa-circle text-danger"></i> Click
            <i class="fa fa-circle text-warning"></i> Click Second Time
          </div>
        </chart-card>
      </div>

      <div class="col-md-6 col-12">
        <chart-card
          v-if="!isFetchingEmailStatistics"
          title="Email Statistics"
          sub-title="Last campaign performance"
          :chart-data="emailStatistics"
          chart-type="Pie"
        >
          <span slot="footer">
            <i class="ti-timer"></i> Campaign set 2 days ago</span>
          <div slot="legend">
            <i class="fa fa-circle text-info"></i> Open
            <i class="fa fa-circle text-danger"></i> Bounce
            <i class="fa fa-circle text-warning"></i> Unsubscribe
          </div>
        </chart-card>
      </div>

      <div class="col-md-6 col-12">
        <chart-card
          v-if="!isFetchingSalesData"
          title="2015 Sales"
          sub-title="All products including Taxes"
          :chart-data="salesData"
          :chart-options="salesChart.options"
        >
          <span slot="footer">
            <i class="ti-check"></i> Data information certified
          </span>
          <div slot="legend">
            <i class="fa fa-circle text-info"></i> Tesla Model S
            <i class="fa fa-circle text-warning"></i> BMW 5 Series
          </div>
        </chart-card>
      </div>

    </div>

  </div>
</template>
<script>
import { mapActions, mapGetters } from "vuex";
import { StatsCard, ChartCard } from "@/components/index";

import Chartist from "chartist";
import DatePicker from "vue2-datepicker";

export default {
  components: {
    StatsCard,
    ChartCard,
    DatePicker
  },
  /**
   * Chart data used to render stats, charts. Should be replaced with server data
   */
  methods: {
    ...mapActions("statistics", ["fetchCardsData"]),
    fetchUsersBehavior() {
      this.isFetchingUsersBehavior = true;
      this.$store
        .dispatch("statistics/fetchUsersBehavior", this.dateRangePayload)
        .catch(err => {
          console.log(err);
        })
        .finally(() => {
          this.isFetchingUsersBehavior = false;
        });
    },
    fetchEmailStatistics() {
      this.isFetchingEmailStatistics = true;
      this.$store
        .dispatch("statistics/fetchEmailStatistics", this.dateRangePayload)
        .catch(err => {
          console.log(err);
        })
        .finally(() => {
          this.isFetchingEmailStatistics = false;
        });
    },
    fetchSalesData() {
      this.isFetchingEmailStatistics = true;
      this.$store
        .dispatch("statistics/fetchSalesData", this.dateRangePayload)
        .catch(err => {
          console.log(err);
        })
        .finally(() => {
          this.isFetchingSalesData = false;
        });
    },
    updateData() {
      this.fetchCardsData();
      this.fetchUsersBehavior();
      this.fetchEmailStatistics();
      this.fetchSalesData();
    }
  },
  mounted() {
    this.updateData();
  },
  computed: {
    ...mapGetters("statistics", [
      "cardsData",
      "usersBehavior",
      "emailStatistics",
      "salesData"
    ]),
    startDate() {
      return this.dateRange[0].toISOString();
    },
    endDate() {
      return this.dateRange[1].toISOString();
    },
    dateRangePyalod() {
      return {
        start: this.startDate,
        end: this.endDate
      };
    }
  },
  data() {
    return {
      isFetchingUsersBehavior: true,
      isFetchingEmailStatistics: true,
      isFetchingSalesData: true,
      dateRange: [new Date(), new Date()],
      usersChart: {
        options: {
          low: 0,
          high: 1000,
          showArea: true,
          height: "245px",
          axisX: {
            showGrid: false
          },
          lineSmooth: Chartist.Interpolation.simple({
            divisor: 3
          }),
          showLine: true,
          showPoint: false
        }
      },
      salesChart: {
        options: {
          seriesBarDistance: 10,
          axisX: {
            showGrid: false
          },
          height: "245px"
        }
      }
    };
  }
};
</script>
