<script setup lang="ts">
</script>

<template>
  <article class="contactUs">
    <h1>CONTACT US</h1>
    <img class="underline" src="/src/assets/Footer/contactUsLine.svg" alt="" />
    <ul>
      <li>
        <img class="icon" src="/src/assets/Footer/phone.svg" alt="" />
        <a href="">1 (723) 393 5824</a>
      </li>
      <li>
        <img class="icon" src="/src/assets/Footer/email.svg" alt="" />
        <a href="">malik@solartrade.com</a>
      </li>
      <li>
        <img class="icon" src="/src/assets/Footer/location.svg" alt="" />
        <a href="">Address NY</a>
      </li>
    </ul>
  </article>
</template>

<style scoped>
article h1 {
  margin: 0px 0px -10px 0px;
  font-family: Montserrat Alternates;
  font-size: 15px;
  font-weight: 600;
  line-height: 22px;
  letter-spacing: 0em;
  text-align: left;
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

article.contactUs ul li img {
  margin-bottom: -2.5px;
  margin-right: 4px;
}

@media only screen and (min-width: 1920px) {
  article h1 {
    font-size: 20px;
  }

  article img.underline {
    width: 238px;
  }

  article ul {
    font-size: 18px;
    line-height: 21.6px;
  }


  article ul li img.icon {
    width: 24px;
    margin-bottom: -6px;
  }
}

@media only screen and (max-width: 1200px) {
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