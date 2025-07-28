<script setup lang="ts">
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend
} from 'chart.js'
import {Line} from 'vue-chartjs'
import type {PropType} from "vue";
import {useTheme} from "vuetify/framework";


const theme = useTheme()

const colorMap = {
  "primary": theme.current.value.colors.primary,
  "secondary": theme.current.value.colors.secondary,
  "success": theme.current.value.colors.success,
  "warning": theme.current.value.colors.warning,
  "info": theme.current.value.colors.info,
  "error": theme.current.value.colors.error,

}

const props = defineProps({
  label: String,
  labels: {
    type: Array as PropType<string[]>,
    required: true
  },
  data: {
    type: Array as PropType<number[]>,
    required: true
  },
  color : {
    type: String,
    required: true,
  }
})

const data = {
  labels: props.labels,
  datasets: [
    {
      label: props.label,
      backgroundColor: colorMap[props.color as keyof typeof colorMap],
      data: props.data,
    },
  ]
}

const styles = {
  position: 'relative',

}
const options = {
  responsive: true,
  maintainAspectRatio: false
}

ChartJS.register(
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend
)

</script>

<template>
  <Line :data="data" :options="options" :style="styles"></Line>
</template>

<style scoped>

</style>