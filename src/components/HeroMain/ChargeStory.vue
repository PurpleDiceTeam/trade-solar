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

div ul {
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
</style>