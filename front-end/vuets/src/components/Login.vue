<template>
  <div class="form">
    <div class="left">
      <h1>The simple way to verify tax number</h1>
    </div>
    <div class="right">
      <div class="form-content">
        <h1 class="login_h2">Hi, welcome back</h1>
        <div class="divider"></div>
        <form>
          <label>Email Address</label>
          <div class="email-input">
            <font-awesome-icon class="icon" icon="fa-solid fa-at" />
            <input v-model="email" placeholder="tax@tax.ca" type="text">
          </div>
          <label>Password</label>
          <div class="password-input">
            <font-awesome-icon class="icon" icon="fa-solid fa-lock" />
            <input v-model="password" placeholder="8+ characters required" type="password">
          </div>
          <a class="password-forget">Forget your password click here!</a>
          <button @click="sendLogin" type="button">
            <font-awesome-icon class="icon" icon="fa-regular fa-circle-check" />
            Login
          </button>
          <div class="create-account">
            <div class="line-create-account"></div>
            <p>or create new account</p>
            <div class="line-create-account"></div>
          </div>
        </form>
      </div>
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


export default defineComponent({
  name: "Login",
  data() {
    return {
      email: '',
      password: ''
    }
  },
  methods: {
    sendLogin() {
      let user = <User>{
        email: this.email,
        password: this.password
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
  width: 100%;
  height: 100%;
  background: transparent;
  color: #535bf2;
  background-size: cover;
  border-radius: 10px;
}

.left {
  position: absolute;
  transform: translate(0, -50%);
  left: 0;
  top: 50%;
  width: 60%;
  height: 100%;
  color: white;
  background: transparent;
  text-align: left;
  border-radius: 10px;
}

.right {
  position: absolute;
  right: 0;
  top: 0;
  width: 40%;
  height: 100%;
  background: #ffffff;
}

.login_h2 {
  position: absolute;
  margin-top: 10%;
  font-size: 40px;
  top: 10%;
}

.divider {
  position: absolute;
  transform: translate(-50%, -50%);
  top: 30%;
  left: 50%;
  text-align: center;
  width: 100%;
  height: 1px;
  background: #888888;
}

form {
  transform: translate(-50%, -50%);
  position: absolute;
  left: 50%;
  top: 68%;
  width: 100%;
  height: 50%;
}

.form-content {
  position: absolute;
  transform: translate(-50%, -50%);
  left: 50%;
  top: 50%;
  width: 70%;
  height: 100%;
}

.email-input, .password-input {
  display: block;
  background: transparent;
  margin-bottom: 40px;
  border-radius: 8px;
  color: #ABAFB9;
  width: 100%;
  height: 10%;
  border: 1px solid #777;
}

label {
  font-weight: bold;
  color: #ABAFB9;
}

input {
  position: relative;
  transform: translate(0, -50%);
  left: 6%;
  top: 45%;
  background: transparent;
  border: none;
  color: #ABAFB9;
  margin: 0;
}

input:focus {
  outline: none;
}


.icon {
  position: relative;
  transform: translate(-50%, -50%);
  left: 5%;
  top: 45%;
  color: #ABAFB9;
  box-sizing: border-box;
}

.password-forget {
  position: relative;
  left: 0;
}

button {
  background: #535bf2;
  border: none;
  border-radius: 8px;
  width: 100%;
  height: 10%;
  color: white;
  font-size: 20px;
  font-weight: bold;
  margin-top: 20px;
}

.create-account {
  position: relative;
  color: #ABAFB9;
  display: inline;
}

.line-create-account {
  height: 1px;
  width: 10%;
  background: #ABAFB9;
  display: inline;
}

</style>