<script setup lang="ts">
</script>

<template>
  <article class="getFreeConsult">
    <p>SOLAR</p>
    <a href="#">GET A FREE CONSULT</a>
  </article>
</template>

<style scoped>
article {
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 20px 0px 0px -11px;
  width: 782px;
}

article p {
  position: absolute;
  display: inline-block;
  margin: 0px;
  color: #E3F2FF;
  font-family: Montserrat;
  font-size: 224px;
  font-weight: 400;
  line-height: 269px;
  letter-spacing: 0em;
  text-align: left;
  z-index: 0;
}

article a {
  font-family: Montserrat;
  font-size: 24px;
  font-weight: 700;
  line-height: 36px;
  letter-spacing: 0.03em;
  text-align: left;
  background: -webkit-linear-gradient(140.89deg, #2755F3 0%, #1475EE 50%, #0095E9 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  z-index: 1;
}

@media only screen and (min-width: 1920px) {
  article p {
    font-size: 307.65px;
    line-height: 369.18px;
    margin: 20px 0px 0px 280px;
  }

  article a {
    text-decoration: underline;
    font-size: 32.91px;
    font-weight: 700;
    line-height: 49.37px;
    margin-left: 230px;
  }
}

@media only screen and (max-width: 1200px) {
  article {
    width: 327px;
    height: 40px;
  }

  article p {
    margin-bottom: -1010px;
    line-height: 0px;
    margin-left: 20px;
    font-size: 96px;
  }

  article a {
    margin: 48px 0px;
  }
}
</style>