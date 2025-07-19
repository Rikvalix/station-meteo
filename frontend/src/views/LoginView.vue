<script setup lang="ts">

import {ref} from "vue";

const authForm = ref({
  pseudo: "",
  password: "",
})


const loginForm = () => {

}

</script>

<template>
  <v-container
  >
    <v-row>
      <v-col
          cols="6"
      >
        <v-img src="/public/images/login_view.svg" alt="login image"  />
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
                  v-model="authForm.password"
                  color="primary"
                  variant="outlined"
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
</template>



<style scoped>

</style>