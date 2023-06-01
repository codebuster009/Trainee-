<script>
 import { useNavigate} from "svelte-navigator";
 const navigate = useNavigate();

  let login = {
    email :"" ,
   password : "" ,
  }


  const handleLogin = (event) => {
           event.preventDefault();
    const loggedUser = JSON.parse(localStorage.getItem("register"));
    console.log(loggedUser)
    if(login.email === loggedUser.email && login.password === loggedUser.password){
         navigate("/meme")
    } else{
      alert("user not found")
    }
  };

  
</script>
<form on:submit={handleLogin}>
  <input type="email" placeholder="Email" bind:value={login.email} />
  <input type="password" placeholder="Password" bind:value={login.password} />
  <button>Login</button>
</form>


<style>
  form {
    display: flex;
    flex-direction: column;
    background-color: #2C2C2C;
    border-radius: 10px;
    padding: 20px;
  }
  input {
    height: 40px;
    width: 100%;
    font-size: 18px;
    border: none;
    border-bottom: 1px solid #E4E4E4;
    outline: none;
    background-color: #ffffff;
  }
  button {
    height: 40px;
    width: 100%;
    font-size: 20px;
    border: none;
    background-color: #4CAF50;
    color: #ffffff;
    border-radius: 10px;
    cursor: pointer;
  }
</style>
