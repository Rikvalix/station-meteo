import {defineStore} from "pinia";
import type UserModel from "../model/UserModel.ts";
import {userApi} from "../api/lib/userApi.ts";
import type StationModel from "../model/StationModel.ts";
import {measureApi} from "../api/lib/measureApi.ts";

export const useUserStore = defineStore('user', {
    state: () => {
        return {
            user: {} as UserModel, // current user
            stations : [] as StationModel[],
        }
    },

    getters: {
        // Recherche l'utilisateur courant
        getCurrentUser: () => {
            if (!sessionStorage.getItem("user") ||  sessionStorage.getItem("user")?.length === 0) {
                return null
            } else {
                return {};
            }
        }
    },

    actions : {
        async loginUser(username: string, password: string) {
            try {
                const result = await userApi.loginUser(username, password);
                if (result) {
                    this.user = result;
                    sessionStorage.setItem("user", JSON.stringify(result));
                    return true
                }
            } catch {
                return null;
            }
        },
        async getAllStations() {
          try {
              const stations = await measureApi.getAllStations()
              if (stations) {
                  this.stations = stations
              }
          } catch {
              this.stations = [];
          }

        }
    }

})