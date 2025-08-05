<script setup lang="ts">

import {useUserStore} from "../stores/UserStore.ts";
import {onMounted, ref} from "vue";

const userStore = useUserStore();
const authKeySee = ref(false)

onMounted(async () => {
  if (userStore.stations.length === 0) {
    await userStore.getAllStations()
  }
})

</script>

<template>
  <v-container class="mt-5" fluid>
    <h1>Stations</h1>
    <v-container fluid>
      <v-row dense>
        <v-col
            cols="12"
            md="6"
            lg="4"
            v-for="station in userStore.stations"
            :key="station.name"
        >
          <v-card elevation="2">
            <v-card-title>
              <v-row>
                <v-col cols="12" class="text-primary font-weight-bold">
                  {{ station.name }}
                </v-col>
              </v-row>
            </v-card-title>
            <v-card-subtitle>
              <v-row>
                <v-col cols="12">
                  {{ station.address }} - <span class="text-info">{{ station.location }}</span>
                </v-col>
              </v-row>
            </v-card-subtitle>
            <v-divider inset></v-divider>
            <v-card-text>
              <v-row>
                <v-col cols="12" sm="12" md="6">
                  <v-list-item-title>Composants</v-list-item-title>
                  <v-list density="compact">
                    <v-list-item
                        v-for="(component, i) in station.components"
                        :key="i">
                      <v-list-item-title>{{ component }}</v-list-item-title>
                    </v-list-item>
                  </v-list>
                </v-col>

                <v-col cols="12" sm="12" md="6">
                  <v-text-field
                      variant="outlined"
                      v-model="station.auth_key"
                      :type="authKeySee ? 'text' : 'password'"
                      label="Clé d’authentification"
                  >
                    <template v-slot:append-inner>
                      <v-icon
                          :icon="authKeySee ? 'mdi-eye' : 'mdi-eye-off'"
                          @click="authKeySee = !authKeySee"
                          class="cursor-pointer"
                      />
                    </template>
                  </v-text-field>
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12" sm="12" md="6">
                  <v-chip
                      :color="station.enabled ? 'success' : 'error'"
                  >
                    {{
                      station.enabled ? 'Activé' : 'désactivé'
                    }}
                  </v-chip>
                </v-col>
              </v-row>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-container>

  </v-container>
</template>

<style scoped>
</style>