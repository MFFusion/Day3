<script setup>

import SearchBar from "./SearchBar.vue";
import CatalogContent from "./CatalogContent.vue";
import {ref} from "vue";
const isFilterShown = ref(false);
const toggleFilter = () => {
  isFilterShown.value = !isFilterShown.value;
};



</script>

<template>
  <section class="min-h-screen">
    <SearchBar />
    <div class="w-full p-2">
      <button @click="toggleFilter" class="my-1 rounded-md border py-1 px-1 transition hover:bg-gray-200 active:shadow-md active:border-2 hover:border-gray-600 ">Show Filter</button>
      <transition name="slide-fade">
        <div v-if="isFilterShown" class="flex flex-col gap-1">
          <div>
            <label for="sort" class="mr-2 bg-transparent border rounded-md focus:ring-0 px-2 py-3">Sort By</label>
            <select  class="bg-transparent border rounded-md focus:ring-0 px-2 py-3">
              <option value="1">1</option>
            </select>
          </div>
          <div >
            <label for="sort" class="mr-2 bg-transparent border rounded-md focus:ring-0 px-2 py-3">Sort By</label>
            <select  class="bg-transparent border rounded-md focus:ring-0 px-2 py-3">
              <option value="1">1</option>
            </select>
          </div>

        </div>
      </transition>
    </div>
    <transition>
      <CatalogContent />

    </transition>
  </section>

</template>

<style scoped>
.slide-fade-enter-active {
  transition: all 0.2s ease-in;
}
.slide-fade-leave-active {
  transition: all 0.2s ease-out;
}
.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateY(2em);
  opacity: 0;
}


.in-out-enter-active {
  transition: all 0.1s ease-out;
}
.in-out-enter-from,
.in-out-leave-to {
  transform: translatey(2em);
  opacity: 0;
}
</style>