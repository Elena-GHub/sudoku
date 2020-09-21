<template>
    <div class="sidenav">
        <div>
            <h3>Gophers<br>Sudoku</h3>
        </div>
        <div class="listItems">
            <ul>
                <li>
                    <router-link to="/"><span class="button-secondary fa fa-home mr-3"></span> Home</router-link>
                </li>
                <li>
                    <router-link to="/settings"><span class="button-secondary fa fa-cog mr-3"></span> Settings</router-link>
                </li>
                <li>
                    <a class="button-secondary"><span class="fa fa-puzzle-piece mr-3"></span> Games</a>
                </li>
                <li>
                    <a class="button-secondary"><span class="fa fa-info mr-3"></span> Credits</a>
                </li>
            </ul>
        </div>
        <div>
            <h3><span></span><strong> Goal</strong></h3>
            <p><span></span> Fill in the numbers 1-9 <u>exactly once</u> in every row, column, and 3x3 region.</p>
        </div>
        <div class="user-gopher">
            <div><img style="background-image" src="file://D:\Sites\Go\sudoku\frontend\src\assets\images\Gopher-peluch-contrapicado-S.png"></div>
        </div>
    </div>
</template>

<script>


</script>

<style>
    * {
        font-family: "Lato", sans-serif;
    }

    .sidenav {
        height: 100%;
        width: 20%;
        position: fixed;
        z-index: 1;
        top: 0;
        left: 0;
        bottom: 0;
        overflow-x: hidden;
        padding-top: 20px;
        text-align: left;
        padding-left: 1rem;
    }

    .sidenav a {
        padding: 6px 8px 6px 5px;
        text-decoration: none;
        font-size: 1rem;
        color: #818181;
        display: block;
    }

    .sidenav a:hover {
        color: #f1f1f1;
    }

    li {
        font-size: .9rem;
        width: 95%;
        margin: 4px;
        list-style: none;
        object-fit: contain;
    }
    h3 {
        color: orange;
        padding-top: 2rem;
    }
    p {
        color: orange;
    }

    .user-gopher {
        position: absolute;
        left: 0;
        bottom: 0;
        background-size: cover;
        background-repeat: no-repeat;
        background-position: bottom;
    }

</style>