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
	min-height: 457px;
	padding: 0px 10vw 0px 10vw;
}
</style>