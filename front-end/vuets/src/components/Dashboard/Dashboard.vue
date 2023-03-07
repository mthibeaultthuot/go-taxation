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
</template>

<style scoped>
.dashboard {
  position: absolute;
  transform: translate(-50%, -50%);
  top: 50%;
  left: 50%;
  width: 80%;
  height: 80%;
  /* From https://css.glass */
  background: rgba(255, 255, 255, 0.14);
  border-radius: 16px;
  box-shadow: 0 4px 30px rgba(0, 0, 0, 0.1);
  backdrop-filter: blur(7.6px);
  -webkit-backdrop-filter: blur(7.6px);
}
</style>
