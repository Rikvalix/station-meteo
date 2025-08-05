import LoginView from "../views/LoginView.vue";
import {createRouter, createWebHistory} from "vue-router";
import HomeView from "../views/HomeView.vue";
import ProfileView from "../views/ProfileView.vue";
import StationView from "../views/StationView.vue";

const routes = [
    {path: "/", component: HomeView},
    {path: "/home", component: HomeView},
    {path: '/login', component: LoginView},
    {path: '/profil', component: ProfileView},
    {path: '/stations', component: StationView}
]


const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router