<script setup lang="ts">
import ChargeStory from './ChargeStory.vue';
import SmartEnergy from './SmartEnergy.vue';

</script>

<template>
  <article class="imageArticle">
    <img src="../../assets/HeroMain/SolarPanelHeroImage.png" alt="" />
    <SmartEnergy />
    <ChargeStory />
  </article>
</template>

<style scoped>
article {
  display: flex;
  position: relative;
}

@media only screen and (min-width: 1920px) {
  article img {
    width: 787.2px;
    height: 821.49px;
  }
}
</style>