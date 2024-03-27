<script setup lang="ts">
import GetStarterButton from './GetStarterButton.vue';
import NavigateUl from './NavigateUl.vue';

</script>

<template>
	<nav class="navbar">
		<img src="/src/assets/navBar/Logo.svg" alt="Trade solar logo" />
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

@media only screen and (min-width: 1920px) {
	.navbar {
		height: 100px;
	}

	.navbar img {
		width: 336px;
		height: 44px;
	}
}
</style>