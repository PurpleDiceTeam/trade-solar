<script setup lang="ts">
import ContactUsLinks from './ContactUsLinks.vue';
import GetFreeConsult from './GetFreeConsult.vue';
import QuickLinks from './QuickLinks.vue';
import SocialMediaLinks from './SocialMediaLinks.vue';
</script>

<template>
  <footer class="footerDiv">
    <link rel="stylesheet" href="/css/footer.css" />
    <div class="articleContainer">
      <GetFreeConsult />
      <div class="linksContainer">
        <div class="linksListsContainer">
          <QuickLinks />
          <ContactUsLinks />
        </div>
        <div class="socialMediaContainer">
          <SocialMediaLinks />
        </div>
      </div>
    </div>
    <p>© 2023 Trade Solar Inc · All Rights Reserved.</p>
  </footer>
</template>

<style>
.footerDiv {
  min-height: 395px;
  background: linear-gradient(140.89deg, rgba(39, 85, 243, 0.14) 0%, rgba(20, 117, 238, 0.14) 50%, rgba(0, 149, 233, 0.14) 100%);
  text-align: center;
}

.footerDiv p {
  margin: 60px 0px 0px 0px;
}

.footerDiv div {
  display: flex;
  justify-content: space-between;
}

.footerDiv div.linksContainer {
  margin: 97px 218px 0px 0px;
  display: flex;
  flex-direction: column;
}

.footerDiv div.linksContainer div.linksListsContainer {
  margin-bottom: 10px;
}


.footerDiv div.linksContainer div.socialMediaContainer {
  margin-left: 150px;
}
</style>