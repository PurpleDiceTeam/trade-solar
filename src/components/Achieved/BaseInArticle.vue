<script setup lang="ts">
</script>

<template>
  <article class="sideArticle baseInArticle">
    <h1>Base in</h1>
    <img src="/src/assets/Achieved/AchievedWhiteArrow.svg" alt="" />
    <a href="">Germany</a>
  </article>
</template>

<style scoped>
article {
  min-width: 370px;
  max-width: 370px;
  min-height: 584px;
  max-height: 584px;
  border-radius: 27px;
  background-color: #fafafa;
  display: flex;
  justify-content: space-between;
  box-sizing: border-box;
  padding: 36px 18px 18px 36px;
  background: url("/src/assets/Achieved/AchievedSolarPanel.png");
  background-position: center center;
}

article h1 {
  margin: 0px;
  font-family: Montserrat Alternates;
  font-size: 25px;
  font-weight: 500;
  line-height: 35px;
  letter-spacing: 0.04em;
  text-align: left;
  color: #fff;
}

article a {
  display: inline-block;
  text-align: center;
  justify-content: center;
  text-decoration: none;
  min-width: 117px;
  max-width: 117px;
  min-height: 37px;
  max-height: 37px;
  font-family: Montserrat;
  font-size: 16px;
  font-weight: 500;
  line-height: 35px;
  letter-spacing: 0em;
  background-image: url("data:image/svg+xml,%3csvg width='100%25' height='100%25' xmlns='http://www.w3.org/2000/svg'%3e%3crect width='100%25' height='100%25' fill='none' rx='15' ry='15' stroke='white' stroke-width='2' stroke-dasharray='14%2c14' stroke-dashoffset='0' stroke-linecap='square'/%3e%3c/svg%3e");
  border-radius: 15px;
  color: #fff;
}

article img {
  align-self: flex-start;
}
</style>