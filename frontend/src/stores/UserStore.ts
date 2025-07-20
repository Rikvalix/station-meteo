import {defineStore} from "pinia";
import type UserModel from "../model/UserModel.ts";
import {userApi} from "../api/lib/userApi.ts";

export const useUserStore = defineStore('user', {
    state: () => {
        return {
            user: {} as UserModel, // current user
        }
    },

    getters: {
        // Recherche l'utilisateur courant
        getCurrentUser: () => {
            if (!localStorage.getItem("userToken")) {
                return null
            }
        }
    },

    actions : {
        async loginUser(username: string, password: string) {
            try {
                const result = await userApi.loginUser(username, password);
                if (result) {
                    // Set le token
                    return result
                }
            } catch {
                return null;
            }
        }
    }

})