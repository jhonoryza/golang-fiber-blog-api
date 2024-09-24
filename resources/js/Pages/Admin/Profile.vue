<script>
import AdminLayout from "@/Layouts/AdminLayout.vue";

export default {
  layout: AdminLayout,
};
</script>

<script setup>
import {Head} from "@inertiajs/vue3";

defineProps({
    name: {
        type: String,
    },
});

</script>

<template>
    <Head title="Profile" />
    <p>Profile {{ name }}</p>

</template>
