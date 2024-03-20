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
</style>