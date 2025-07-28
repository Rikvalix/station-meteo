import {defineStore} from "pinia";
import type MeasureModel from "../model/MeasureModel.ts";
import {measureApi} from "../api/lib/measureApi.ts";

export const useMeasureStore = defineStore('measure', {
    state: () => ({
        latestMeasure: {} as MeasureModel,
        measures : [] as MeasureModel[],
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

        async fetchMeasures(stationId: string, limit: number) {
            try {
                const result = await measureApi.getAllMeasuresByStationId(stationId,limit)
                if (result) {
                    this.measures = result
                }
            } catch (err) {
                console.error("Erreur lors de la récupération des mesures :", err)
            }
        }
    },
})
