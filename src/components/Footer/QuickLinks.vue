<script setup lang="ts">
</script>

<template>
  <article class="quickLinks">
    <h1>QUICK LINKS</h1>
    <img src="/src/assets/Footer/quickLinkLine.svg" alt="" />
    <ul>
      <li><a href="">About</a></li>
      <li><a href="">Services</a></li>
      <li><a href="">Partners</a></li>
    </ul>
  </article>
</template>

<style scoped>
article {
  margin-right: 44px;
}

article ul {
  padding: 0px;
  margin: 13px 0px;
  list-style-type: none;
  font-family: Montserrat;
  font-size: 13px;
  font-weight: 500;
  line-height: 16px;
  letter-spacing: 0.01em;
  text-align: left;
}

article ul li {
  margin-bottom: 18px;
}

article ul li a {
  text-decoration: none;
  color: #4E4D4D;
}

article h1 {
  margin: 0px 0px -10px 0px;
  font-family: Montserrat Alternates;
  font-size: 15px;
  font-weight: 600;
  line-height: 22px;
  letter-spacing: 0em;
  text-align: left;
}

@media only screen and (min-width: 1920px) {
  article h1 {
    font-size: 20px;
  }

  article img {
    width: 167px;
  }

  article ul {
    font-size: 18px;
    line-height: 21.6px;
  }
}

@media only screen and (max-width: 1200px) {
  article {
    margin: 0px 0px 14px 0px;
  }

  article h1 {
    text-align: center;
    leading-trim: both;
    font-size: 16px;
    line-height: 24px;
  }

  article img.underline {}

  article ul {
    text-align: center;
    leading-trim: both;
    font-size: 14px;
    line-height: 16.8px;
    letter-spacing: 0.14px;
  }
}
</style>