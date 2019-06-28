document.addEventListener("scroll", function (e){
    if (window.scrollY > 0) {
        document.querySelector('header').classList.add("scrolled");
    } else {
        document.querySelector('header').classList.remove("scrolled");        
    }
}, {passive: true});

document.querySelector("button.menu-toggle").addEventListener("mousedown", function(ev) {
    this.classList.add("active");
    let ctx = this;
    document.addEventListener("mouseup", function mouseupHandler(ev){
        ctx.classList.remove("active");
    });
});

document.querySelector("button.menu-toggle").addEventListener("touchstart", function(ev) {
    this.classList.add("active");
    let ctx = this;
    document.addEventListener("touchend", function touchendHandler(ev){
        ctx.classList.remove("active");
    });
}, {passive: true});

document.querySelector("button.menu-toggle").addEventListener("click", function handler(ev){
    document.dispatchEvent(new CustomEvent('toggle-menu', {}));
});