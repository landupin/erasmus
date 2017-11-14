function refactorMore(action) {
    if (action === "show") {
        document.getElementById('swimmers').style.display = "block";
        document.getElementById('divers').style.display = "block";
        document.getElementById('hide').style.display = "block";
        document.getElementById('show').style.display = "none";
        ScrollTo('#' + 'swimmers')
    } else {
        document.getElementById('swimmers').style.display = "none";
        document.getElementById('divers').style.display = "none";
        document.getElementById('hide').style.display = "none";
        document.getElementById('show').style.display = "block";
        ScrollTo('#' + 'welcome');
    }
}
