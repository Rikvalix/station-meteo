import axios from "axios";

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
