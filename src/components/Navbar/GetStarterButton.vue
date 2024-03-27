<script setup lang="ts">
</script>

<template>
  <a class="startedButton" href="#">Get Started</a>
  <img class="navArrow" src="/src/assets/navBar/navArrow.svg" alt="arrowIcon" />
</template>

<style scoped>
img {
  position: absolute;
  right: 9vw;
}

a {
  display: inline-block;
  text-align: center;
  justify-content: center;
  text-decoration: none;
  width: 181px;
  height: 56px;
  border: solid 1px #1F1F1F;
  border-radius: 64px;
  font-family: "Montserrat Alternates";
  font-weight: 500;
  line-height: 56px;
  font-size: 20px;
  color: #fff;
  background: linear-gradient(140.89deg, #2755F3 0%, #1475EE 50%, #0095E9 100%);
}

@media only screen and (min-width: 1920px) {
  img {
    position: absolute;
    right: 9vw;
    width: 47px;
    height: 47px;
  }

  a {
    width: 248px;
    height: 77px;
    font-size: 27.43px;
    line-height: 77px;
  }
}
</style>