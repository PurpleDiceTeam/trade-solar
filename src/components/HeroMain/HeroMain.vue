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
	min-height: 746px;
	display: flex;
	justify-content: space-around;
	align-items: center;
	padding: 0px 10vw 0px 10vw;
}
</style>