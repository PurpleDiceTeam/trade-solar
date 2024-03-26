<script setup lang="ts">
import PoweringUpArticle from './PoweringUpArticle.vue'
import ImageArticle from './ImageArticle.vue'
import PoweringUpStats from './PoweringUpStats.vue'
</script>

<template>
  <PoweringUpArticle />
  <ImageArticle />
  <div class="statsContainer">
    <PoweringUpStats amount="2.5k" title="Happy Clients" />
    <PoweringUpStats amount="8.4k" title="Generated" />
  </div>
</template>

<style scoped>
div.statsContainer {
  display: flex;
  margin-left: 60px;
  box-sizing: border-box;
}
</style>