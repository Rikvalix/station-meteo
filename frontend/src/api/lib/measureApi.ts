import {client} from "../client.ts";

export const measureApi = {
    getAllStations: async () => {
        const result = await client.get("/api/v1/user/station")
        if (result.status === 200) {
            return result.data
        } else {
            return []
        }
    }
}