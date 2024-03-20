<script setup lang="ts">
import GetStarterButton from './GetStarterButton.vue';
import NavigateUl from './NavigateUl.vue';

</script>

<template>
	<nav class="navbar">
		<img src="/src/assets/NavBar/Logo.svg" alt="Trade solar logo" />
		<NavigateUl />
		<GetStarterButton />
	</nav>
</template>

<style scoped>
.navbar {
	height: 80px;
	padding: 12px 10vw 12px 10vw;
	display: flex;
	justify-content: space-between;
	align-items: center;
}
</style>