<script setup lang="ts">
import {useUserStore} from "../stores/UserStore.ts";
import {onMounted, watch, ref, computed} from "vue";
import {useMeasureStore} from "../stores/MeasureStore.ts";
import {format} from 'date-fns'
import LineChart from "../components/charts/LineChart.vue";
import type MeasureModel from "../model/MeasureModel.ts";
import {dateUtils} from "../utils/dateUtils.ts";
import CustomSnackbar from "../components/Snackbar.vue";

const userStore = useUserStore()
const measureStore = useMeasureStore()
const snackbar = ref({status: false, message: "", type: "primary"})

onMounted(async () => {
  await userStore.getAllStations()
})


// Loader
const stationLoad = ref(true)
const measuresLoad = ref(true)

// Config
const limit = ref(1000)

const currentStation = computed(() => {
  if (userStore.stations[0]) {
    return userStore.stations[0]
  }
})

watch(
    () => userStore.stations,
    async (newVal) => {
      if (newVal && newVal.length > 0) {
        stationLoad.value = false
        try {
          if (currentStation.value) {
            await measureStore.fetchLatestMesure(currentStation.value.public_id)
            await measureStore.fetchMeasures(currentStation.value.public_id, limit.value)
            if (measureStore.latestMeasure) {
              measuresLoad.value = false
            }
          }
        } catch (err) {
         snackbar.value = {status: true, message: "Erreur lors de la récupération des mesures", type: "error"}
        }
      }
    },
    {immediate: true}
)

watch(
    () => limit.value,
    async (newVal) => {
      measuresLoad.value = true
      try {
        if (currentStation.value) {
          await measureStore.fetchMeasures(currentStation.value.public_id, newVal)
        }
      } catch (err) {
        snackbar.value = {status: true, message: "Erreur lors de la récupération des mesures", type: "error"}
      }
    }
)

watch(() => measureStore.measures, (newVal) => {
  console.log("updates valeures")
  if (newVal && newVal.length > 0) {
    measuresLoad.value = false
  }
})



</script>

<template>
  <v-container fluid>
    <v-row>
      <v-col cols="12">
        <v-card
            class="pa-2"
            rounded="lg"
        >
          <v-card-title>
            Station actuelle
          </v-card-title>
          <v-card-text>
            <v-skeleton-loader
                :loading="stationLoad"
                type="paragraph"
            >
              {{ userStore.stations[0].name }} - {{ userStore.stations[0].address }}
            </v-skeleton-loader>
          </v-card-text>
          <v-card-actions>
            <v-select
                label="Changer de station"
                :items=userStore.stations
                item-title="name"
                variant="outlined"
                v-model="currentStation"
            />

          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12">
        <v-card
            rounded="lg">
          <v-card-title>Actuellement</v-card-title>
          <v-skeleton-loader
              :loading="measuresLoad"
              type="paragraph"
          >
            <v-card-subtitle>Dernière mesure: {{
                format(new Date(measureStore.latestMeasure.date), 'dd/MM/yyyy HH:mm')
              }}
            </v-card-subtitle>
          </v-skeleton-loader>
          <v-card-text>
            <v-row>
              <v-col cols="6">
                <v-skeleton-loader
                    :loading="measuresLoad"
                    type="paragraph"
                >
                  <v-card
                      variant="plain"
                  >
                    <v-card-title>
                      <h3>{{ measureStore.latestMeasure.temperature }}°c</h3>
                    </v-card-title>
                  </v-card>
                </v-skeleton-loader>
              </v-col>
              <v-col cols="6">
                <v-skeleton-loader
                    :loading="measuresLoad"
                    type="paragraph"
                >
                  <v-card
                      variant="plain"
                  >
                    <v-card-title>
                      <h3>{{ measureStore.latestMeasure.humidity }}%</h3>
                    </v-card-title>
                  </v-card>
                </v-skeleton-loader>
              </v-col>
            </v-row>


          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12">
        <v-card>
          <v-row justify="space-between">
            <v-col cols="auto">
              <v-card-title>Relevés</v-card-title>
            </v-col>
            <v-col cols="auto">
              <v-select
                  variant="plain"
                  :items="[10,100,1000,2000,5000,10000]"
                  v-model="limit"
                  color="primary"
              />
            </v-col>
          </v-row>
          <v-card-subtitle>Température</v-card-subtitle>
          <v-card-text>
            <v-skeleton-loader
                :loading="measuresLoad"
                type="paragraph"
            >
              <LineChart
                  :data="measureStore.measures.map((item: MeasureModel) => item.temperature)"
                  label="Température"
                  :labels="dateUtils.groupDate(measureStore.measures)"
                  color="primary"
              />
            </v-skeleton-loader>
          </v-card-text>
          <v-card-subtitle>Humidité</v-card-subtitle>
          <v-card-text>
            <v-skeleton-loader
                :loading="measuresLoad"
                type="paragraph"
            >
              <LineChart
                  :data="measureStore.measures.map((item: MeasureModel) => item.humidity)"
                  label="Humidité"
                  :labels="dateUtils.groupDate(measureStore.measures)"
                  color="secondary"
              />
            </v-skeleton-loader>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

  </v-container>
  <CustomSnackbar :snackbar="snackbar" />

</template>

<style scoped>

</style>