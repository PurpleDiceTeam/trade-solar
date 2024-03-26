<script setup lang="ts">
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
  width: 370px;
  min-width: 370px;
  height: 410px;
  margin: 0px 16px 0px 16px;
  padding: 16px 32px;
  box-sizing: border-box;
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
  margin-top: 20px
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

@media only screen and (min-width: 1920px) {
  article {
    min-width: 506.51px;
    height: 558.63px;
    padding: 32px;
    margin-top: 30px;
  }

  article img.iconArticle {
    width: 68.61px;
    height: 62.23px;
  }

  article h2 {
    font-size: 27.43px;
    line-height: 32.91px;
  }

  article p {
    font-size: 21.94px;
    line-height: 35.11px;
  }

  article div {
    margin-top: 50px;
  }

  article div a {
    font-size: 27.43px;
    line-height: 32.91px;
  }

  article div img {
    width: 40px;
    margin-bottom: -4px;
  }
}
</style>