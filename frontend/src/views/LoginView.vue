<script setup lang="ts">

import {ref} from "vue";
import {useUserStore} from "../stores/UserStore.ts";
import CustomSnackbar from "../components/Snackbar.vue";
import {useRouter} from "vue-router";

const userStore = useUserStore();
const authForm = ref({
  pseudo: "",
  password: "",
})
const snackbar = ref({status: false, message: "", type: "primary"})
const seePassword = ref(false)
const router = useRouter()


const loginForm = async () => {
    const result = await userStore.loginUser(authForm.value.pseudo, authForm.value.password);
    if (result) {
      snackbar.value = {status: true, message: "Connecté", type: "success"}
      await router.push("/home")
    } else {
      snackbar.value = {status: true, message: "Erreur de connexion", type: "error"}
    }
}

</script>

<template>
  <v-container
  >
    <v-row justify="center">
      <v-col
          cols="auto"
      >
        <v-card
            variant="flat"
        >
          <v-card-title class="text-h4 font-weight-bold text-center d-flex justify-center align-center gap-2 mb-6">
            <v-icon size="32" color="primary">mdi-lock-open</v-icon>
            <span class="text-primary">Connexion</span>
          </v-card-title>


          <v-card-text>
            <v-form @submit.prevent="loginForm">

              <v-text-field
                  label="Pseudo"
                  prepend-inner-icon="mdi-email-outline"
                  v-model="authForm.pseudo"
                  color="primary"
                  autocomplete="username"
                  variant="outlined"
              />

              <v-text-field
                  label="Mot de passe"
                  prepend-inner-icon="mdi-lock"
                  :append-inner-icon="seePassword ? 'mdi-eye' : 'mdi-eye-off'"
                  v-model="authForm.password"
                  :type="seePassword ? 'text' : 'password'"
                  color="primary"
                  variant="outlined"
                  autocomplete="current-password"
                  @click:append-inner="seePassword = !seePassword"
              />

              <v-btn
                  class="mt-6"
                  color="primary"
                  variant="outlined"
                  rounded="lg"
                  size="large"
                  type="submit"
              >
                Connexion
              </v-btn>
            </v-form>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
  <CustomSnackbar :snackbar="snackbar" />
</template>



<style scoped>

</style>