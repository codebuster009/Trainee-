<script>
const userName = JSON.parse(localStorage.getItem("register"))
import {useNavigate } from "svelte-navigator";
const navigate = useNavigate();


function handleLogout(){
localStorage.removeItem("register")
navigate("/login")
}
</script>

<main>
<div class="header"></div>
{#if userName}
<h1 class = "user">{userName.userName}</h1>
{:else}
<h1>Username</h1>
{/if}
<button type ="button" on:click={handleLogout}>Logout</button>
</main>

<style>
.user{
    color: rgb(255, 234, 234);
    margin-left: 660px;
}
button{
    margin-left: 750px;
    background: linear-gradient(to right top, #672280, #105685);
  color: white;
}
</style>
