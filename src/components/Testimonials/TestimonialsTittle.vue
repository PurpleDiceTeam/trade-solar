<script setup lang="ts">
</script>

<template>
  <div class="tittleContainer">
    <h1>Testimonials</h1>
    <img src="/src/assets/Testimonials/Arrow.svg" alt="Arrow" />
    <p>See what people are saying</p>
  </div>
</template>

<style scoped>
div {
  display: flex;
  align-items: center;
}

div h1 {
  margin: 0px 24px 0px 0px;
  font-family: Montserrat Alternates;
  font-size: 46px;
  font-weight: 500;
  line-height: 55px;
  letter-spacing: 0em;
  text-align: left;

}

div p {
  margin: 0px 0px 0px 32px;
  font-family: Montserrat Alternates;
  font-size: 20px;
  font-weight: 500;
  line-height: 24px;
  letter-spacing: 0em;
  text-align: left;
}

@media only screen and (min-width: 1920px) {
  div h1 {
    font-size: 63.09px;
    line-height: 75.7px;
    margin-right: 50px;
  }

  div img {
    width: 152.96px;
  }

  div p {
    font-size: 27.43px;
    line-height: 32.91px;
    margin: 0px 0px 0px 46px;
  }
}
</style>