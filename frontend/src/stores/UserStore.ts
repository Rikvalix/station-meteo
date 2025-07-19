import {defineStore} from "pinia";
import type UserModel from "../model/UserModel.ts";

export const userStore = defineStore('user', {
    state: () => {
        return {
            user: {} as UserModel, // current user
        }
    },

    getters: {
        // Recherche l'utilisateur courant
        getCurrentUser: () => {
            //TODO
        }
    }

})