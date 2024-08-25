<template>
  <div class="mx-auto mt-20 max-w-6xl text-primary text-base sm:text-lg px-4">
    <h1 class="text-4xl sm:text-5xl font-bold">All Articles</h1>
    <div
        class="my-4 flex flex-col sm:flex-row justify-between items-start sm:items-center"
        v-for="post in posts"
        :key="post.id"
    >
      <Link
          :href="'/articles/' + post.slug"
          class="text-link hover:cursor-pointer hover:text-link-hover hover:underline"
      >
        {{ post.title }}
      </Link>
      <p class="text-secondary">{{ post.published_at }}</p>
    </div>
  </div>
</template>
<script setup>
import { Link } from "@inertiajs/vue3";

defineProps({
  posts: {
    type: Object
  }
})
</script>
