<template>
  <div class="form-body">
    <h2>Qst/Pst verification</h2>
      <form>
        <label>Pst</label>
        <input v-model="Pst" type="text">
        <label>Qst</label>
        <input v-model="Qst" type="text">
        <label>Entreprise name</label>
        <input v-model="entreprise" type="text">
        <button @click="emitTaxForm" type="button">Check validity</button>
      </form>
  </div>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import axios from "axios";
import {Tax, TaxRequest} from "../../../Model";


export default defineComponent ({
  name: "TaxVerification",
  props: ['newTax'],
  data(){
    return {
      Tax : <Tax>{},
      Pst: '',
      Qst: '',
      entreprise  : ''
    }
  },
  methods : {
    emitTaxForm() {
      let tax = <TaxRequest>{
        PstNumber: this.Pst,
        QstNumber: this.Qst
      }
      this.$emit('taxFormEmit', tax)
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