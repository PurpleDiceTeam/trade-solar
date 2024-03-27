<script setup lang="ts">
import TestimonialArticle from './TestimonialArticle.vue';
import TestimonialsTittle from './TestimonialsTittle.vue';
</script>

<template>
	<section class="testimonials">
		<TestimonialsTittle />
		<TestimonialArticle />
	</section>
</template>

<style scoped>
.testimonials {
	height: 457px;
	padding: 0px 10vw 0px 10vw;
}

@media only screen and (min-width: 1920px) {
	.testimonials {
		height: 627.24px;
	}
}

@media only screen and (max-width: 1200px) {
	.testimonials {
		width: 327px;
		height: 377px;
		padding: 0px;
		box-sizing: border-box;
	}
}
</style>