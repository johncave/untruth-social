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
import { SocialSignUp } from "@/components/partials/socials";
useSeoMeta({
  title: "Forgot Password",
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
const { isFieldDirty, handleSubmit, setFieldError } = useForm({
  validationSchema: schema,
});
const { toast } = useToast();
const onSubmit = handleSubmit((values) => {
  toast({
    title: "Form Submitted",
    description: JSON.stringify(values, null, 2),
  })
  console.log("Submitted")
});

</script>

<template>
  <frame-content>
    <auth-container>
      <auth-banner-container
        heading="Start your journey with us"
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
              <form-field
                v-slot="{ componentField }"
                name="userName"
                :validate-on-blur="!isFieldDirty"
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
