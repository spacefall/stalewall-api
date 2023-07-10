const cssvars = document.querySelector(":root");
const timeBetweenImagesInSecs = 10;
let nextwallurl = "";
let visibleBg = 1; //is set to 1 to have a nice transition when first loading the page

function newWall() {
    fetch("https://stalewall.vercel.app/api?bQlt=80")
        .then((res) => res.json())
        .then((out) => preloadAndSet(out["url"]))
        .catch((err) => {
            throw err;
        });
}

function preloadAndSet(url) {
    cssvars.style.setProperty("--bg-preload", "url(" + url + ")");
    if (visibleBg == 0) {
        cssvars.style.setProperty("--bg-img-1", "url(" + nextwallurl + ")");
        cssvars.style.setProperty("--bg-1-opacity", "1  ");
        visibleBg = 1;
    } else {
        cssvars.style.setProperty("--bg-img-0", "url(" + nextwallurl + ")");
        cssvars.style.setProperty("--bg-1-opacity", "0");
        visibleBg = 0;
    }
    nextwallurl = url;
}

// Gets time every second and sets the appropriate divs
function tick() {
    const now = dayjs();
    document.getElementById("clock").innerHTML = now.format("HH:mm");
    document.getElementById("date").innerHTML = now.format("dddd, DD MMMM");
}

function setupAndRun() {
    if (navigator.language === "it-IT") {
        dayjs.locale("it");
    }
    newWall();
    tick();
    newWall();
    setInterval(tick, 1000);
    setInterval(newWall, timeBetweenImagesInSecs * 1000);
}

document.addEventListener("DOMContentLoaded", setupAndRun);
