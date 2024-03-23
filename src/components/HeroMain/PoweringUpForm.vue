<script setup lang="ts">
</script>

<template>
  <input type="text" placeholder="Subscribehere@email.com" />
  <div>
    <a href="#" class="howItsWork">How It's Work</a>
    <a href="" class="lightUp">Light Up!</a>
  </div>
</template>

<style scoped>
div {
  min-width: 522.61px;
}

input {
  display: block;
  width: 470px;
  min-width: none;
  height: 62px;
  margin: 0px 0px 28px 0px;
  padding: 24px, 32px, 24px, 32px;
  border-radius: 64px;
  border: 2px solid #2755F3;
  font-family: Montserrat Alternates;
  box-sizing: border-box;
  padding-left: 32px;
  font-size: 20px;
  font-weight: 500;
  line-height: 24px;
  letter-spacing: 0em;
  text-align: left;
}

a.howItsWork {
  display: inline-block;
  text-align: center;
  justify-content: center;
  text-decoration: none;
  margin-right: 14px;
  width: 219px;
  border: solid 2px #1F1F1F;
  border-radius: 64px;
  font-family: "Montserrat Alternates";
  font-weight: 500;
  line-height: 62px;
  font-size: 20px;
  color: #fff;
  background: linear-gradient(140.89deg, #2755F3 0%, #1475EE 50%, #0095E9 100%);
}

a.lightUp {
  display: inline-block;
  text-align: center;
  justify-content: center;
  text-decoration: none;
  margin-left: 14px;
  width: 219px;
  border: solid 2px #acacac;
  border-radius: 64px;
  font-family: "Montserrat Alternates";
  font-weight: 500;
  line-height: 62px;
  font-size: 20px;
  color: #323232;
}

@media only screen and (min-width: 1920px) {
  input {
    width: 644.57px;
    height: 84.83px;
    font-size: 27.43px;
    line-height: 32.91px;
  }

  a.howItsWork {
    width: 300px;
    height: 85px;
    font-size: 27.43px;
    line-height: 85px;
  }

  a.lightUp {
    width: 300px;
    height: 85px;
    font-size: 27.43px;
    line-height: 85px;
  }
}
</style>