<template>
  <h1>Monitoring microservice</h1>
  <hr>
  <button @click="sendBroker()" id="btn_test" class="m">Test broker service</button>
  <button @click="sendVerification()" id="btn_test" class="m">Test verification service</button>
  <div id="output">
    <p>output</p>
  </div>
  <div id="send" class="m">
    <h3>Sent</h3>
    <div>{{ request }}</div>
  </div>
  <div id="received">
    <h3>Received</h3>
    <div> {{ response }}</div>
  </div>
</template>

<script lang="ts">
import {defineComponent} from 'vue'
import axios from "axios";

export interface Tax {
  Pst: string,
  Qst: string,
}

export default defineComponent({
  name: "monitoring",
  data() {
    return {
      response: '',
      request: '',
    }
  },
  methods: {
    sendBroker() {
      axios.post('http://localhost:8181/', "")
          .then((response) => {
            this.response = response.data;
          })
          .catch(error => console.log(error))
    },
    sendVerification() {
      let tax = <Tax>{
        Pst: '121476329RT',
        Qst: '1018199561TQ0001'
      }
      let json = JSON.stringify(tax)
      this.request = json

      axios.post('http://localhost:8081/service/taxation/verify', json)
          .then((response) => {
            this.response = response.data;
          })
          .catch(error => console.log(error))
    }
  }
})
</script>

<style scoped>
textarea {
  background: transparent;
  color: #333;
  border: none;
}
#btn_test {
  position: relative;
  border: solid black 1px;
  background: white;
  color: #1a1a1a;

  border-radius: 8px;
  padding: 0.6em 1.2em;
  font-size: 1em;
  font-weight: 500;
  font-family: inherit;
  cursor: pointer;
  transition: border-color 0.25s;
  margin-top: 10px;
  margin-bottom: 10px;
}

#output {
  position: relative;
  border: solid #1a1a1a 1px;
  width: 100%;
  height: 20%;
  margin-bottom: 20px;
}


#send {
  position: relative;
  border: solid #1a1a1a 1px;
  width: 40%;
  float: left;
}


#received {
  position: relative;
  border: solid #1a1a1a 1px;
  width: 40%;
  float: right;
}


</style>