document.addEventListener("scroll", function (e){
    if (window.scrollY > 0) {
        document.querySelector('header').classList.add("scrolled");
    } else {
        document.querySelector('header').classList.remove("scrolled");        
    }
}, {passive: true});