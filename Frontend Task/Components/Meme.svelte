<script>
    let topText = "";
    let bottomText = "";
    function updateTop(e) {
        topText = e.target.value
    }
    function updateBottom(e) {
        bottomText = e.target.value
    }

    const getMemes= async () =>  {
       return await fetch ("https://api.imgflip.com/get_memes")
              .then((response) => response.json())
              .then((data) => data.data.memes)
    }

    const getImage = async () => {
        const memes = await getMemes();
    const randomIndex = Math.floor(Math.random() * memes.length);
    const url = memes[randomIndex].url;
    document.querySelector(".meme img").src = url;
    }
</script>

<main>
  <div class="form">
    <div class="inputs">
      <input on:input={updateTop} id = "top-text" class="form-input" type="text" placeholder="Top Text" />
      <input on:input={updateBottom} id = "bottom-text "class="form-input" type="text" placeholder="Bottom Text" />
    </div>
    <button on:click={getImage}> Get Image</button>
    <div class="meme">
      <img src="https://i.imgflip.com/1bgw.jpg" alt="A meme" />
      <h2 class = "top-h2"> {topText}</h2>
      <h2 class = "bottom-h2"> {bottomText}</h2>
    </div>
  </div>
</main>

<style>
    .form{
  display: flex;
  flex-direction: column;
  height: 100vh;
}
.inputs {
margin-top: 10px;
height: 100px;
display: flex;
justify-content: space-evenly;
align-items: center;
}
.form-input{
  height: 30px;
  width: 220px;
  font-size: 1rem;
  font-family: monospace;
  border: 1px solid #d5d4d8;
  text-indent: 5px;
  border-radius: 5px;
}
.form > button{
  width: 850px;
  height: 40px;
  align-self: center;
  background: linear-gradient(to right top, #672280, #105685);
  color: white;
  border-radius: 5px;
}
.meme > img{
    position : relative;
    margin-left: 500px;
}
.top-h2 , .bottom-h2 {
    position: absolute;
}
.top-h2 {
    top : 0 ;
    margin-top: 390px;
    margin-left: 720px;
    color: yellowgreen;
}
.bottom-h2{
    bottom : 0;
    margin-left: 720px;
    margin-bottom: 150px;
    color: yellowgreen;
}
.meme > img {
    max-width: 500px;
    max-height: 500px;
    position : relative;
} 
</style>
