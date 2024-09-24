<script>
import AdminLayout from "@/Layouts/AdminLayout.vue";

export default {
  layout: AdminLayout,
};
</script>

<script setup>
import ButtonPrimary from "@/Components/ButtonPrimary.vue";
import {Head} from "@inertiajs/vue3";

defineProps({
    name: {
        type: String,
    },
});
</script>

<template>
    <Head title="Dashboard" />

    <p>Dashboard {{ name }}</p>

</template>
