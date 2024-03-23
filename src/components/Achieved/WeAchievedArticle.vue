<script setup lang="ts">
</script>

<template>
  <article class="sideArticle weAchievedArticle">
    <h1>We achieved</h1>
    <p class="percentage">56%</p>
    <p class="mainText">
      Green energy in
      <img src="/src/assets/Achieved/WeAchivedFocusElipse.svg" alt="dd" />
      <br />
      just the past year
    </p>
    <div class="wannaSeeHowContianer">
      <p>
        Wanna see
        <br />
        how we did it?
      </p>
      <img src="/src/assets/Achieved/AchievedBlueArrow.svg" alt="dd" />
    </div>
    <a href="">Views Details</a>
  </article>
</template>

<style scoped>
article {
  width: 370px;
  height: 584px;
  border-radius: 27px;
  background-color: #fafafa;
  box-sizing: border-box;
  padding: 48px;
}

article p img {
  position: absolute;
  margin: -5px 0px 0px -228px;
}

article h1 {
  margin: 0px;
  text-wrap: nowrap;
  font-family: Montserrat;
  font-size: 37px;
  font-weight: 500;
  line-height: 52px;
  letter-spacing: 0em;
  text-align: left;
}

article p.percentage {
  margin: 0px;
  font-family: Montserrat;
  font-size: 110px;
  font-weight: 700;
  line-height: 137px;
  letter-spacing: 0em;
  text-align: left;
}

article p.mainText {
  margin: 0px;
  font-family: Montserrat;
  font-size: 27px;
  font-weight: 400;
  line-height: 45px;
  letter-spacing: 0em;
}

article div {
  display: flex;
  margin: 80px 0px 18px 0px;
}

article div p {
  margin: 0px 50px 0px 0px;
  text-wrap: nowrap;
  font-family: Montserrat;
  font-size: 18px;
  font-weight: 500;
  line-height: 22px;
  letter-spacing: 0em;
  text-align: left;
}

article a {
  display: inline-block;
  text-align: center;
  justify-content: center;
  text-decoration: none;
  width: 166px;
  height: 53px;
  border: solid 2px #acacac;
  border-radius: 64px;
  font-family: Montserrat Alternates;
  font-size: 17px;
  font-weight: 500;
  line-height: 50px;
  letter-spacing: 0em;
  color: #323232;
}

@media only screen and (min-width: 1920px) {
  article {
    width: 507px;
    height: 801px;
    padding: 65.75px;
  }

  article h1 {
    font-size: 50.1px;
    line-height: 71.39px;
  }

  article p.percentage {
    font-size: 150.29px;
    line-height: 187.87px;
  }

  article p.mainText {
    font-size: 37.57px;
    line-height: 61.62px;
  }

  article p img {
    position: absolute;
    width: 280.49px;
    margin: -7px 0px 0px -314px;
  }

  article div {
    margin-top: 120px;
  }

  article div p {
    font-size: 25.05px;
    line-height: 30.06px;
  }

  article div img {
    width: 87.67px;
  }

  article a {
    width: 227px;
    height: 72px;
    font-family: Montserrat Alternates;
    font-size: 23.45px;
    line-height: 72px;
  }
}
</style>