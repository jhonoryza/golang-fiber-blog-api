<script setup>
const startWebDevYear = 2019;
const currentYear = new Date().getFullYear();
const expYear = currentYear - startWebDevYear;
</script>

<template>
  <div>
    <div class="mx-auto mt-20 max-w-4xl text-primary text-base sm:text-lg px-4">
      <div class="flex flex-col gap-2">
        <h1 class="text-4xl sm:text-5xl font-bold">Work with me</h1>
        <p class="mt-4">
          I'm available to help you choose your tech stack, develop your application, or
          review your code.
        </p>
        <p class="mt-4">
          I have over {{ expYear }} years experience building web applications, APIs
          services using different languages and technologies. I value friendliness and
          kindness in interactions, and simplicity and clarity in code, architecture and
          workflows.
        </p>
        <p class="mt-4">
          I speak both technology and business. I'm just as happy to talk with you about
          return on investment as relational databases. In fact, where the two areas meet
          is where I like to work most.
        </p>

        <h2 class="mt-4 text-3xl font-semibold">What I can help you with</h2>
        <ol class="list-disc list-outside mt-4 flex flex-col gap-4">
          <li>
            <p>
              <strong>Planning for your project or startup</strong> — You can hire me to
              help you define the requirements for your project, pick the right technology
              stack, and architect the infrastructure. I can help create both high-level
              roadmaps and detailed technical specifications to turn your vision into
              reality.
            </p>
          </li>
          <li>
            <p>
              <strong>Software development</strong> — I can work with you to write the
              code for your web applications, APIs and RPC services. I'm proficient with
              Laravel, Nestjs or Go.
            </p>
          </li>
          <li>
            <p>
              <strong>Code review</strong> — If you'd like an experienced second set of
              eyes to go over your code then I'm happy to help.
            </p>
          </li>
        </ol>
        <h2 class="mt-4 text-3xl font-semibold">How it works</h2>
        <p class="mt-4">
          To start with, we'll schedule a video call to discuss what you want to
          accomplish and then we can move forward from there. Sessions start at one hour.
        </p>
        <p class="mt-4">
          I only take on consultancy work where I'm positive that it's a good fit for both
          you and me, and I'm confident that I'll be able to provide support in the long
          term if needed. I'm a native English speaker and wherever you are in the world,
          I'm happy to work remotely with you.
        </p>
        <p class="mt-4">
          Please feel free to email me at
          <a
              href="mailto:jardik.oryza@gmail.com"
              class="hover:underline hover:text-link-hover text-link"
              target="_blank"
          >
            jardik.oryza@gmail.com
          </a>
          to introduce yourself and start the conversation.
        </p>
      </div>
    </div>
  </div>
</template>
