document.querySelectorAll("aside span").forEach((v, i) => {
    v.addEventListener("click", function handler(e) {
        if (this.classList.contains("slide-opened")){
            this.parentNode.style.height = "43px";            
            this.classList.remove("slide-opened");
        } else {
            this.parentNode.style.height = `${(this.nextElementSibling.childElementCount + 1) * 43}px`;            
            this.classList.add("slide-opened");
        }
    });
});

document.addEventListener("toggle-menu", function handler(ev) {
    let menu = this.querySelector("aside");
    if (menu.classList.contains("opened")){
        menu.classList.remove("opened");
        /* REMOVE BLUR */
        this.querySelector("div.scrim").remove();
    } else {
        menu.classList.add("opened");
        /* ADD BLUR */
        let scrim = this.createElement("div");
        scrim.classList.add("scrim");
        scrim.addEventListener("click", function handler(ev){
            document.dispatchEvent(new CustomEvent("toggle-menu"), {});
        });
        this.body.appendChild(scrim);
    }
});