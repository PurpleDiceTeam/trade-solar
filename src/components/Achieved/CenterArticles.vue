<script setup lang="ts">
interface Props {
  id: string;
  titleFirstWorld: string;
  titleSecondWorld: string;
  amount: string;
}
const props = defineProps<Props>()
</script>

<template>
  <article class="centerArticle" :id="props.id">
    <h1>{{ props.titleFirstWorld }}<br>{{ props.titleSecondWorld }} </h1>
    <p>{{ props.amount }}</p>
  </article>
</template>

<style scoped>
article {
  width: 370px;
  height: 274px;
  margin: 0px 20px;
  box-sizing: border-box;
  padding: 48px;
  border-radius: 27px;
  background-color: #fafafa;
  background-image: url("data:image/svg+xml,%3csvg width='100%25' height='100%25' xmlns='http://www.w3.org/2000/svg'%3e%3crect width='100%25' height='100%25' fill='none' rx='27' ry='27' stroke='%232755F3FF' stroke-width='2' stroke-dasharray='15%2c 16' stroke-dashoffset='12' stroke-linecap='square'/%3e%3c/svg%3e");
  border-radius: 27px;
}

article h1 {
  margin: 0px;
  font-family: Montserrat Alternates;
  font-size: 37px;
  font-weight: 500;
  line-height: 44px;
  letter-spacing: 0em;
  text-align: left;
}


article p {
  margin: 0px;
  font-family: Montserrat;
  font-size: 73px;
  font-weight: 500;
  line-height: 102px;
  letter-spacing: 0.02em;
}

article#clienteleValuations {
  margin-bottom: 36px;
}

article#inFunding {
  margin-top: 36px;
}

@media only screen and (min-width: 1920px) {
  article {
    width: 507.43px;
    height: 375.74px;
    padding: 65px;
  }

  article h1 {
    font-size: 50.1px;
    line-height: 60.12px;
  }

  article p {
    font-size: 100.2px;
    line-height: 140.28px;
    letter-spacing: 0.02em;
  }

  article#clienteleValuations {
    margin-bottom: 50px;
  }

  article#inFunding {
    margin-top: 50px;
  }
}

@media only screen and (max-width: 1200px) {
  article {
    width: 327px;
    padding: 48px 32px;
    margin: 0px;
  }
}
</style>