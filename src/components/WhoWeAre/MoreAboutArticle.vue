<script setup lang="ts">
</script>

<template>
  <article class="moreAboutArticle">
    <h2>
      Know more about
      <br />
      our company
    </h2>
    <p>
      Harness the sun's energy with solar panels,
      <br />
      paving the way for renewable and green
      <br />
      energy solutions
    </p>
    <div>
      <a href="#">Learn more</a>
      <img src="/src/assets/WhoWeAre/ArrowWhatWeOffer.svg" alt="" />
    </div>
  </article>
</template>

<style scoped>
article {
  height: 417px;
  width: 660px;
  padding: 0px 0px 0px 64px;
  border-radius: 32px;
  background-color: #fafafa;
}

article h2 {
  font-family: Montserrat Alternates;
  font-size: 46px;
  font-weight: 500;
  line-height: 55px;
  letter-spacing: 0em;
  text-align: left;
  margin: 38px 0px;
}

article p {
  margin: 18px 0px;
  text-transform: capitalize;
}

article div {
  display: flex;
  justify-content: initial;
  align-items: center;
}

article div a {
  font-family: Montserrat Alternates;
  font-size: 20px;
  font-weight: 500;
  line-height: 24px;
  letter-spacing: 0em;
  text-align: left;
  color: #000;
  border-bottom: 1.5px solid;
  text-decoration: none;
  margin-right: 6px;
}

article p {
  font-family: Montserrat;
  font-size: 18px;
  font-weight: 400;
  line-height: 32px;
  letter-spacing: 0em;
  text-align: left;
}

@media only screen and (min-width: 1920px) {
  article {
    height: 572px;
    width: 1048px;
    box-sizing: border-box;
    padding: 47px 80px 47px 80px;
  }

  article h2 {
    font-size: 63.09px;
  }

  article p {
    font-size: 24.69px;
    line-height: 44.43px;
  }

  article div {
    margin-top: 77px;
  }

  article div a {
    font-size: 27.43px;
    line-height: 32.91px;
  }

  article div img {
    width: 40px;
    margin-bottom: -6px;
  }
}

@media only screen and (max-width: 1200px) {
  article {
    width: 327px;
    height: 417px;
    padding: 0px 0px 0px 24px;
    box-sizing: border-box;
    margin-top: 24px;
  }

  article h2 {
    font-size: 28px;
    line-height: 38.4px;
  }

  article p {
    font-size: 18px;
    line-height: 32.4px;
    margin-bottom: 50px;
    margin-right: 20px;
  }

  article p br {
    display: none;
  }
}
</style>