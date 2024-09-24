<template>
    <Head :title="message || 'Not Found'"/>
    <div class="mx-auto mt-20 max-w-4xl text-primary px-4 text-base sm:text-lg flex-1 flex flex-col items-center justify-center">
      <h1 class="text-5xl font-semibold">{{ message || 'Not Found'}}</h1>
    </div>
</template>

<script setup lang="ts">

import {Head} from "@inertiajs/vue3";

defineProps({
  message: {
    type: String,
  }
})

</script>