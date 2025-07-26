import axios from "axios";
import router from "../router/router.ts";

const baseUrl = import.meta.env.VITE_BASE_API_URL;

export const client = axios.create({
    baseURL :baseUrl,
    headers : {
        'Content-Type': 'application/json',
    }
})

client.interceptors.request.use(
    config => {
        const apiKey = sessionStorage.getItem("user")
        if (apiKey) {
            config.headers.Authorization = JSON.parse(apiKey).authToken;
        }else {
            delete client.defaults.headers.common.Authorization;
        }
        return config;
    },
    error => Promise.reject(error)
)

client.interceptors.response.use(
    function (response) {
        return response
    },
    async function (error) {
        if (error.response && error.response.status === 401) {
            await router.push("/login")
        }
        return Promise.reject(error)
    }
)
