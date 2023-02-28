<template>
  <div class="form">
    <div class="left">
      <h1>The simple way to verify tax number</h1>
    </div>
    <div class="right">
      <h2 class="login_h2">Tax validation</h2>
      <div class="divider"></div>
      <form>
        <label>Pst</label>
        <input v-model="Pst" type="text">
        <label>Qst</label>
        <input v-model="Qst" type="text">
        <label>Entreprise name</label>
        <input v-model="entreprise" type="text">
        <button @click="sendLogin" type="button">Check validity</button>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import axios from "axios";

export interface Tax {
  Pst: String,
  Qst: String,
}


export default defineComponent ({
  name: "TaxVerification",
  data(){
    return {
      Pst: '',
      Qst: '',
      entreprise  : ''
    }
  },
  methods : {
    sendLogin() {
      let tax = <Tax>{
        Pst : this.Pst,
        Qst : this.Qst
      }

      let json = JSON.stringify(tax)

      const bearer = {
        headers: { Authorization: localStorage.getItem("token") }
      };

      console.log(json)



      axios.post('http://localhost:8081/service/taxation/verify', json, bearer)
          .then((response) => {
            console.log(response.data)

          })
          .catch(error => console.log(error))
    }
  }
})
</script>

<style scoped>


.form {
  position: absolute;
  transform: translate(-50%, -50%);
  left: 50%;
  top: 50%;
  width: 80%;
  height: 80%;
  background: #fff;
  color: #535bf2;
  background-size: cover;
  border-radius: 10px;
}

.left {
  position: absolute;
  transform: translate(0, -50%);
  left: 2%;
  top: 50%;
  width: 34%;
  height: 95%;
  color: white;
  background: #646cff;
  text-align: center;
  border-radius: 10px;
}

.right {
  position: absolute;
  left: 35%;
  top: 0;
  text-align: center;
  width: 66%;
  height: 100%;
}

.login_h2 {
  margin-top: 10%;
  font-size: 45px;
}

.divider {
  position: absolute;
  transform: translate(-50%, -50%);
  top: 30%;
  left: 50%;
  text-align: center;
  width: 80%;
  height: 1px;
  background: #888888;
}

form {
  transform: translate(-50%, -50%);
  position: absolute;
  left: 50%;
  top: 68%;
  width: 50%;
  height: 50%;
  margin: 10px;
}

input {
  display: block;
  border: solid 1px #333333;
  background: transparent;
  margin-bottom: 50px;
  border-radius: 2px;
  color: #111111;
  width: 100%;
  height: 6.5%;
}

button {
  background: #535bf2;
  border: none;
  border-radius: 2px;
  width: 100%;
  height: 6.5%;
}

</style>