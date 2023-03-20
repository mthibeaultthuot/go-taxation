<script lang="ts">
import TaxVerification from "./TaxForm.vue";

export interface Tax {
  Username: String,
  IsPstValid: boolean
  IsQstValid: boolean
  PstNumber: String
  QstNumber: String
  Enterprise: String
  Date: String
}

import {defineComponent} from "vue";
import {TaxRequest} from "../../../Model";
import axios from "axios";

export default defineComponent({
  components: {TaxVerification},
  props: ['tax', 'newTax'],
  data() {
    return {
      hasNewTax: false,
    }
  },
  methods : {
    sendTax() {
      let tax = <TaxRequest>{
        PstNumber : this.newTax.PstNumber,
        QstNumber : this.newTax.QstNumber
      }
      axios.post('http://localhost:8081/service/taxation/verify',
          {
            Pst : tax.PstNumber,
            Qst : tax.QstNumber
          },
          {
            headers: {
              'Content-Type': 'application/json',
              Authorization: localStorage.getItem("token")
            }
          }
      )
          .then((response) => {
            console.log(response.data)
            console.log(response.request)
          })
          .catch(error => console.log(error))
    },
    handleNewTax() {
      console.log(this.newTax.QstNumber)
      if (this.newTax.QstNumber !== undefined) {
        this.hasNewTax = true;
        this.sendTax()
      }
    }
  },
  updated() {
    this.handleNewTax()
  }
})
</script>

<template>
  <div class="table-header">
    <div class="table-item left">Pst number</div>
    <div class="table-item center">Qst number</div>
    <div class="table-item right">Entreprise</div>
  </div>

  <div v-if="this.hasNewTax" class="table-tax-for">
    <div class="table-tax">
      <div class="table-item left">{{ this.newTax.PstNumber }}
        <div class="table-item left-icon" v-if="this.newTax.IsPstValid">
          <font-awesome-icon class="table-true-icon" icon="fa-solid fa-check" />
        </div>
      </div>
      <div class="table-item center">{{ this.newTax.QstNumber }}
        <div class="table-item left-icon" v-if="this.newTax.IsQstValid">
          <font-awesome-icon class="table-true-icon" icon="fa-solid fa-check" />
        </div>
      </div>
      <div class="table-item right">{{ this.newTax.Enterprise }}</div>
    </div>
    <div class="table-line"></div>
  </div>

  <div class="table-tax-for" v-for="item in tax">
    <div class="table-tax">
      <div class="table-item left">{{ item.PstNumber }}
        <div class="table-item left-icon" v-if="item.IsPstValid">
          <font-awesome-icon class="table-true-icon" icon="fa-solid fa-check" />
        </div>
      </div>
      <div class="table-item center">{{ item.QstNumber }}
        <div class="table-item left-icon" v-if="item.IsPstValid">
          <font-awesome-icon class="table-true-icon" icon="fa-solid fa-check" />
        </div>
      </div>
      <div class="table-item right">{{ item.Enterprise }}</div>
    </div>
    <div class="table-line"></div>
  </div>

</template>

<style scoped>

.table-header {
  position: relative;
  transform: translate(-50%, -50%);
  left: 50%;
  width: 95%;
  margin-bottom: 30px;
}
.table-tax {
  position: relative;
  transform: translate(-50%, -50%);
  left: 50%;
  width: 95%;
  text-align: left;
}

.table-line {
  position: relative;
  transform: translate(-50%, -50%);
  left: 50%;
  width: 100%;
  height: .5px;
  background: #ABAFB9;
  margin-bottom: 25px;
}

.table-item {
  display: inline;
  margin-right: 50px;
  font-family: Inter, system-ui, Avenir, Helvetica, Arial, sans-serif;
  font-weight: bold;
}
.left {
  position: absolute;
  left: 5%;
}

.center {
  position: absolute;
  left: 25%;

}

.right {
  position: absolute;
  right: 0;

}

.table-true-icon {
  color: green;
}

</style>
