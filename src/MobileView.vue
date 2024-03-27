<script setup lang="ts">
import NavBarMobile from './components/Navbar/NavBarMobile.vue'
import HeroMainMobile from './components/HeroMain/HeroMainMobile.vue'
import WhatWeOffer from './components/WhatWeOffer/WhatWeOffer.vue'
import WhoWeAre from './components/WhoWeAre/WhoWeAre.vue'
import AchievedSection from './components/Achieved/AchievedSection.vue'
import TestimonialsSection from './components/Testimonials/TestimonialsSection.vue'
</script>

<template>
  <div>
    <NavBarMobile />
    <HeroMainMobile />
    <WhatWeOffer />
    <WhoWeAre />
    <AchievedSection />
    <TestimonialsSection />
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