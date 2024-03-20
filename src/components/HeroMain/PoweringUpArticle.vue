<script setup lang="ts">
import PoweringUpForm from './PoweringUpForm.vue';
import PoweringUpStats from './PoweringUpStats.vue';
import PoweringUpTitle from './PoweringUpTitle.vue';
</script>

<template>
	<article class="poweringUpArticle">
		<PoweringUpTitle />
		<p>
			Harness the sun's energy with solar panels, paving
			<br />
			the way for renewable and green energy solutions
		</p>
		<PoweringUpForm />
		<div>
			<PoweringUpStats amount="2.5k" title="Happy Clients" />
			<PoweringUpStats amount="8.4k" title="Generated" />
		</div>
	</article>
</template>

<style scoped>
article {
	min-width: 522px;
	max-width: 522px;
	min-height: 618px;
}

article p {
	font-family: Montserrat;
	text-wrap: balance;
	margin: 26px 0px 28px 0px;
	font-size: 18px;
	font-weight: 400;
	line-height: 32px;
	letter-spacing: 0em;
}

article div {
	display: flex;
}
</style>