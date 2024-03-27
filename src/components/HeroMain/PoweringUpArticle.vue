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
	width: 522px;
	height: 618px;
	margin-right: 20px;
}

article p {
	color: #404040;
	font-family: Montserrat;
	text-wrap: balance;
	margin: 26px 0px 28px 0px;
	font-size: 18px;
	font-weight: 400;
	line-height: 32px;
	letter-spacing: 0em;
	text-transform: capitalize;
}

article div {
	display: flex;
}

@media only screen and (min-width: 1920px) {
	article {
		width: 715.89px;
		height: 847.54px;
	}

	article p {
		font-size: 24.69px;
		line-height: 44.43px;
	}
}

@media only screen and (max-width: 1200px) {
	article {
		min-width: 327px;
		max-width: 327px;
		height: 400px;
		margin: 0px;
	}

	article p {
		font-size: 14px;
		line-height: 140%;
		text-wrap: wrap;
	}

	article p br {
		display: none;
	}

	article div {
		display: none;
	}
}
</style>