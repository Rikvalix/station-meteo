<script setup lang="ts">
import {useUserStore} from "../stores/UserStore.ts";
import {onMounted, watch, ref, computed} from "vue";
import {useMeasureStore} from "../stores/MeasureStore.ts";
import { format } from 'date-fns'

const userStore = useUserStore()
const measureStore = useMeasureStore()

onMounted(async () => {
  await userStore.getAllStations()
})


// Loader
const stationLoad = ref(true)
const measuresLoad = ref(true)

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
            if (measureStore.latestMeasure) {
              measuresLoad.value = false
            }
          }
        } catch (err) {
          console.error("Erreur lors de la récupération des mesures :", err)
        }
      }
    },
    {immediate: true}
)


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
            <v-card-subtitle>Dernière mesure: {{ format(new Date(measureStore.latestMeasure.date), 'dd/MM/yyyy HH:mm') }}</v-card-subtitle>
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
          <v-card-title>Relevés</v-card-title>
          <v-card-text>

          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

  </v-container>

</template>

<style scoped>

</style>