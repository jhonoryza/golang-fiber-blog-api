import { createApp, h } from 'vue';
import { createInertiaApp } from '@inertiajs/vue3';
import { resolvePageComponent } from '@/Plugins/vite/inertia-helpers';
import GuestLayout from '@/Layouts/GuestLayout.vue';
import '../css/app.css'
import VueClickAway from "vue3-click-away";

const appName = import.meta.env.VITE_APP_NAME || 'Artefak';

createInertiaApp({
    title: (title) => `${title} - ${appName}`,
    resolve: async (name) => {
        const page = await resolvePageComponent(`./Pages/${name}.vue`, import.meta.glob('./Pages/**/*.vue'));
        page.default.layout ??= GuestLayout;
        return page;
    },
    setup({ el, App, props, plugin }) {
        return createApp({ render: () => h(App, props) })
            .use(plugin)
            .use(VueClickAway)
            .mount(el);
    },
    progress: {
        color: '#4B5563',
    },
});
