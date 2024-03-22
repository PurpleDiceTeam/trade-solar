<script setup lang="ts">
</script>

<template>
  <section class="clientsSection">
    <link rel="stylesheet" href="/css/clients.css" />
    <img src="../../assets/ClientsSection/Samsung.svg" alt="Samsung logo" />
    <img src="../../assets/ClientsSection/Apple.svg" alt="Apple logo" />
    <img src="../../assets/ClientsSection/Patagonia.svg" alt="Patagonia logo" />
    <img src="../../assets/ClientsSection/Walmart.svg" alt="Walmart logo" />
    <img src="../../assets/ClientsSection/Spotify.svg" alt="Spotify logo" />
    <img src="../../assets/ClientsSection/Etsy.svg" alt="Etsy logo" />
  </section>
</template>

<style scoped>
.clientsSection {
  height: 154px;
  display: flex;
  padding: 0px 10vw 0px 10vw;
  justify-content: space-between;
  align-items: center;
  background-color: #fafafa;
  box-shadow:
    4px 0px 8px 0px #0000001A inset,
    15px 0px 15px 0px #00000017 inset,
    34px 0px 20px 0px #0000000D inset,
    61px 0px 24px 0px #00000003 inset,
    95px 0px 27px 0px #00000000 inset;
}

@media only screen and (min-width: 1920px) {
  .clientsSection {
    height: 211px;
  }

  .clientsSection img {
    height: 20%;
  }
}
</style>