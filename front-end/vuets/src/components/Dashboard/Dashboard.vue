<script lang="ts">
import {defineComponent} from "vue";
import Menu from "./Menu.vue";
import Tax from "./Tax/Tax.vue";
export default defineComponent({
  components: {Tax, Menu}

})
</script>

<template>
  <div class="dashboard">
    <Menu/>
    <Tax/>
  </div>
<h1>dashboard</h1>
</template>

<style scoped>

</style>
