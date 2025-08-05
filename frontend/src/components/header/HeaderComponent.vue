<script setup lang="ts">
import CustomSnackbar from "../Snackbar.vue";
import {ref, shallowRef} from "vue"
import {useRouter} from "vue-router";

const snackbar = ref({status: false, message: "", type: "primary"})
const router = useRouter()
const drawer = shallowRef(false)

const handleLogout = () => {
  sessionStorage.clear()
  snackbar.value = {status: true, message: "Déconnecté", type: "primary"}
  router.push('/login')
}

</script>

<template>
  <v-app-bar app
             color="primary"
             dark
             scroll-behavior="collapse elevate"
  >
    <v-app-bar-nav-icon @click="drawer = !drawer"/>


    <v-toolbar-title>Station météo</v-toolbar-title>
    <template v-slot:append>
      <v-btn icon="mdi-power" @click="handleLogout"></v-btn>
    </template>

  </v-app-bar>
  <v-navigation-drawer v-model="drawer">
    <v-list>
      <v-list-item link to="/home">
        <v-list-item-title>Accueil</v-list-item-title>
      </v-list-item>
      <v-divider inset></v-divider>
      <v-list-item link to="/profil">
        <v-list-item-title>Profil</v-list-item-title>
      </v-list-item>
      <v-divider inset></v-divider>
      <v-list-item link to="/stations">
        <v-list-item-title>Toutes les stations</v-list-item-title>
      </v-list-item>
      <v-divider inset></v-divider>
    </v-list>
  </v-navigation-drawer>
  <CustomSnackbar :snackbar="snackbar"/>
</template>

<style scoped>

</style>