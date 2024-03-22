<script setup lang="ts">
</script>

<template>
  <article class="goalsArticle">
    <h2>200k+</h2>
    <p>PROJECTS COMPLETED</p>
    <h2>150k+</h2>
    <p>HAPPY COSTUMERS</p>
    <h2>500+</h2>
    <p>AWARDS ACHIEVEMENT</p>
  </article>
</template>

<style scoped>
article {
  height: 417px;
  width: 387px;
  margin: 0px 40px;
  box-sizing: border-box;
  padding-top: 27px;
  border-radius: 32px;
  text-align: center;
  background: linear-gradient(180deg, rgba(39, 85, 243, 0.5) 100%, rgba(32, 68, 190, 0.253928) 100%, rgba(25, 50, 135, 0) 100%, rgba(39, 85, 243, 0) 100%), url('/src/assets/WhoWeAre/WhoWeAreSolarPanel.png');
  background-blend-mode: overlay;
  background-position: -80px -210px;
}

article h2 {
  font-family: Montserrat Alternates;
  font-size: 64px;
  font-weight: 500;
  line-height: 90px;
  letter-spacing: 0em;
  margin: 0px;
  color: #fff;
}

article p {
  font-family: Montserrat;
  font-size: 16px;
  font-weight: 500;
  line-height: 22px;
  letter-spacing: 0em;
  margin: 0px;
  color: #fff;
}

@media only screen and (min-width: 1920px) {
  article {
    width: 515.66px;
    height: 572px;
    padding-top: 57.6px;
    background-position: -80px -100px;
  }

  article h2 {
    font-size: 87.77px;
  }

  article p {
    font-size: 21.94px;
    margin-bottom: 65px;
  }
}
</style>