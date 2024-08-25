<template>
  <div v-html="renderedMarkdown" />
</template>

<script setup>
import MarkdownIt from "markdown-it";
import Shiki from '@shikijs/markdown-it'
import {computed, ref} from "vue";

const md = ref(MarkdownIt());

// Async function to initialize Shiki
async function initializeShiki() {
  md.value.use(await Shiki({
    themes: {
      light: 'one-dark-pro',
      dark: 'one-dark-pro',
    }
  }));

  // Override the default rendering of tables to add a wrapper div
  const defaultRender = md.value.renderer.rules.table_open || function (tokens, idx, options, env, self) {
    return self.renderToken(tokens, idx, options);
  };

  md.value.renderer.rules.table_open = function (tokens, idx, options, env, self) {
    // Add wrapper div with classes before table
    return '<div class="overflow-x-scroll scroll-smooth shadow-md p-2 border border-gray-200 rounded-md">' + defaultRender(tokens, idx, options, env, self);
  };

  md.value.renderer.rules.table_close = function (tokens, idx, options, env, self) {
    // Close the wrapper div after the table
    return defaultRender(tokens, idx, options, env, self) + '</div>';
  };
}

// Call the initialization function
initializeShiki();

// import MarkdownItAbbr from "markdown-it-abbr";
// import MarkdownItAnchor from "markdown-it-anchor";
// import MarkdownItFootnote from "markdown-it-footnote";
// import MarkdownItHighlightjs from "markdown-it-highlightjs";
// import MarkdownItSub from "markdown-it-sub";
// import MarkdownItSup from "markdown-it-sup";
// import MarkdownItTasklists from "markdown-it-task-lists";
// import MarkdownItTOC from "markdown-it-toc-done-right";
//
// const markdown = new MarkdownIt()
//     .use(MarkdownItAbbr)
//     .use(MarkdownItAnchor)
//     .use(MarkdownItFootnote)
//     .use(MarkdownItHighlightjs)
//     .use(MarkdownItSub)
//     .use(MarkdownItSup)
//     .use(MarkdownItTasklists)
//     .use(MarkdownItTOC);

const props = defineProps({
  source: {
    type: String,
    default: ""
  }
});

// Computed property to render Markdown content to HTML
const renderedMarkdown = computed(() => md.value.render(props.source));
</script>