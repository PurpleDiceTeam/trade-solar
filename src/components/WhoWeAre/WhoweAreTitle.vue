<script setup lang="ts">
</script>

<template>
  <div class="tittleContainer">
    <h1>Who we are</h1>
    <img src="/src/assets/WhoWeAre/Arrow.svg" alt="Arrow" />
    <a href="#">How It's Work</a>
  </div>
</template>

<style scoped>
div {
  display: flex;
  align-items: center;
  margin-bottom: 32px;
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

div a {
  display: inline-block;
  text-align: center;
  justify-content: center;
  text-decoration: none;
  margin: 0px 0px 0px 32px;
  width: 212px;
  height: 56px;
  border: solid 1px #1F1F1F;
  border-radius: 64px;
  font-family: Montserrat Alternates;
  font-weight: 500;
  line-height: 56px;
  font-size: 20px;
  color: #fff;
  background: linear-gradient(140.89deg, #2755F3 0%, #1475EE 50%, #0095E9 100%);
}

@media only screen and (min-width: 1920px) {
  div h1 {
    font-size: 63.09px;
    line-height: 75.7px;
  }

  div img {
    width: 153px;
  }

  div a {
    width: 290.77px;
    height: 84.83px;
    font-size: 27.43px;
    line-height: 84.83px;
  }
}

@media only screen and (max-width: 1200px) {
  div {
    display: none;
  }
}
</style>