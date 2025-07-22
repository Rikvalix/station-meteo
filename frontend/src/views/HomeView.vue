<script setup lang="ts">

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
            Le merlerault - Chambre
          </v-card-text>
          <v-card-actions>
            <v-select
                label="Changer de station"
                :items="['Le Merlerault', 'Caen', 'Nantes']"
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