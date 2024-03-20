<script setup lang="ts">
import WhatWeOfferArticle from './WhatWeOfferArticle.vue';
</script>

<template>
  <section class="whatWeOfferSection">
    <h1>We Offer Differents Services</h1>
    <div class="articleContainer">
      <WhatWeOfferArticle title="Green Savings Consult" iconUrl="Houses.svg" articleClass=""
        text="Discover a smarter way to power your life with our exclusive consultation services. We'll guide you in minimizing waste, unlocking savings, and seamlessly transitioning into a greener, more sustainable future." />
      <WhatWeOfferArticle title="Seamless Solar Integration" iconUrl="SolarPanel.svg" articleClass="SpecialArticle"
        text="Transform your space effortlessly with our expert solar panel installation services. Our dedicated team ensures a smooth integration process, optimizing energy efficiency and liberating you from conventional power sources." />
      <WhatWeOfferArticle title="Carbon-Neutral Living" iconUrl="Hand.svg" articleClass=""
        text="Be a climate change champion with our innovative CO2e compensation programs. When you choose Trade Solar, you actively contribute to offsetting carbon emissions, achieving carbon neutrality effortlessly." />
    </div>
  </section>
</template>

<style>
.whatWeOfferSection {
  min-height: 654px;
  padding: 12px 10vw 12px 10vw;
}

.whatWeOfferSection h1 {
  margin: 30px 0px 52px 0px;
  font-family: "Montserrat Alternates", "Montserrat";
  font-size: 46px;
  font-weight: 500;
  line-height: 55px;
  letter-spacing: 0em;
  text-align: center;
}

.whatWeOfferSection div {
  display: flex;
  justify-content: space-evenly;
  align-items: center;
}
</style>