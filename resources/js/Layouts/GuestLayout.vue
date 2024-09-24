<script setup>
import {Head, Link} from "@inertiajs/vue3";

const weekday = [
  "Sunday",
  "Monday",
  "Tuesday",
  "Wednesday",
  "Thursday",
  "Friday",
  "Saturday",
];
const now = new Date();
const today = weekday[now.getDay()];

defineProps({
  name: {
    type: String,
  },
});

</script>

<template>
  <Head>
    <meta name="description" content="my personal notes about technology" />
  </Head>

  <div class="container mx-auto flex flex-col min-h-screen font-rubik">
    <nav class="flex flex-col sm:flex-row gap-2 sm:gap-0 justify-between items-start sm:items-center uppercase text-base font-semibold
    p-4 fixed sm:relative bg-white sm:bg-transparent shadow-lg sm:shadow-none w-full z-50"
    >
      <Link
          href="/"
          class="text-white bg-primary p-2 text-xl hover:bg-link hover:-rotate-6"
      >
        Fajar SP
      </Link>
      <div class="flex gap-2 sm:gap-12">
        <Link href="/articles" class="text-primary hover:bg-link hover:text-white p-2">
          Articles
        </Link>
        <Link href="/work-with-me" class="text-primary hover:bg-link hover:text-white p-2">
          Work With Me
        </Link>
        <Link href="/about" class="text-primary hover:bg-link hover:text-white p-2">
          About
        </Link>
        <Link href="/login" class="text-primary hover:bg-link hover:text-white p-2">
          {{ name ? 'Dashboard' : 'Login' }}
        </Link>
      </div>
    </nav>
    <main class="mt-20 sm:mt-0 flex-1 h-full flex flex-col justify-start items-center">
      <slot />
    </main>
    <footer>
      <div class="flex flex-col sm:flex-row justify-between items-center text-secondary text-sm p-4">
        <div class="self-start">
          © Copyright 2024 Fajar SP<br/>Code snippets are
          <a
              href="https://opensource.org/licenses/MIT"
              class="hover:text-link hover:underline"
              target="_blank"
          >
            MIT licensed
          </a>
          <br/>
          <Link href="/disclaimer" class="text-link hover:text-link-hover hover:underline">
            Disclaimer
          </Link>
        </div>
        <div class="self-end">
          <i>
            Enjoy the rest of your <span>{{ today }}</span>!
          </i>
        </div>
      </div>
    </footer>
  </div>
</template>
