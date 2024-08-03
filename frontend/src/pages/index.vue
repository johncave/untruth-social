<script setup lang="ts">
import { BlueprintAuthenticated } from '@/components/layouts/blueprint';
import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
} from '@/components/ui/card';

import {
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form'
import { Textarea } from '@/components/ui/textarea'

import { Button } from '@/components/ui/button'
import { createAvatar } from '@dicebear/core';
import { adventurer } from '@dicebear/collection';
import { useBreadcrumbsStore } from '@/stores/layout';
import { useSeoMeta } from '@unhead/vue';

import axios from 'axios';
import { create } from 'domain';

useSeoMeta({
  title: 'Dashboard',
  description: 'Officia pariatur eu minim commodo velit ut ad.',
  ogTitle: 'Dashboard',
  ogDescription: 'Officia pariatur eu minim commodo velit ut ad.',
});
const { setBreadcrumbs } = useBreadcrumbsStore();

setBreadcrumbs([
  { label: 'Home', to: '/dashboard' },
  { label: 'Dashboard', to: '/dashboard' },
]);



axios.get(import.meta.env.VITE_API_URL + "posts").then((res) => {
  posts = res.data

  for (let i = 0; i < posts.length; i++) {
    posts[i].avatar = createAvatar(adventurer, {
      seed: posts[i].username
    }).toDataUri();
  }

  console.log(posts)
}).catch((err) => {
  console.log(err.response.data);

});


let posts
let newPost:string 
let message:string = "This post will be strictly fact checked. Any facts will be removed."
const self = this

const onPost = (e) => {
  e.preventDefault()
  console.log("Submitted", e, newPost, message)
  message = "Fact checking..."

  axios.post(import.meta.env.VITE_API_URL + "check", {
    content: newPost
  }).then((res) => {
    console.log(res.data, res.data.allowed);

    if (!res.data.allowed) {
      console.log(res.data.allowed)
      console.log(message)
      message = "Looks like your submission is too factual. We have updated it to remove any facts."
      newPost = res.data.suggestion
    }
  }).catch((err) => {
    console.log(err.response.data);
  });

  // axios.post(import.meta.env.VITE_API_URL + "posts", {
  //   content: newPost
  // }).then((res) => {
  //   console.log(res.data);
  //   posts.push({
  //     username: e.target.username.value,
  //     content: newPost,
  //     description: e.target.description.value,
  //     avatar: createAvatar(adventurer, {
  //       seed: e.target.username.value
  //     }).toDataUri()
  //   })
  // }).catch((err) => {
  //   console.log(err.response.data);
  // });
}

</script>

<template>
  <blueprint-authenticated>
    <div class="space-y-6">
      <v-card>
        <v-card-header>
          <v-card-title> Untruth Social! </v-card-title>
        </v-card-header>
        <v-card-content>
          <p>
            Every day, we are bombarded with information from all sides. It's hard to know what to believe. Blah blah blah, this content was AI generated and I didn't check it, blah. That's why
            we created Untruth Social.
            A place where every post has been verified as untrue by our team of experts (an LLM).
          </p>
        </v-card-content>
      </v-card>
    </div>
    <div class="grid grid-cols-2 lg:grid-cols-3 gap-6">
      <div class="w-full">
        <card class="shadow-sm">
          <card-header>
            <card-title> New Untruth! </card-title>
            <card-description>
              Your post will be strictly fact-checked by our state of the art Large Language Model we downloaded from
              some Hugging Face repo.
            </card-description>
          </card-header>
          <card-content>
            <form class="w-full space-y-6" @submit="onPost">
              <FormField v-slot="{ componentField }" name="username">
                <FormItem>
                  <FormLabel>Post</FormLabel>
                  <FormControl>
                    <!-- <Input type="textarea" placeholder="shadcn" v-bind="componentField" /> -->
                    <Textarea placeholder="Type your message here." v-model="newPost"></Textarea>
                  </FormControl>
                  <FormDescription>
                    {{ message }}
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              </FormField>
              <Button type="submit">
                Submit
              </Button>
            </form>
          </card-content>
        </card>
      </div>
      <div class="w-full" v-for="item in posts">
        <card class="shadow-sm">
          <card-header class="relative">
            <card-title class="text-md tracking-normal text-card-foreground/60">
              <div class=" flex items-center space-x-4 rounded-md ">
                <img :src="item.avatar" width="64px" />
                <div class="flex-1">
                  <p class="">
                    @{{ item.username }}
                  </p>
                </div>

              </div>
            </card-title>

            <card-description>
              {{ item.description }}
            </card-description>
            <icon v-if="item.icon" :icon="item.icon" class="absolute top-6 right-6" width="24" height="24" />
          </card-header>
          <card-content>
            {{ item.content }}
          </card-content>
        </card>
      </div>
    </div>

  </blueprint-authenticated>
</template>

<style scoped></style>
