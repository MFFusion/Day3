<script setup>
import NavBar from "./components/NavBar.vue";
</script>

<template>
 <NavBar />
  <router-view class="flex flex-col items-center max-w-3xl mx-auto" v-slot="{ Component }">
    <transition name="slide-fade">
      <component :is="Component" />
    </transition>
  </router-view>

</template>

<style scoped>
.slide-fade-enter-active {
  transition: all 0.5s ease-out;
}
.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateY(2em);
  opacity: 0;
}
</style>
