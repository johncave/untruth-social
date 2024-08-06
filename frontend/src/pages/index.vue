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
import { ref } from 'vue';
import axios from 'axios';
import { create } from 'domain';

import { useToast } from '@/components/ui/toast';
const { toast } = useToast();

useSeoMeta({
  title: 'TRUTH',
  description: 'On Untruth Social, every post is certified false, saving you the guesswork.',
  ogTitle: 'TRUTH',
  ogDescription: 'On Untruth Social, every post is certified false, saving you the guesswork.',
});
const { setBreadcrumbs } = useBreadcrumbsStore();

setBreadcrumbs([
  { label: 'Home', to: '/dashboard' },
  { label: 'Truth', to: '/dashboard' },
]);


const fetchPosts = async () => {
  axios.get(import.meta.env.VITE_API_URL + "posts").then((res) => {
    posts.value = res.data

    console.log(typeof(posts.value), typeof(posts.value.length))

    for (let i = 0; i < posts.value.length; i++) {
      console.log("Hi")
      console.log(posts.value[i])
      posts.value[i].avatar = createAvatar(adventurer, {
        seed: posts.value[i].username
      }).toDataUri();
    }

    console.log(posts.value)
  }).catch((err) => {
    console.log(err);
    // console.log(err.response.data);

  });
}
fetchPosts();
// Poll for new posts every few seconds
setInterval(fetchPosts, 5000);




let posts = ref()       
let newPost = ref("")
let message = ref("This post will be strictly fact checked. Any facts will be removed.")
const self = this


axios.get(import.meta.env.VITE_API_URL + "suggest").then((res) => {
  newPost.value = res.data.suggestion

  console.log(newPost.value, res.data.suggestion)
}).catch((err) => {
  console.log(err);
  // console.log(err.response.data);

});

const onPost = (e) => {
  e.preventDefault()
  console.log("Submitted", e, newPost, message)
  message.value = "Fact checking..."

  axios.post(import.meta.env.VITE_API_URL + "check", {
    content: newPost.value
  }).then((res) => {
    console.log(res.data, res.data.allowed);

    if (!res.data.allowed) {
      console.log(res.data.allowed)
      console.log(message)
      message.value = "Looks like your submission is too factual. We have updated it to remove any facts."
      newPost.value = res.data.suggestion
    } else {
      axios.post(import.meta.env.VITE_API_URL + "posts", {
        content: newPost.value
      }).then((res) => {
        console.log(res.data);
        window.location.reload()
      }).catch((err) => {
        toast({
          title: "Error",
          description: err.response.data.error,
        });
        console.log(err.response.data);
      });
    }
  }).catch((err) => {
    toast({
      title: "Error",
      description: err.response.data.error,
    });
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
    <!-- <div class="space-y-6">
      <v-card>
        <v-card-header>
          <v-card-title> Untruth Social! </v-card-title>
        </v-card-header>
        <v-card-content>
          <p>
            Every day, we are bombarded with information from all sides. It's hard to know what to believe. Blah blah
            blah, this content was AI generated and I didn't check it, blah. That's why
            we created Untruth Social.
            A place where every post has been verified as untrue by our team of experts (an LLM).
          </p>
        </v-card-content>
      </v-card>
    </div> -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <div class="w-full">
        <card class="shadow-sm">
          <card-header>
            <card-description>
              <p>On Untruth Social, every post has been verified as untrue, saving you the guesswork.</p><br />Hint: Try posting something true
            </card-description>
          </card-header>
          <card-content>
            <form class="w-full space-y-6" @submit="onPost">
              <FormField v-slot="{ componentField }" name="username">
                <FormItem>
                  <FormLabel>Post</FormLabel>
                  <FormControl>
                    <!-- <Input type="textarea" placeholder="shadcn" v-bind="componentField" /> -->
                    <Textarea placeholder="Type your message here." v-model="newPost" rows="5"></Textarea>
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
