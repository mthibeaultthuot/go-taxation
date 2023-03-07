<script lang="ts">
import Monitoring from "./components/Dashboard/Admin/monitoring.vue";
import NotFound from "./components/NotFound.vue";
import Login from "./components/Login.vue";
import TaxVerification from "./components/Dashboard/Tax/TaxForm.vue";
import {defineComponent} from "vue";
import NavBar from "./components/NavBar.vue";
import Dashboard from "./components/Dashboard/Dashboard.vue";

const routes = <any>{
  '/login': Login,
  '/dashboard': Dashboard
}

export default defineComponent({
  components: {NavBar, Login},
  data() {
    return {
      currentPath: window.location.hash
    }
  },
  computed: {
    currentView() {
      return routes[this.currentPath.slice(1) || '/'] || NotFound
    }
  },
  mounted() {
    window.addEventListener('hashchange', () => {
      this.currentPath = window.location.hash
    })
  }
})
</script>

<template>
  <component :is="currentView" />
</template>

<style scoped>
.form-opacity {
  background: #333333;
}
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
