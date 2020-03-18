<template>
  <div id="app">
    <nav class="navbar navbar-default">
    <div class="continer-fluid">
      <a class="navbar-brand" href="#">Kartaca</a>
    <div class="collapse navbar-collapse" id="navbarNavAltMarkup">
      <div class="nav navbar-nav">
       <li><router-link to="home">Home</router-link></li>
          <li v-if="!isAuthenticated()"><router-link to="login">Login</router-link></li>
          <li v-if="!isAuthenticated()"><router-link to="signup">Sign Up</router-link></li>
          <li v-if="isAuthenticated()"><router-link to="newpost">New Post</router-link></li>
          <li v-if="isAuthenticated()"><router-link to="userinfo">User Info</router-link></li>
          <li v-if="isAuthenticated()"><router-link to="logout" @click.native="logout()">Logout</router-link></li>
      </div>
    </div>
    </div>
</nav>



    <div class="container">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import auth from './auth'

export default {

  data () {
    return {
    }
  },

  methods: {
    logout () {
      auth.logout(this)
    },
    isAuthenticated () {
      return auth.isAuthenticated()
    }
  }
}
</script>
