<script setup lang="ts">
interface Props {
  amount: string;
  title: string;
}
const props = defineProps<Props>()
</script>

<template>
  <div>
    <h2>{{ props.amount }}</h2>
    <p>{{ props.title }}</p>
  </div>
</template>

<style scoped>
div {
  display: flex;
  flex-direction: column;
  margin: 0px 46px 0px 0px;
  text-wrap: nowrap;
}

h2 {
  font-family: "Montserrat Alternates";
  margin: 78px 0px 16px 0px;
  font-size: 60px;
  font-weight: 500;
  line-height: 72px;
  letter-spacing: 0em;
  text-align: left;
}

div p {
  font-family: "Montserrat Alternates";
  margin: 0px;
  font-size: 24px;
  font-weight: 500;
  line-height: 29px;
  letter-spacing: 0em;
  text-align: left;
}

@media only screen and (min-width: 1920px) {
  h2 {
    margin-top: 160px;
    font-size: 82.29px;
  }

  div p {
    font-size: 32.91px;
  }
}

@media only screen and (max-width: 1200px) {
  h2 {
    margin: 14px 0px 0px 0px;
    font-size: 45.854px;
    line-height: 55.025px;
  }

  div p {
    font-size: 18.342px;
    line-height: 22.01px;
  }
}
</style>