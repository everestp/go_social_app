<template>
<<<<<<< HEAD
  <nav>
    <router-link to="/">Home</router-link> |
    <router-link to="/about">About</router-link>
  </nav>
  <router-view/>
</template>

=======
  <q-layout view="lHh Lpr lFf">
    <NavBar/>
    <q-page-container class="bg-grey-1">
      <router-view/>

    </q-page-container>
  </q-layout>
</template>

<script>
import NavBar from "@/views/NavBar.vue";

export default {
  name: "MainLayout",
  components:{NavBar}
}
</script>

>>>>>>> 60f9811 (Create a navbr)
<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
<<<<<<< HEAD
</style>
=======
</style>
>>>>>>> 60f9811 (Create a navbr)
