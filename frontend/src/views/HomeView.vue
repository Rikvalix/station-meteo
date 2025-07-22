<script setup lang="ts">
import {useUserStore} from "../stores/UserStore.ts";
import {onMounted} from "vue";

const userStore = useUserStore()

onMounted(async () => {
  await userStore.getAllStations()
})

</script>

<template>
  <v-container fluid>
    <v-row align-content="start">
      <v-col cols="4">
        <v-card
            class="pa-2"
            rounded="lg"
        >
          <v-card-title>
            Station actuelle
          </v-card-title>
          <v-card-text>
            {{userStore.stations[0].name}} - {{userStore.stations[0].address}}
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
      <v-col cols="4">
        <v-card>
        <v-card-title>Actuellement</v-card-title>
          <v-card-text>

          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="4">
        <v-card>
          <v-card-title>Mes informations</v-card-title>
          <v-card-text>

          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

  </v-container>

</template>

<style scoped>

</style>