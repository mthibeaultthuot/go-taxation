<template>
  <div class="opacity-form"></div>
  <div class="form-body">
    <h2>Qst/Pst verification</h2>
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
      axios.post('http://localhost:8081/service/taxation/verify', json, bearer)
          .then((response) => {
            console.log(response.data)

          })
          .catch(error => console.log(error))
    },
    opacity() {
      const el = document.body;
      el.classList.add("opacity-form")
    }
  },

})
</script>

<style scoped>
.opacity-form {
  position: absolute;
  height: 100vh;
  width: 100vh;
  opacity: 0.1;
  top: 0;
  left: 0;
  background: #333333;
}
.form-body {
  position: absolute;
  transform: translate(-50%, -50%);
  left: 50%;
  top: 25%;
  width: 50%;
  height: 100%;
  border-radius: 10px;
  background: #ffffff;
  text-align: center;
}

form {
  position: absolute;
  transform: translate(-50%, -50%);
  left: 50%;
  top: 50%;
  width: 80%;
}

input {
  display: block;
  border: solid 1px #ABAFB9;
  background: transparent;
  margin-bottom: 50px;
  border-radius: 2px;
  color: #ABAFB9;
  width: 100%;
  height: 30px;
  text-align: center;

}



button {
  background: #535bf2;
  border: none;
  border-radius: 2px;
  width: 50%;
  height: 30px;
  color: white;
}

label, button {
  font-family: Inter, system-ui, Avenir, Helvetica, Arial, sans-serif;
  font-weight: bold;
}

</style>