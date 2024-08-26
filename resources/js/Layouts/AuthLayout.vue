<script setup>
import ButtonPrimary from "@/Components/ButtonPrimary.vue";
import ButtonDarkMode from "@/Components/ButtonDarkMode.vue";

const back = () => {
    window.location = "/";
};
</script>

<template>
    <div
        class="min-h-screen flex flex-col sm:justify-center items-center pt-6 sm:pt-0 bg-gray-100 dark:bg-gray-900"
    >
        <div class="flex gap-2 justify-start items-center">
            <p class="font-bold text-2xl text-gray-700 dark:text-gray-300">
                Login
            </p>
            <ButtonDarkMode />
        </div>

        <div
            class="w-full sm:max-w-md mt-6 px-6 py-4 bg-white dark:bg-gray-800 shadow-md overflow-hidden sm:rounded-lg"
        >
            <slot />
        </div>
        <ButtonPrimary @click="back()" class="mt-6"> Back </ButtonPrimary>
    </div>
</template>
