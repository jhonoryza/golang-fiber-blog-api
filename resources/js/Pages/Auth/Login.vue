<script>
import AuthLayout from "@/Layouts/AuthLayout.vue";

export default {
    layout: AuthLayout,
};
</script>

<script setup>
import InputCheckbox from "@/Components/InputCheckbox.vue";
import InputError from "@/Components/InputError.vue";
import InputLabel from "@/Components/InputLabel.vue";
import ButtonPrimary from "@/Components/ButtonPrimary.vue";
import InputText from "@/Components/InputText.vue";
import { Head, Link, useForm } from "@inertiajs/vue3";

const getCSRFToken = () => {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; csrf_=`);
    if (parts.length === 2) return parts.pop().split(";").shift();
};

const props = defineProps({
    canResetPassword: Boolean,
    errors: Object,
    message: String,
});

const form = useForm({
    email: "",
    password: "",
    remember: false,
});

const submit = () => {
    form.post("/login", {
        onFinish: () => form.reset("password"),
        headers: {
            "X-Csrf-Token": getCSRFToken(),
        },
    });
};
</script>

<template>
    <Head title="Log in" />

    <InputError class="my-2" :message="message" />

    <form @submit.prevent="submit">
        <div>
            <InputLabel for="email" value="Email" />

            <InputText
                id="email"
                type="email"
                class="mt-1 block w-full"
                v-model="form.email"
                required
                autofocus
                autocomplete="username"
            />

            <InputError class="mt-2" :message="errors?.email" />
        </div>

        <div class="mt-4">
            <InputLabel for="password" value="Password" />

            <InputText
                id="password"
                type="password"
                class="mt-1 block w-full"
                v-model="form.password"
                required
                autocomplete="current-password"
            />

            <InputError class="mt-2" :message="errors?.password" />
        </div>

        <div class="block mt-4">
            <label class="flex items-center">
                <InputCheckbox
                    name="remember"
                    v-model:checked="form.remember"
                />
                <span class="ms-2 text-sm text-gray-600 dark:text-gray-400"
                    >Remember me</span
                >
            </label>
        </div>

        <div class="flex items-center justify-end mt-4">
            <Link
                v-if="canResetPassword"
                href="/forgot-pass"
                class="underline text-sm text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-100 rounded-md focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 dark:focus:ring-offset-gray-800"
            >
                Forgot your password?
            </Link>

            <ButtonPrimary
                class="ms-4"
                :class="{ 'opacity-25': form.processing }"
                :disabled="form.processing"
            >
                Log in
            </ButtonPrimary>
        </div>
    </form>
</template>
