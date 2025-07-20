<script setup lang="ts">
import HeaderComponent from "./components/header/HeaderComponent.vue";
import {useUserStore} from "./stores/UserStore.ts";
import {onMounted} from "vue";
import {useRouter} from "vue-router";

const userStore = useUserStore();
const router = useRouter();

onMounted(() => {
  if (userStore.getCurrentUser === null) {
    router.push("/login");
  }
})
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
