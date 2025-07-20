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
    <v-row>
      <v-col
          cols="6"
      >
        <v-img src="/images/login_view.svg" alt="login image"  />
      </v-col>
      <v-col
          cols="6"
      >
        <v-card
            class="pa-8"
            variant="flat"
            max-width="400"
            width="100%"
        >
          <v-card-title class="text-h4 text-center mb-6">
            Connexion
          </v-card-title>

          <v-card-text>
            <v-form @submit.prevent="loginForm">
              <v-text-field
                  label="Pseudo"
                  prepend-inner-icon="mdi-email-outline"
                  v-model="authForm.pseudo"
                  color="primary"
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