<template>
  <div class="menu">
    <a href="#/">Home</a>
    <a href="#/login">Login</a>
    <a href="#/registration">Register</a>
    <a href="#/dashboard">DashBoard</a>
  </div>

</template>

<script>
export default {
  name: "NavBar"
}
</script>

<style scoped>
.menu {
  position: absolute;
  transform: translate(-50%, -50%);
  left: 50%;
  top: 5%;
  color: #ffff;
  width: 80%;
  padding: 20px 20px;
  text-align: center;

}

a {
  margin-left: 50px;
  transition: 1s;
}

a:hover {
  color: #333333;
  transition: 1s;
}

</style>