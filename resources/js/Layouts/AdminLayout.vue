<script setup>
import {Link, useForm} from "@inertiajs/vue3";
import {ref} from "vue";
import {getCSRFToken} from "@/Composables/helper.js";
import {IconUser} from "@tabler/icons-vue";

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
const showDropdownUser = ref(false);
const form = useForm({});

const submit = () => {
  form.post("/auth/logout", {
    headers: {
      "X-Csrf-Token": getCSRFToken(),
    },
  });
};

defineProps({
  name: String,
})
</script>

<template>
  <div class="container mx-auto flex flex-col min-h-screen font-rubik">
    <nav class="flex flex-col sm:flex-row gap-2 sm:gap-0 justify-between items-start sm:items-center uppercase text-base font-semibold
    p-4 fixed sm:relative bg-white sm:bg-transparent shadow-lg sm:shadow-none w-full z-20"
    >
      <Link
          href="/"
          class="text-white bg-primary p-2 text-xl hover:bg-link hover:-rotate-6"
      >
        Fajar SP
      </Link>
      <div class="flex gap-2 sm:gap-12">
        <Link href="/auth/dashboard" class="text-primary hover:bg-link hover:text-white p-2">
          Dashboard
        </Link>
        <div class="relative">
          <button class="flex gap-2 items-center text-primary hover:cursor-pointer p-2 uppercase hover:bg-link hover:text-white" @click="() => showDropdownUser = !showDropdownUser">
            {{ name }}
            <IconUser class="size-4" />
          </button>
          <div class="z-30 bg-white absolute top-12 right-0 w-48 flex flex-col gap-2 rounded-lg shadow-lg border border-gray-200 overflow-hidden" v-show="showDropdownUser">
            <Link href="/auth/profile" class="text-primary p-2 hover:bg-link hover:text-white">
              Profile
            </Link>
            <button type="button" @click="submit" class="text-primary p-2 uppercase inline-flex hover:bg-link hover:text-white">
              Logout
            </button>
          </div>
        </div>
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
