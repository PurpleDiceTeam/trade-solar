<script setup lang="ts">
</script>

<template>
  <div class="smartEnergy">
    <img src="../../assets/HeroMain/PlantHero.svg" alt="" />
    <h1>Smart Energy</h1>
    <p>
      For a better
      <span>FUTURE </span>
    </p>
  </div>
</template>

<style scoped>
div {
  position: absolute;
  box-sizing: border-box;
  left: -50px;
  bottom: 121px;
  padding: 10px;
  width: 157px;
  height: 193px;
  background-color: #fff;
  border-radius: 22px;
  border: 1px solid black;
}

div h1 {
  margin: 12px 0px 5px;
  font-family: Montserrat Alternates;
  margin-bottom: 5px;
  font-size: 18px;
  font-weight: 500;
  line-height: 18px;
  letter-spacing: 0em;
  text-align: left;
}

div p {
  margin: 0px;
  font-family: Montserrat;
  font-size: 12px;
  font-weight: 400;
  line-height: 17px;
  letter-spacing: 0em;
  text-align: left;

}

div p span {
  font-family: Montserrat;
  font-size: 12px;
  font-weight: 600;
  line-height: 17px;
  letter-spacing: 0em;
  text-align: left;
  color: #2755F3;
}

@media only screen and (min-width: 1920px) {
  div {
    width: 215.31px;
    height: 264.69px;
    padding: 14.4px;
  }

  div img {
    width: 186.51px;
    height: 146.74px;
  }

  div h1 {
    font-size: 24.69px;
    margin: 22px 0px 14px;
  }

  div p {
    font-size: 16.92px;
  }

  div p span {
    font-size: 16.92px;
  }
}
</style>