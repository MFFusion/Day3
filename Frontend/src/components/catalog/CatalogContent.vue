<script setup>

</script>

<template>
  <section>
    <div class="grid grid-cols-3 gap-2.5 p-2">
      <div v-for="_ in 10" class="border rounded-md grid grid-cols-2 p-2">
        <div class="h-fit mx-1">
          <h1 class="text-3xl">
            Title
          </h1>
          <p>
            Lorem ipsum dolor sit amet consectetur adipisicing elit. Quisquam, voluptatum.
          </p>
        </div>
        <img src="https://picsum.photos/200/300" alt="" class=" rounded-md" />
        </div>

    </div>

  </section>
</template>

<style scoped>

</style>