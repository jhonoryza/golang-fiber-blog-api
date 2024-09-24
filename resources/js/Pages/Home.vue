<script setup>
import { Head } from '@inertiajs/vue3'

</script>

<template>
  <Head title="Home" description="Home" />

  <div>
    <div class="mx-auto mt-20 max-w-6xl px-4">
      <div class="flex flex-col xl:flex-row gap-8">
        <img
            class="self-start rounded-full w-32 sm:w-36 h-32 sm:h-36"
            src="storage/pic.jpeg"
            alt="fajar sp"
        />
        <div class="text-primary text-base sm:text-lg">
          <h1 class="text-4xl sm:text-5xl font-bold mt-8 xl:mt-0">Hey, I'm Fajar SP</h1>
          <p class="mt-8">
            I'm a full-stack web developer and author living near Bandung, Indonesia.
          </p>
          <p class="mt-6 font-semibold">
            My tech summary experience
          </p>
          <ol class="flex flex-col gap-2 mt-4 list-disc list-inside">
            <li>backend: laravel, nestjs and golang.</li>
            <li>frontend: svelte and vue.</li>
            <li>deployment: ansible and docker stuff.</li>
            <li>cloud: aws, google cloud, alibaba cloud, oracle cloud and
              linode.
            </li>
          </ol>
          <p class="mt-6">
            I also work as a freelance. I'm happy talking both technology and business,
            and can help you with application development and deployment. If you're
            interested in working together you can
            <a href="/work-with-me">find out more here</a>.
          </p>
          <p class="mt-6">
            Beyond that, I developed some open-source project in
            <a
                href="https://github.com/jhonoryza/"
                class="hover:underline hover:text-link-hover text-link"
                target="_blank"
            >
              github </a
            >, some of them laravel library, mobile and desktop app.
          </p>
          <p class="mt-6">
            If you've got any questions, suggestions, or would like to discuss a project
            then please email me or get in touch on any of the platforms below. I'd be
            happy to hear from you.
          </p>
          <div class="mt-6 flex gap-4 flex-wrap">
            <span class="no-mobile">↳</span>
            <a
                href="mailto:jardik.oryza@gmail.com"
                class="hover:underline hover:text-link-hover text-link"
                target="_blank"
            >
              jardik.oryza@gmail.com
            </a>
            <a
                href="https://github.com/jhonoryza/"
                class="hover:underline hover:text-link-hover text-link"
                target="_blank"
            >
              GitHub
            </a>
            <a
                href="https://twitter.com/jardik7"
                class="hover:underline hover:text-link-hover text-link"
                target="_blank"
            >
              Twitter
            </a>
            <a
                href="https://www.youtube.com/@labkita/"
                class="hover:underline hover:text-link-hover text-link"
                target="_blank"
            >
              Youtube
            </a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
