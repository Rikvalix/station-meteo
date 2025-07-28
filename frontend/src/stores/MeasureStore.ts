import {defineStore} from "pinia";
import type MeasureModel from "../model/MeasureModel.ts";
import {measureApi} from "../api/lib/measureApi.ts";

export const useMeasureStore = defineStore('measure', {
    state: () => ({
        latestMeasure: {} as MeasureModel,
    }),

    actions: {
        async fetchLatestMesure(stationId: string) {
            try {
                const result = await measureApi.getLatestMesureByStationId(stationId)
                if (result) {
                    this.latestMeasure = result
                }
            } catch (err) {
                console.error("Erreur lors de la récupération de la mesure :", err)
            }
        },
    },
})
