<script lang="ts" setup>
import {
  AuthContainer,
  AuthBannerContainer,
  AuthFormContainer,
} from "@/components/partials/auth";
import { useSeoMeta } from "@unhead/vue";
import ImgSvg from "@/assets/svg/register.svg";
import { toTypedSchema } from "@vee-validate/zod";
import { FrameContent } from "@/components/layouts/frame";
import {
  Form,
  FormField,
  FormControl,
  FormItem,
  FormLabel,
  FormMessage,
  FormDescription,
} from "@/components/ui/form";
import { Toaster, useToast } from "@/components/ui/toast";
import * as z from "zod";
import { useForm } from "vee-validate";
import { useRouter } from "vue-router";
import { createAvatar } from '@dicebear/core';
import { adventurer } from '@dicebear/collection';
import { watchEffect, ref  } from 'vue';
const router = useRouter();
import axios from "axios";
axios.defaults.withCredentials = true

useSeoMeta({
  title: "Register User",
  description: "Forgot your password? No worries, we got you covered.",
  ogTitle: "Forgot Password",
  ogDescription: "Forgot your password? No worries, we got you covered.",
});
const schema = toTypedSchema(
  z
    .object({
      userName: z.string(),
     
    })
);

var username = ref('');
const { isFieldDirty, handleSubmit, setFieldError } = useForm({
  validationSchema: schema,
});
const { toast } = useToast();
const onSubmit = handleSubmit((values) => {
  toast({
    title: "Registering...",
    description: JSON.stringify(values, null, 2),
  })
  console.log("Submitted")
  axios.post(import.meta.env.VITE_API_URL+"register", values, {withCredentials: true}).then((res) => {
    console.log(res.data);
    toast({
      title: "Registered",
      description: "You have successfully registered",
    });
    router.push({ name: "index" });
  }).catch((err) => {
    console.log(err.response.data);
    toast({
      title: "Error",
      description: err.response.data.error,
    });
  });
});

var avatar = ref('')

watchEffect(() =>{
  //  console.log("Value is", username.value)
  avatar =   createAvatar(adventurer,{
    seed: username.value
  }).toDataUri();

})


</script>

<template>
  <frame-content>
    <auth-container>
      <auth-banner-container
        heading="Start speaking the untruth"
        subheading="Create an account to get started"
      >
        <template #image>
          <img-svg class="w-full" />
        </template>
      </auth-banner-container>
      <auth-form-container>
        <div class="w-full space-y-8">
          <form @submit="onSubmit" class="w-full space-y-8">
            <div class="w-full flex items-center gap-6">
                <div class="flex items-center rounded border">
                  <img :src="avatar" width="96px" />
                </div>

              <form-field
                v-slot="{ componentField }"
                name="userName"
                :validate-on-blur="!isFieldDirty"
                v-model="username"
              >
                <form-item class="w-full">
                  <form-label>Username</form-label>
                  <form-control>
                    <v-input
                      v-bind="componentField"
                      class="w-full h-14"
                      placeholder="realtruth"
                    />
                  </form-control>
                  <form-message />
                </form-item>
              </form-field>
              
            </div>

           
          
           
            <v-button type="submit" class="w-full h-14"> Submit </v-button>
            <!-- <social-sign-up /> -->
          </form>
        </div>
      </auth-form-container>
    </auth-container>
    <toaster />
  </frame-content>
</template>
