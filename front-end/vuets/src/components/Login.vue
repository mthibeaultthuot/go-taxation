<template>
  <div class="form">
    <div class="left">
      <h1>The simple way to verify tax number</h1>
    </div>
    <div class="right">
      <h2 class="login_h2">Login</h2>
      <div class="divider"></div>
      <form>
        <label>Email</label>
        <input v-model="email" type="text">
        <label>Password</label>
        <input v-model="password" type="password">
        <button @click="sendLogin" type="button">Submit</button>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import axios from "axios";

export interface User {
  email: String,
  password: String,
}


export default defineComponent ({
  name: "Login",
  data(){
    return {
      email: '',
      password: ''
    }
  },
  methods : {
    sendLogin() {
      let user = <User>{
        email : this.email,
        password : this.password
      }

      const params = new URLSearchParams();
      params.append('username', user.email.toString());
      params.append('password', user.password.toString());

      axios.post('http://localhost:9000/v1/auth/login', params)
          .then((response) => {
            localStorage.setItem("token", response.headers.authorization)
            console.log(response)
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