<script>
	import { Router, Link, Route } from "svelte-navigator";
	import Header from "./Components/Header.svelte";
	import Meme from "./Components/Meme.svelte"
	import Login from "./Components/Login.svelte";
	import Register from "./Components/Register.svelte";
  </script>
  
  <Router>
	<Router>
  
	<nav>
	  <Link to="/">
		<h1>Registration</h1>
	  </Link>
	  <Link to="/login">
		<h1>Login</h1>
	  </Link>
	  <!-- <Link to="/meme">
		Meme
	  </Link> -->
	</nav>
  
	<Route path="/">
	  <Register/>
	</Route>
  
	<Route path="/meme">
	  <Header />
	  <Meme/>
	</Route>
  
	<Route path="/login">
	  <Login />
	</Route>
</Router>
  </Router>
  
  <style>
	h1{
		color: rgb(230, 239, 239);
		color: yellowgreen;
	}
  </style>
  
