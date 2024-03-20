<script setup lang="ts">
import BaseInArticle from './BaseInArticle.vue';
import CenterArticles from './CenterArticles.vue';
import WeAchievedArticle from './WeAchievedArticle.vue';
</script>

<template>
  <section class="achieved">
    <WeAchievedArticle />
    <div>
      <CenterArticles id="clienteleValuations" title-first-world="Clientele" title-second-world="valuations"
        amount="+1.5b" />
      <CenterArticles id="inFunding" title-first-world="In" title-second-world="funding" amount="+500m" />
    </div>
    <BaseInArticle />
  </section>
</template>

<style scoped>
.achieved {
  min-height: 664px;
  padding: 0px 10vw 0px 10vw;
  display: flex;
  justify-content: space-between;
}
</style>