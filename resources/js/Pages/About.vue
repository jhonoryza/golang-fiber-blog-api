<template>
  <div class="mx-auto mt-20 max-w-4xl text-primary px-4 text-base sm:text-lg">
    <h1 class="text-4xl sm:text-5xl font-bold">About this site</h1>
    <p class="mt-6">
      The website that you're reading right now is powered by a custom Go application and
      Vue application. The HTML and CSS is a custom theme which took me embarrassingly
      long to design 🙃.
    </p>
    <p class="mt-6">
      Everything is hand-written. It runs on my home server, using zero trust from
      cloudflare. Cloudflare is used to automatically manage TLS certificates and Traefik
      act as as a reverse proxy to the sites. I use Ansible to manage the server and make
      deployments.
    </p>
  </div>
</template>
