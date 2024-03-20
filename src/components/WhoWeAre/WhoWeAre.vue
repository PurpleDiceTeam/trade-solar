<script setup lang="ts">
import GoalsArticle from './GoalsArticle.vue';
import MoreAboutArticle from './MoreAboutArticle.vue';
import WhoweAreTitle from './WhoweAreTitle.vue';
</script>

<template>
  <section class="whoWeAre">
    <WhoweAreTitle />
    <div class="articlesContainer">
      <GoalsArticle />
      <MoreAboutArticle />
    </div>
  </section>
</template>

<style scoped>
.whoWeAre {
  min-height: 543px;
  background-color: #fff;
  padding: 0px 10vw 0px 10vw;
}

.whoWeAre div.articlesContainer {
  display: flex;
  justify-content: space-around;
  align-items: center;
}
</style>