<script setup lang="ts">
</script>

<template>
  <div class="articlesContainer">
    <img class="carruselLeftArrow" src="/src/assets/Testimonials/CarruselLeftArrow.svg" alt="Carrusel left arrow" />
    <article>
      <img class="testimonialsPhoto" src="/src/assets/Testimonials/testimonialsPhoto.png" alt="" />
      <div class="textArticleContainer">
        <img class="leftQuotes" src="/src/assets/Testimonials/leftQuotes.svg" />
        <p class="mainText">
          Trade Solar has exceeded my expectations. I wholeheartedly recommend them to anyone considering a switch to
          solar energy. Choosing Trade Solar was undeniably a bright decision for me and my home.
        </p>
        <img class="rightQuotes" src="/src/assets/Testimonials/rightQuotes.svg" />
        <p class="autor">Heather Bennet</p>
        <p class="aboutAutor">CEO - CO2E Measurement Solutions</p>
      </div>
    </article>
    <img class="carruselRightArrow" src="/src/assets/Testimonials/CarruselRightArrow.svg" alt="Carrusel right arrow" />
  </div>
</template>

<style scoped>
div.articlesContainer {
  margin-top: 44px;
  display: flex;
  justify-content: space-around;
  align-items: center;
}

div.articlesContainer img.carruselLeftArrow {
  margin-right: 10px;
}

div.articlesContainer article {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin: 40px 0px 0px 0px;
}

div.articlesContainer article div {
  padding-left: 32px;
}

div.articlesContainer article div img.leftQuotes {
  position: absolute;
  margin: 20px 0px 0px -30px;
}

div.articlesContainer article div img.rightQuotes {
  position: absolute;
  margin: -100px 0px 0px 550px;
}

div.articlesContainer article div p.mainText {
  margin: 24px 0px 24px 0px;
  font-family: Montserrat;
  font-size: 24px;
  font-weight: 400;
  line-height: 34px;
  letter-spacing: 0em;
  text-align: left;
}

div.articlesContainer article div p.autor {
  margin: 24px 0px 0px 0px;
  font-family: Montserrat;
  font-size: 24px;
  font-weight: 500;
  line-height: 36px;
  letter-spacing: 0em;
  text-align: left;
}

div.articlesContainer article div p.aboutAutor {
  margin: 0px 0px 22px 0px;
  font-family: Montserrat;
  font-size: 12px;
  font-weight: 500;
  line-height: 18px;
  letter-spacing: 0em;
  text-align: left;
}

div.articlesContainer article {
  min-width: 1000px;
  max-width: 1000px;
  height: 283px;
}

@media only screen and (min-width: 1920px) {
  div.articlesContainer article img.testimonialsPhoto {
    height: 388.81px;
    margin-left: -200px;
  }

  div.articlesContainer img.carruselLeftArrow {
    margin-right: 190px;
  }

  div.articlesContainer img.carruselRightArrow {
    margin-left: 190px;
  }

  .textArticleContainer {
    min-width: 871.68px;
    height: 323.2px;
  }

  div.articlesContainer article div p.mainText {
    font-size: 32.91px;
    font-weight: 400;
    line-height: 46.08px;
  }

  div.articlesContainer article div p.autor {
    font-size: 32.91px;
    line-height: 49.37px;
  }

  div.articlesContainer article div p.aboutAutor {
    font-size: 16.46px;
    line-height: 24.69px;

  }

  div.articlesContainer article div img.rightQuotes {
    position: absolute;
    margin: -100px 0px 0px 800px;
  }
}
</style>