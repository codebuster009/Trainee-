<script>
    import { useNavigate } from "svelte-navigator";
    const navigate = useNavigate();

let register = {
userName :"",
email : "",
password : "",
    }
function handleSubmit(e) {
    e.preventDefault();
    localStorage.setItem("register" , JSON.stringify(register))
    navigate("/login")
}
</script>

<main>
<form on:submit={handleSubmit} >
    <input type="text" placeholder="Create your username" bind:value={register.userName} />
    <input type="email" placeholder="Enter your email" bind:value={register.email} />
    <input type="password" placeholder="Create your password" bind:value={register.password} />
    <button>Register</button>
  </form>
</main>
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
