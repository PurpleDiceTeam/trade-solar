<script setup lang="ts">
</script>

<template>
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
</template>

<style scoped>
div {
  width: 250px;
  height: 169px;
  position: absolute;
  box-sizing: border-box;
  align-self: self-end;
  right: 31px;
  bottom: 27.5px;
  padding: 0px 15px 0px 15px;
  background-color: #fff;
  border: 1px solid #000;
  border-radius: 18px;
}

div h1 {
  margin: 15px 0px 0px 0px;
  font-family: Montserrat Alternates;
  font-size: 16px;
  font-weight: 500;
  line-height: 16px;
  letter-spacing: 0em;
  text-align: left;
  color: #1F1F1F;
}

div p {
  margin: 10px 20px 0px 0px;
  font-family: Montserrat Alternates;
  font-size: 13px;
  font-weight: 400;
  line-height: 13px;
  letter-spacing: 0em;
  text-align: right;
  color: #888888;
}

div ul {
  width: 219px;
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

@media only screen and (min-width: 1920px) {
  div {
    width: 342.86px;
    height: 231px;
    padding: 0px 22px 0px 22px;
  }

  div h1 {
    font-size: 21.94px;
    margin: 28px 0px 0px 0px;
  }

  div p {
    font-size: 17.48px;
    margin-top: 16px;
  }

  div img {
    width: 298px;
    height: 124px;
  }

  div ul {
    width: 299px;
    font-size: 17.48px;
  }
}
</style>