document.querySelectorAll(".interview").forEach(function (p, i) {
    p.addEventListener("click", function(e){
        //make a rim
        //let rim = document.createElement("div");
        //rim.classList.add("interview-rim");
        this.querySelector(".interview-box").classList.add("active")
    });
});