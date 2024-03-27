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
  height: 543px;
  background-color: #fff;
  padding: 0px 10vw 0px 10vw;
}

.whoWeAre div.articlesContainer {
  display: flex;
  justify-content: space-around;
  align-items: center;
}

@media only screen and (min-width: 1920px) {
  .whoWeAre {
    height: 744.49px;
  }
}

@media only screen and (max-width: 1200px) {
  .whoWeAre {
    width: 327px;
    box-sizing: border-box;
    height: 904px;
  }

  .whoWeAre div.articlesContainer {
    margin-top: 14px;
    display: flex;
    flex-direction: column;
  }
}
</style>