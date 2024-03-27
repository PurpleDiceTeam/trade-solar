<script setup lang="ts">
import NavBarMobile from './components/Navbar/NavBarMobile.vue'
import HeroMainMobile from './components/HeroMain/HeroMainMobile.vue'
import WhatWeOffer from './components/WhatWeOffer/WhatWeOffer.vue'
import WhoWeAre from './components/WhoWeAre/WhoWeAre.vue'
</script>

<template>
  <div>
    <NavBarMobile />
    <HeroMainMobile />
    <WhatWeOffer />
    <WhoWeAre />
  </div>
</template>

<style scoped>
div {
  width: 327px;
  display: flex;
  flex-direction: column;
  margin: auto;
  align-items: center;
}
</style>