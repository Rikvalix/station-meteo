<script setup lang="ts">
import {defineComponent} from "vue";

defineComponent({
  name: 'CustomSnackbar',
})

interface SnackbarInterface {
  type: string;
  message: string;
  status: boolean;
}

const props = defineProps<{
  snackbar: SnackbarInterface
}>()

</script>

<template>
  <v-snackbar
      rounded="pill"
      :color="props.snackbar.type"
      v-model="props.snackbar.status">

    {{props.snackbar.message}}
  </v-snackbar>
</template>

<style scoped>

</style>