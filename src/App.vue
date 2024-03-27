<script setup lang="ts">
import DesktopView from './DesktopView.vue';
import MobileView from './MobileView.vue';
import { ref, onMounted, onBeforeUnmount } from 'vue';

const isMobile = ref(window.outerWidth <= 1200);

const handleResize = () => {
  isMobile.value = window.outerWidth <= 1200;
};

onMounted(() => {
  window.addEventListener('resize', handleResize);
});

onBeforeUnmount(() => {
  window.removeEventListener('resize', handleResize);
});
</script>

<template>
  <MobileView v-if="isMobile" />
  <DesktopView v-else />
</template>

<style scoped></style>
