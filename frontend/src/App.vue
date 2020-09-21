<template>
  <div flex-container>
    <div id="app">
      <div id="sidebar">
        <Sidebar/>
      </div>
      <div id="header">
        <div><img alt="Main header logo" src="./assets/images/Sudoku-UI-HeaderA.png" class="logo1 zoomIn"></div>
        <div><img alt="Wails logo" src="./assets/images/wails-logo-100px-height.png" class="logo2 zoomIn"></div>
        <div><img alt="Vue logo" src="./assets/images/120px-Vue.js_Logo_2.svg.png" class="logo2 zoomIn"></div>
      </div>
      <div id="mainarea">
        <div>
          <form v-on:submit.prevent="console.log(form)" onsubmit="return false;">
            <div>
              <div id="sudokuInput" class="flex-container">
                <div>
                  <div>5</div><div>3</div><div><input type="text" id="arrayindex" name="index" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[0]"><input type="submit" class="hidden"></div>
                  <div>6</div><div><input type="text" id="arrayindex" name="index" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[1]"><input type="submit" class="hidden"></div>
                  <div><input type="text" id="arrayindex" name="index" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[2]"><input type="submit" class="hidden"></div>
                  <div><input type="text" id="arrayindex" name="index" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[3]"><input type="submit" class="hidden"></div><div>9</div><div>8</div>
                </div>
                <div>
                  <div><input type="text" id="arrayindex" name="4" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[4]"><input type="submit" class="hidden"></div><div>7</div>
                  <div><input type="text" id="arrayindex" name="5" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[5]"><input type="submit" class="hidden"></div><div>1</div>
                  <div>9</div><div>5</div><div><input type="text" id="arrayindex" name="6" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[6]"><input type="submit" class="hidden"></div>
                  <div><input type="text" id="arrayindex" name="7" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[7]"><input type="submit" class="hidden"></div>
                  <div><input type="text" id="arrayindex" name="8" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[8]"><input type="submit" class="hidden"></div>
                </div>
                <div>
                  <div><input type="text" id="arrayindex" name="9" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[9]"><input type="submit" class="hidden"></div>
                  <div><input type="text" id="arrayindex" name="10" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[10]"><input type="submit" class="hidden"></div>
                  <div><input type="text" id="arrayindex" name="11" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[11]"><input type="submit" class="hidden"></div>
                  <div><input type="text" id="arrayindex" name="12" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[12]"><input type="submit" class="hidden"></div>
                  <div><input type="text" id="arrayindex" name="13" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[13]"><input type="submit" class="hidden"></div>
                  <div><input type="text" id="arrayindex" name="14" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[14]"><input type="submit" class="hidden"></div>
                  <div><input type="text" id="arrayindex" name="15" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[15]"><input type="submit" class="hidden"></div><div>6</div>
                  <div><input type="text" id="arrayindex" name="16" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[16]"><input type="submit" class="hidden"></div>
                </div>
                <div>
                  <div>8</div><div><input type="text" id="arrayindex" name="17" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[17]"><input type="submit" class="hidden"></div>
                  <div><input type="text" id="arrayindex" name="18" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[18]"><input type="submit" class="hidden"></div><div>4</div>
                  <div><input type="text" id="arrayindex" name="19" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[19]"><input type="submit" class="hidden"></div>
                  <div><input type="text" id="arrayindex" name="20" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[20]"><input type="submit" class="hidden"></div><div>7</div>
                  <div><input type="text" id="arrayindex" name="21" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[21]"><input type="submit" class="hidden"></div>
                  <div><input type="text" id="arrayindex" name="22" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[22]"><input type="submit" class="hidden"></div>
                </div>
                <div>
                  <div><input type="text" id="arrayindex" name="23" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[23]"><input type="submit" class="hidden"></div><div>6</div>
                  <div><input type="text" id="arrayindex" name="24" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[24]"><input type="submit" class="hidden"></div><div>8</div>
                  <div><input type="text" id="arrayindex" name="25" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[25]"><input type="submit" class="hidden"></div><div>3</div>
                  <div><input type="text" id="arrayindex" name="26" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[26]"><input type="submit" class="hidden"></div><div>2</div>
                  <div><input type="text" id="arrayindex" name="27" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[27]"><input type="submit" class="hidden"></div>
                </div>
                <div>
                  <div><input type="text" id="arrayindex" name="28" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[28]"><input type="submit" class="hidden"></div>
                  <div><input type="text" id="arrayindex" name="29" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[29]"><input type="submit" class="hidden"></div><div>3</div>
                  <div><input type="text" id="arrayindex" name="30" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[30]"><input type="submit" class="hidden"></div>
                  <div><input type="text" id="arrayindex" name="31" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[31]"><input type="submit" class="hidden"></div><div>1</div>
                  <div><input type="text" id="arrayindex" name="32" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[32]"><input type="submit" class="hidden"></div>
                  <div><input type="text" id="arrayindex" name="33" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[33]"><input type="submit" class="hidden"></div><div>6</div>
                </div>
                <div>
                  <div><input type="text" id="arrayindex" name="34" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[34]"><input type="submit" class="hidden"></div>
                  <div>6</div><div><input type="text" id="arrayindex" name="35" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[35]"><input type="submit" class="hidden"></div>
                  <div><input type="text" id="arrayindex" name="36" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[36]"><input type="submit" class="hidden"></div>
                  <div><input type="text" id="arrayindex" name="37" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[37]"><input type="submit" class="hidden"></div>
                  <div><input type="text" id="arrayindex" name="38" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[38]"><input type="submit" class="hidden"></div>
                  <div><input type="text" id="arrayindex" name="39" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[39]"><input type="submit" class="hidden"></div>
                  <div><input type="text" id="arrayindex" name="40" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[40]"><input type="submit" class="hidden"></div>
                  <div><input type="text" id="arrayindex" name="41" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[41]"><input type="submit" class="hidden"></div>
                </div>
                <div>
                  <div><input type="text" id="arrayindex" name="42" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[42]"><input type="submit" class="hidden"></div>
                  <div><input type="text" id="arrayindex" name="43" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[43]"><input type="submit" class="hidden"></div>
                  <div><input type="text" id="arrayindex" name="44" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[44]"><input type="submit" class="hidden"></div>
                  <div>4</div><div>1</div><div>9</div>
                  <div><input type="text" id="arrayindex" name="45" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[45]"><input type="submit" class="hidden"></div><div>8</div>
                  <div><input type="text" id="arrayindex" name="46" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[46]"><input type="submit" class="hidden"></div>
                </div>
                <div>
                  <div>2</div><div>8</div><div><input type="text" id="arrayindex" name="47" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[47]"><input type="submit" class="hidden"></div>
                  <div><input type="text" id="arrayindex" name="48" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[48]"><input type="submit" class="hidden"></div>
                  <div><input type="text" id="arrayindex" name="49" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[49]"><input type="submit" class="hidden"></div>
                  <div>5</div><div><input type="text" id="arrayindex" name="50" pattern="[1-9]{1}" title="Allowed only numbers from 1 to 9" v-model="playerInput[50]"><input type="submit" class="hidden"></div>
                  <div>7</div><div>9</div>
                </div>
              </div>
              <h1>{{message}}</h1>
                <div id="button-panel">
                  <div><button class="button" @clic="onSubmit">VERIFY</button></div>
                  <div><button @click='resetForm' class="button">RESTART</button></div>
                </div>
            </div>
          </form> 
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Sidebar from "./components/Sidebar.vue";
//import LaunchSudokuHome from "./components/LaunchSudokuHome.vue";
//import Home from "./components/Home.vue";
//import Settings from "./components/Settings.vue";
import "./assets/css/main.css";

