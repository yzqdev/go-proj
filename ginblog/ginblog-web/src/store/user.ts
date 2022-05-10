import { UserState } from "@/type/storeTypes";

export const useUserStore = defineStore({
  id: "user",
  state: (): UserState => {
    return {
      token: "",
      username: "",
    };
  },
  getters: {
    getToken(): string {
      return this.token;
    },
  },
  actions: {
    setUserToken(val: string) {
      this.token = val;
    },
    setUserInfo(val: UserState) {
      this.$state = val;
    },
  },
});
