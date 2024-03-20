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
  text-align: left;
  margin: 0px 46px 0px 0px;
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
</style>