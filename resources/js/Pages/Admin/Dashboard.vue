<script setup>
import ButtonPrimary from "@/Components/ButtonPrimary.vue";
import { useForm } from "@inertiajs/vue3";

defineProps({
    name: {
        type: String,
    },
});

const form = useForm({});

const getCSRFToken = () => {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; csrf_=`);
    if (parts.length === 2) return parts.pop().split(";").shift();
};

const submit = () => {
    form.post("/auth/logout", {
        headers: {
            "X-Csrf-Token": getCSRFToken(),
        },
    });
};
</script>

<template>
    <h1>Dashboard</h1>

    <p>Welcome {{ name }}</p>

    <form @submit.prevent="submit">
        <ButtonPrimary
            class="ms-4"
            :class="{ 'opacity-25': form.processing }"
            :disabled="form.processing"
        >
            Log out
        </ButtonPrimary>
    </form>
</template>
