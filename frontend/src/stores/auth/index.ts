import { users } from "@/data/mock";
import { User } from "@/interfaces";
import { defineStore } from "pinia";
import { ref } from "vue";
import {roles } from "@/data/mock/roles";
import router from "@/router";
import * as bcrypt from "bcryptjs";
import { nanoid } from "nanoid";
import { useRouter } from "vue-router";
import axios from "axios";
axios.defaults.withCredentials = true;

export const useAuthStore = defineStore("auth", () => {
  const isAuth = ref(false);
  const token = ref("");
  const user = ref<User | undefined>(undefined);
  const setToken = (data: string) => (token.value = data);
  const setUser = (data: User) => (user.value = data);

  const getUser = async () => {
    console.log("BASE URL IS ", import.meta.env.VITE_API_URL);
    return axios.get(import.meta.env.VITE_API_URL+"user", {withCredentials: true});
  }

  const login = async (email: string, password: string) => {
    // const router = useRouter();
    // // fetch user data from API

    // // const user = users.find((u) => u.email === email);
    // // if (!user) return { key: "email", error: "User not found" };
    // // if (!bcrypt.compareSync(password, user.password))
    // //   return { key: "password", error: "Password is incorrect" };

    // isAuth.value = true;
    // const token = nanoid(64);
    // localStorage.setItem("token", token);
    // setToken(token);
    // setUser(user);
    // localStorage.setItem("user", JSON.stringify(user));
    // return { key: "", error: "" };
  };
  const checkAuth = () => {
    console.log("checking auth");
    getUser().then((res) => {
      console.log(res.data);
      let user: User
      user = {name: res.data.username, role: roles[0]};
      setUser(user as User);
      isAuth.value = true;
      console.log("user set", user)
    }).catch((err) => {
      console.log("Not registered?", err);
      router.push({ name: "register" });
    });
  };
  const logout = () => {
    isAuth.value = false;
    localStorage.removeItem("token");
    localStorage.removeItem("user");
    token.value = "";
    user.value = undefined;
  };
  return { isAuth, token, user, login, checkAuth, logout };
});
