<script setup lang="ts">
import {useUserStore} from "../stores/UserStore.ts";
import {onMounted, watch, ref} from "vue";

const userStore = useUserStore()

onMounted(async () => {
  await userStore.getAllStations()
})

// Loader
const stationLoad = ref(true)

watch(userStore.stations, (newVal) => {
  if (newVal.length > 0) {
    stationLoad.value = false
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
          <v-card-text>
            <v-row>
              <v-col cols="6">
                <v-card
                    variant="plain"
                >
                  <v-card-title>
                    <h3>21.2°c</h3>
                  </v-card-title>
                </v-card>
              </v-col>
              <v-col cols="6">
                <v-card
                    variant="plain"
                >
                  <v-card-title>
                    <h3>52%</h3>
                  </v-card-title>
                </v-card>
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