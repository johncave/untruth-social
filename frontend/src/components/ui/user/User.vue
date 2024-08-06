<script setup lang="ts">
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import { createAvatar } from '@dicebear/core';
import { adventurer } from '@dicebear/collection';
import { computed, PropType } from 'vue';
const props = defineProps({
  name: {
    type: String,
    required: true,
  },
  role: {
    type: String,
    required: true,
  },
  url: {
    type: String,
    required: true,
  },
  direction: {
    type: String as PropType<'left' | 'right'>,
    default: 'left',
  },
});

const data = createAvatar(adventurer, {
  seed: props.name,
}).toDataUri();

const fallBackName = computed(() => {
  return props.name
    .split(' ')
    .map((n) => n[0])
    .join('')
    .toUpperCase();
});
</script>
<template>
  <div class="flex items-center gap-4">
    <avatar
      :class="[
        'bg-primary text-primary-foreground',
        direction === 'left' ? 'order-1' : 'order-2',
      ]"
    >
      <avatar-image :src="data" alt="@radix-vue" />
      <avatar-fallback>{{ fallBackName }}</avatar-fallback>
    </avatar>
    <div
      :class="[
        'flex flex-col items-start',
        direction === 'left' ? 'order-2' : 'order-1',
      ]"
    >
      <span>{{ name }}</span>
    </div>
  </div>
</template>