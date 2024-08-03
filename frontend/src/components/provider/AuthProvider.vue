<script setup lang="ts">
import { useAuthStore } from "@/stores";
import { watchEffect } from "vue";
import { useRouter } from "vue-router";

const auth = useAuthStore();
const router = useRouter();
watchEffect(() => {
  console.log("Triggering check auth");
  auth.checkAuth();
  console.log("Authed", auth.isAuth);
});

// // watch for auth changes
// watchEffect(() => {
//   console.log("Authed", auth.isAuth);
//   //auth.checkAuth();
//   if (!auth.isAuth) {
//     router.push({ name: "register" });
//   } else {
//     router.push({ name: "index" });
//   }
// });
</script>

<template>
  <div v-if="auth.isAuth">
    <slot />
  </div>
  <div v-else></div>
</template>
