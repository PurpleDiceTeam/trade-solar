<script setup lang="ts">
</script>

<template>
  <ul>
    <li><a href="#">Home</a></li>
    <li><a href="#">About</a></li>
    <li><a href="#">Service</a></li>
    <li><a href="#">Office</a></li>
  </ul>
</template>

<style scoped>
ul {
  list-style-type: none;
  text-align: center;
}

ul li {
  display: inline;
  padding: 16px;
  margin: 0px 16px 0px 16px;
  border-radius: 32px;
  border: 1px solid #0000;
  background:
    linear-gradient(#fafafa, #fafafa) padding-box,
    linear-gradient(140.89deg, #2755F3 0%, #1475EE 50%, #0095E9 100%) border-box;
}

ul li a {
  text-decoration: none;
  font-family: "Montserrat";
  color: #777777;
  font-display: 500;
  font-size: 18px;
  line-height: 21.6px;
}
</style>