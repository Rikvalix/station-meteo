import {client} from "../client.ts";

export const userApi = {

    loginUser: async (username: string, password: string) => {
        const result = await client.post('/user/login', {
            "username": username,
            "password": password
        })
        if (result) {
            return result.data
        } else {
            return null
        }
    },
}