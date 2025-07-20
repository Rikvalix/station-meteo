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
            if (!localStorage.getItem("user")) {
                return null
            }
        }
    },

    actions : {
        async loginUser(username: string, password: string) {
            try {
                const result = await userApi.loginUser(username, password);
                if (result) {
                    this.user = result;
                    localStorage.setItem("user", JSON.stringify(result));
                    return true
                }
            } catch {
                return null;
            }
        }
    }

})