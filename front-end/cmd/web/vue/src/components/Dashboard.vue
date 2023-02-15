<template>
  <div class="glass"></div>
  <div id="dashboard">
    <h1>dash</h1>
  </div>
</template>

<script>
export default {
  name: 'App',
  components: {

  }
}
</script>

<style>

.glass {
  position: absolute;
  transform: translate(-50%, -50%);
  left: 50%;
  top: 50%;
  width: 80%;
  height: 80%;

  background: rgba( 71, 67, 67, 0.15 );
  backdrop-filter: blur( 5.5px );
  -webkit-backdrop-filter: blur( 5.5px );
  border-radius: 10px;
  border: 1px solid rgba( 255, 255, 255, 0.18 );
}
#dashboard {
  position: absolute;
  transform: translate(-50%, -50%);
  left: 50%;
  top: 50%;
  width: 80%;
  height: 80%;
}
</style>