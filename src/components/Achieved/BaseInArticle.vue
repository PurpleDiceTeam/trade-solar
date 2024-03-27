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
  width: 370px;
  height: 584px;
  border-radius: 27px;
  background-color: #fafafa;
  display: flex;
  justify-content: space-between;
  box-sizing: border-box;
  padding: 36px 18px 18px 20px;
  background: url("/src/assets/Achieved/AchievedSolarPanel.png");
  background-position: center center;
}

article h1 {
  margin: 0px;
  text-wrap: nowrap;
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
  width: 117px;
  height: 37px;
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

@media only screen and (min-width: 1920px) {
  article {
    width: 507px;
    height: 801px;
    padding: 50px 30px 30px 50px;
    background-size: cover;
  }

  article h1 {
    font-size: 34.44px;
  }

  article img {
    width: 81.41px;
  }

  article a {
    width: 161.82px;
    height: 51.05px;
    font-size: 21.94px;
    line-height: 51.05px;
    background-image: url("data:image/svg+xml,%3csvg width='100%25' height='100%25' xmlns='http://www.w3.org/2000/svg'%3e%3crect width='100%25' height='100%25' fill='none' rx='22' ry='22' stroke='white' stroke-width='3' stroke-dasharray='7%2c 16' stroke-dashoffset='42' stroke-linecap='square'/%3e%3c/svg%3e");
    border-radius: 22px;
    margin-top: -6px;
  }
}

@media only screen and (max-width: 1200px) {
  article {
    width: 327px;
    height: 276px;
    padding: 36px 12px 14px 12px;
    background-position: center 340px;
  }
}
</style>