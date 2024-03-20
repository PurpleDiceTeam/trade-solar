<script setup lang="ts">
</script>

<template>
  <article class="imageArticle">
    <img src="../../assets/HeroMain/SolarPanelHeroImage.png" alt="" />
    <div class="smartEnergy">
      <img src="../../assets/HeroMain/PlantHero.svg" alt="" />
      <h1>Smart Energy</h1>
      <p>
        For a better
        <span>FUTURE </span>
      </p>
    </div>
    <div class="chargeStory">
      <h1>Charge Story</h1>
      <p>+350kwh</p>
      <img src="../../assets/HeroMain/ChargeStory.svg" alt="" />
      <ul>
        <li>Jan</li>
        <li>Feb</li>
        <li>Mar</li>
        <li>Apr</li>
        <li>May</li>
        <li>June</li>
      </ul>
    </div>
  </article>
</template>

<style scoped>
article {
  display: flex;
  position: relative;
}

article div.chargeStory {
  position: absolute;
  box-sizing: border-box;
  align-self: self-end;
  right: 31px;
  bottom: 27.5px;
  padding: 0px 15px 0px 15px;
  min-width: 250px;
  max-width: 250px;
  min-height: 169px;
  max-height: 169px;
  background-color: #fff;
  border: 1px solid #000;
  border-radius: 18px;
}

article div.chargeStory ul {
  max-width: 219px;
  display: flex;
  list-style-type: none;
  justify-content: space-between;
  padding-left: 0;
  margin: -15px 0px 0px 0px;
  font-family: Montserrat Alternates;
  font-size: 13px;
  font-weight: 400;
  line-height: 13px;
  letter-spacing: 0em;
  color: #888888;
}

article div.chargeStory h1 {
  margin: 15px 0px 0px 0px;
  font-family: Montserrat Alternates;
  font-size: 16px;
  font-weight: 500;
  line-height: 16px;
  letter-spacing: 0em;
  text-align: left;
  color: #1F1F1F;
}

article div.chargeStory p {
  margin: 10px 20px 0px 0px;
  font-family: Montserrat Alternates;
  font-size: 13px;
  font-weight: 400;
  line-height: 13px;
  letter-spacing: 0em;
  text-align: right;
  color: #888888;
}

article div.smartEnergy {
  position: absolute;
  box-sizing: border-box;
  left: -50px;
  bottom: 121px;
  padding: 10px;
  min-width: 157px;
  max-width: 157px;
  min-height: 193px;
  max-height: 193px;
  background-color: #fff;
  border-radius: 22px;
  border: 1px solid black;
}

article div.smartEnergy h1 {
  margin: 12px 0px 5px;
  font-family: Montserrat Alternates;
  margin-bottom: 5px;
  font-size: 18px;
  font-weight: 500;
  line-height: 18px;
  letter-spacing: 0em;
  text-align: left;
}

article div.smartEnergy p {
  margin: 0px;
  font-family: Montserrat;
  font-size: 12px;
  font-weight: 400;
  line-height: 17px;
  letter-spacing: 0em;
  text-align: left;

}

article div.smartEnergy p span {
  font-family: Montserrat;
  font-size: 12px;
  font-weight: 600;
  line-height: 17px;
  letter-spacing: 0em;
  text-align: left;
  color: #2755F3;
}
</style>