<script>
import AdminLayout from "@/Layouts/AdminLayout.vue";

export default {
  layout: AdminLayout,
};
</script>

<script setup>
import ButtonPrimary from "@/Components/ButtonPrimary.vue";
import { useForm } from "@inertiajs/vue3";
import {getCSRFToken} from "@/Composables/helper.js";

defineProps({
    name: {
        type: String,
    },
});

</script>

<template>
    <p>Profile {{ name }}</p>

</template>
