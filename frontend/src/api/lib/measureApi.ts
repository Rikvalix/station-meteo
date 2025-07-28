import {client} from "../client.ts";

export const measureApi = {
    getAllStations: async () => {
        const result = await client.get("/api/v1/user/station")
        if (result.status === 200) {
            return result.data
        } else {
            return []
        }
    },

    getLatestMesureByStationId : async (stationId: string) => {
        const result = await client.get("/api/v1/station/measurements/latest/" + stationId)
        if (result.status === 200) {
            return result.data
        } else {
            return null
        }
    },

    getAllMeasuresByStationId: async (stationId: string, limit: number) => {
        const result = await client.get("/api/v1/station/measurements?limit="+limit+"&station=" + stationId)
        if (result.status === 200) {
            return result.data
        }
        return null
    }
}