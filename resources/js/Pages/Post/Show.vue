<template>
  <div class="mx-auto my-20 max-w-6xl text-primary text-base sm:text-lg px-4">
    <Link href="/articles" class="p-2 bg-primary text-white rounded-lg hover:bg-link">
      ⬅ Go Back
    </Link>
    <article v-if="post" class="flex flex-col gap-1 my-12">
      <img class="w-48 h-48 object-cover rounded-md" :src="post.image_url" alt="post image"/>
      <h1 class="mt-4 text-4xl sm:text-5xl font-bold">{{ post.title }}</h1>
      <span class="mt-4 text-base text-secondary">
        Published on: {{ post.published_at }}
      </span>
      <span class="mt-0 text-base">Categories: {{ post.categories_name }}</span>
      <span class="mt-0 text-base">Author: {{ post.author_name }}</span>
      <div class="mt-4 overflow-x-auto">
      <MarkdownRender class="prose lg:prose-xl" :source="post.content" />
      </div>
    </article>
    <div class="flex justify-between">
      <Link
          href="/articles"
          class="p-2 bg-primary text-white rounded-lg hover:bg-link"
      >
        ⬅ Go Back
      </Link>
      <button
          @click="scrollToTop"
          class="px-4 py-2 text-white rounded-full bg-primary hover:bg-gray-400"
      >
        ⬆
      </button>
    </div>
  </div>
</template>

<script setup>
import { Link } from "@inertiajs/vue3";
import MarkdownRender from "@/Components/MarkdownRender.vue";

defineProps({
  post: {
    type: Object
  }
})

const scrollToTop = () => {
  window.scrollTo({top: 0, behavior: "smooth"});
};
</script>
