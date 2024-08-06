<script setup lang="ts">
import { useAuthStore, useNavigationMenu, useThemeStore } from '@/stores';
import { Button } from '@/components/ui/button';
import { User } from '@/components/ui/user';

import { BrandLogo } from '@/components/ui/brand';
import { computed, ref } from 'vue';
import { useBreadcrumbsStore } from '@/stores/layout';
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbSeparator,
} from '@/components/ui/breadcrumb';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';
import {
  Dialog,
  DialogHeader,
  DialogFooter,
  DialogTitle,
  DialogDescription,
  DialogContent,
  DialogClose,
} from '@/components/ui/dialog';

import { useRouter } from 'vue-router';
const theme = useThemeStore();
const navMenu = useNavigationMenu();
const { breadcrumbs } = useBreadcrumbsStore();
const breadcrumbsLength = computed(() => breadcrumbs.length);
const breadcrumbsList = computed(() => {
  const lists = [...breadcrumbs];
  if (breadcrumbsLength.value > 3) {
    lists.splice(1, breadcrumbsLength.value - 3, {
      label: '...',
    });
  }
  return lists;
});
const showDialogLogout = ref(false);
const auth = useAuthStore();
const quit = () => {
  showDialogLogout.value = false;
  auth.logout();
};
const router = useRouter();
</script>

<template>
  <div
    class="w-full px-6 h-20 flex items-center justify-between sticky top-0 bg-background/10 backdrop-blur-md z-10"
  >
    <div class="w-full">
      <BrandLogo />
    </div>
    <div class="w-full flex items-center justify-between">
      <div class="w-full hidden lg:block">
        <breadcrumb>
          <breadcrumb-list>
            <template
              v-for="(breadcrumb, index) in breadcrumbsList"
              :key="breadcrumb.label"
            >
              <breadcrumb-item>
                <breadcrumb-link :to="breadcrumb.to">
                  {{ breadcrumb.label }}
                </breadcrumb-link>
                <Breadcrumb-separator v-if="index !== breadcrumbsLength - 1" />
              </breadcrumb-item>
            </template>
          </breadcrumb-list>
        </breadcrumb>
      </div>
      <div class="w-full flex items-center justify-end gap-4">
        
        <toggle-theme />
        
            <user
              :name="auth.user?.name!"
              :role="auth.user?.role.name!"
              url=""
              direction="left"
            />
          
        
      </div>
    </div>
    <Dialog :open="showDialogLogout">
      <dialog-content
        class="z-50 p-10"
        :class="[theme.darkMode && 'dark', 'bg-card text-card-foreground']"
        v-bind="$attrs"
      >
        <dialog-header>
          <dialog-title>Logout</dialog-title>
          <dialog-description>
            Are you sure you want to logout?
          </dialog-description>
        </dialog-header>
        <dialog-footer>
          <div class="w-full flex flex-col lg:flex-row gap-4 justify-end">
            <v-button @click="quit" variant="destructive" class="h-12 w-full"
              >Yes</v-button
            >
            <v-button
              @click="showDialogLogout = false"
              variant="secondary"
              class="h-12 w-full"
            >
              No
            </v-button>
          </div>
        </dialog-footer>
      </dialog-content>
    </Dialog>
  </div>
</template>
