document.addEventListener("scroll", function (e){
    if (window.scrollY > 0) {
        document.querySelector('header').classList.add("scrolled");
    } else {
        document.querySelector('header').classList.remove("scrolled");        
    }
}, {passive: true});

document.querySelector("button.menu").addEventListener("mousedown", function(ev) {
    this.classList.add("active");
    this.addEventListener("mouseup", function mouseupHandler(ev){
        this.classList.remove("active");
        this.removeEventListener("animationend", mouseupHandler);
    });
});

document.querySelector("button.menu").addEventListener("touchstart", function(ev) {
    this.classList.add("active");
    this.addEventListener("touchend", function touchendHandler(ev){
        this.classList.remove("active");
        this.removeEventListener("animationend", touchendHandler);
    });
}, {passive: true});

document.querySelector("button.menu").addEventListener("click", function handler(ev){
    document.dispatchEvent(new CustomEvent('toggle-menu', {}));
});