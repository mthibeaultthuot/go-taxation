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
      <Table :tax="tax"/>
      <TaxVerification v-if="showModal"></TaxVerification>
    </div>
    </div>

</template>

<script lang="ts">
import axios from "axios";

export interface Tax {
  Username:   String,
  IsPstValid: boolean
  IsQstValid: boolean
  PstNumber:  String
  QstNumber:  String
  Enterprise: String
  Date:       String
}

import {defineComponent} from "vue";
import Table from "./Table.vue";
import TaxVerification from "./TaxForm.vue";


export default defineComponent({
  name: "Tax",
  components: {TaxVerification, Table},
  data() {
    return {
      tax: [],
      showModal : false,
    }
  },
  methods: {
    fetchTax() {
      const bearer = {
        headers: { Authorization: localStorage.getItem("token") }
      };
      axios.get('http://localhost:8081/service/taxation/test', bearer)
          .then((response) => {
            this.tax = response.data
          })
          .catch(error => console.log(error))
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