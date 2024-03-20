<script setup lang="ts">
import { defineProps } from 'vue';
interface Props {
  iconUrl: string;
  title: string;
  text: string;
  articleClass: string;
}
const props = defineProps<Props>()
</script>

<template>
  <article :class="props.articleClass">
    <img class="iconArticle" :src="'/src/assets/WhatWeOffer/' + props.iconUrl" alt="">
    <h2>{{ props.title }}</h2>
    <p>{{ props.text }}</p>
    <div>
      <a href="#">Learn More</a>
      <img src="../../assets/WhatWeOffer/ArrowWhatWeOffer.svg" alt="">
    </div>
  </article>
</template>

<style scoped>
article {
  min-width: 290px;
  max-width: 338px;
  min-height: 410px;
  max-height: 410px;
  margin: 0px 16px 0px 16px;
  padding: 32px;
  border: solid 3px #0000;
  border-radius: 26px;
  background:
    linear-gradient(#fafafa, #fafafa) padding-box,
    linear-gradient(179.9deg, rgba(39, 85, 243, 0) -5.6%, rgba(20, 117, 238, 0.32) 47.15%, #0095E9 99.91%) border-box;
}

article.SpecialArticle {
  background:
    linear-gradient(#fff, #fff) padding-box,
    linear-gradient(179.9deg, #0095E9 -3.07%, rgba(20, 117, 238, 0.32) 50.37%, rgba(39, 85, 243, 0) 103.81%) border-box;
}

article img.iconArticle {
  margin-top: 20px;
}

article h2 {
  font-family: "Montserrat Alternates", "Montserrat";
  margin: 16.6px 0px;
  font-size: 20px;
  font-weight: 500;
  line-height: 24px;
  letter-spacing: 0em;
  text-align: left;
}

article p {
  font-family: Montserrat;
  font-size: 16px;
  font-weight: 400;
  line-height: 26px;
  letter-spacing: 0em;
  text-align: left;
  text-wrap: pretty;
}

article div {
  display: flex;
  justify-content: initial;
  margin-top: 45px
}

article div a {
  font-family: Montserrat Alternates;
  font-size: 20px;
  font-weight: 500;
  line-height: 24px;
  letter-spacing: 0em;
  text-align: left;
  color: #000;
  border-bottom: 1.5px solid;
  text-decoration: none;
  margin-right: 6px;
}
</style>