export default {
  name: "app",
  components: {
    Sidebar,
    //LaunchSudokuHome,
    //Home,
    //Settings,
  },
  data() {
    return {
      sendPlayerInput: "",
      playerInput: [],
      message: " ",
      //test: "",
    };
  },
  methods: {
    onSubmit: function() {
      window.backend.sendPlayerInput(this.playerInput).then(result => {
        this.sendPlayerInput = result;
        this.message = result;
      });
    },
    resetForm: function() {
      this.playerInput = [];
    }
  }
};
</script>

<style scoped>
  .flex-container {
    display: flex;
    flex-wrap: wrap;
    width: 540px;
    background-color: DodgerBlue;
    margin-top: 7%;
  }

  .flex-container > div {
    display: flex;
    flex-wrap: wrap;
    color:black;
    width: 165px;
    height: 165px;
    margin: 5px;
    text-align: center;
    line-height: 160px;
    border: 2px solid black;
  }

  .flex-container > div > div {
    background-color: #f1f1f1;
    color:black;
    width: 45px;
    height: 45px;
    margin: 5px;
    text-align: center;
    line-height: 45px;
    font-size: 20px;
  }

  .hidden {
    opacity: 0;
  }

  #arrayindex {
    text-align: center;
  }

  #button-panel {
    display: flex;
    justify-content: space-around;
    align-items: center;
    width: 80%;
    position: relative;
    left: 10%;
    margin-top: 5%;
  }

  .button {
  background-color: orange;
  border: none;
  color: white;
  padding: 15px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  border-radius: 20px;
  font-size: 20px;
  margin: 2px 2px;
  cursor: pointer;
}

input{
  box-sizing:border-box;
  width: 90%;
  font-size: 20px;
} 

h1 {
  position:relative;
  z-index:9999;
}

</style>
