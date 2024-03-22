<script setup lang="ts">
import ImageArticle from './ImageArticle.vue';
import PoweringUpArticle from './PoweringUpArticle.vue';

</script>

<template>
	<section class="heroSection">
		<PoweringUpArticle />
		<ImageArticle />
	</section>
</template>

<style>
.heroSection {
	background-color: #fff;
	height: 746px;
	display: flex;
	justify-content: space-around;
	align-items: center;
	padding: 0px 10vw 0px 10vw;
}

@media only screen and (min-width: 1920px) {
	.heroSection {
		height: 1023px;
	}
}
</style>