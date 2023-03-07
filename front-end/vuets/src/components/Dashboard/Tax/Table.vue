<script lang="ts">
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

export default defineComponent({
  props: ['tax']
})
</script>

<template>
  <div class="table-header">
    <div class="table-item left">Pst number</div>
    <div class="table-item center">Qst number</div>
    <div class="table-item right">Entreprise</div>
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
