import { createInertiaApp } from '@inertiajs/vue3'
import { resolvePageComponent } from "@/Plugins/vite/inertia-helpers";
import createServer from '@inertiajs/vue3/server'
import { renderToString } from '@vue/server-renderer'
import { createSSRApp, h } from 'vue'
import GuestLayout from '@/Layouts/GuestLayout.vue';
import '../css/app.css'
import VueClickAway from "vue3-click-away";

const appName = import.meta.env.VITE_APP_NAME || 'Artefak';

createServer(page =>
    createInertiaApp({
        title: (title) => `${title} - ${appName}`,
        page,
        render: renderToString,
        resolve: async (name) => {
            const page = await resolvePageComponent(`./Pages/${name}.vue`, import.meta.glob('./Pages/**/*.vue'));
            page.default.layout ??= GuestLayout;
            return page;
        },
        setup({ App, props, plugin }) {
            return createSSRApp({
                render: () => h(App, props),
            })
                .use(plugin)
                .use(VueClickAway)
        },
        progress: {
            color: '#4B5563',
        },
    }),
)