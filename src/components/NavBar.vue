<script setup lang="ts">
</script>

<template>
  <nav class="navbar">
		<img src="../assets/navBar/Logo.svg" alt="Trade solar logo"/>
		<ul>
			<li><a href="#">Home</a></li>
			<li><a href="#">About</a></li>
			<li><a href="#">Service</a></li>
			<li><a href="#">Office</a></li>
		</ul>
		<a class="startedButton" href="#">Get Started</a>
		<img
			class="navArrow"
			src="../assets/navBar/navArrow.svg"
			alt="arrowIcon"
		/>
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

	.navbar img.navArrow{
		position: absolute;
		right: 9vw;
	}

	.navbar ul {
		list-style-type: none;
		text-align: center;
	}

	.navbar ul li {
		display: inline;
		padding: 16px;
		margin: 0px 16px 0px 16px;
		border-radius: 32px;
		border: 1px solid #0000;
		background:
			linear-gradient(#fafafa, #fafafa) padding-box,
			linear-gradient(140.89deg, #2755F3 0%, #1475EE 50%, #0095E9 100%) border-box;
	}

	.navbar ul li a {
		text-decoration: none;
		font-family: "Montserrat";
		color: #777777;
		font-display: 500;
		font-size: 18px;
		line-height: 21.6px;
	}

	.startedButton {
		display: inline-block;
		text-align: center;
		justify-content: center;
		text-decoration: none;
		width: 181px;
		height: 56px;
		border: solid 1px #1F1F1F;
		border-radius: 64px;
		font-family: "Montserrat Alternates";
		font-weight: 500;
		line-height: 56px;
		font-size: 20px;
		color: #fff;
		background: linear-gradient(140.89deg, #2755F3 0%, #1475EE 50%, #0095E9 100%);
	}
</style>
