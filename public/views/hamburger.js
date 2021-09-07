// {{ define "js-hamburger" }}
const main = document.getElementById("main");
const hamburgerMenu = document.getElementById("hamburger-menu");
const hamburgerIcon = document.getElementById("hamburger-icon");
const hamburgerExitIcon = document.getElementById("hamburger-exit-icon");
hamburgerIcon.addEventListener("click", () => {
    hamburgerMenu.style.display = '';
    main.style.display = 'none';
});
hamburgerExitIcon.addEventListener("click", () => {
    hamburgerMenu.style.display = 'none';
    main.style.display = '';
});
// {{ end }}
