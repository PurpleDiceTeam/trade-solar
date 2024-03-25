<script setup lang="ts">
import NavBarMobile from './components/Navbar/NavBarMobile.vue'
import HeroMainMobile from './components/HeroMain/HeroMainMobile.vue'

</script>

<template>
  <div>
    <NavBarMobile />
    <HeroMainMobile />
  </div>
</template>

<style scoped>
div {
  display: flex;
  flex-direction: column;
  align-items: center;
}
</style>