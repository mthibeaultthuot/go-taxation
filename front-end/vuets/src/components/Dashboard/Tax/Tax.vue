<template>
  <div>
    <div class="header">
      <h1>Qst/Pst verification</h1>
      <button @click="showModal=true">New verification</button>

      <div class="nb-verification">
      </div>
      <div class="nb-true">
      </div>
      <div class="nb-false">

      </div>
    </div>

    <div class="body">
      <Table :newTax="newTax" :tax="tax"/>
      <TaxVerification @tax-form-emit="HandleNewTax" :new-tax="tax" v-if="showModal"></TaxVerification>
    </div>
  </div>

</template>

<script lang="ts">
import axios from "axios";
import {defineComponent} from "vue";
import {Tax, TaxRequest} from "../../../Model";
import Table from "./Table.vue";
import TaxVerification from "./TaxForm.vue";



export default defineComponent({
  name: "Tax",
  components: {TaxVerification, Table},
  data() {
    return {
      tax: [],
      newTax : <TaxRequest>{},
      showModal: false,
    }
  },
  methods: {
    fetchTax() {
      const bearer = {
        headers: {
          'content-type': ' application/json',
          Authorization: localStorage.getItem("token")
        }
      };
      axios.get('http://localhost:8081/service/taxation/test', bearer)
          .then((response) => {
            this.tax = response.data
          })
          .catch(error => console.log(error))
    },
    HandleNewTax(newTax : TaxRequest) {
      this.newTax = newTax;
      this.showModal = false;
    }

  },
  beforeMount() {
    this.fetchTax()
  },
})
</script>

<style scoped>
.header {
  position: relative;
  transform: translate(-50%, 0);
  top: 10%;
  left: 50%;
  width: 80%;
  height: 20%;
}

.body {
  position: absolute;
  transform: translate(-50%, 0);
  top: 30%;
  left: 50%;
  background-size: cover;
  border-radius: 10px;
  width: 80%;
  height: 60%;
}

.table-tax-for {
  position: relative;
  top: 10%;
}


</style>