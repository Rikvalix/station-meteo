<script setup lang="ts">
import HeaderComponent from "./components/header/HeaderComponent.vue";

</script>

<template>
  <v-app>
    <HeaderComponent/>
    <v-main>
      <v-container class="pa-4">
        <RouterView/>
      </v-container>
    </v-main>
  </v-app>

</template>

<style scoped>

</style>
