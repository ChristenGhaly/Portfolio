// Header [Burger Menu]-------------------------------------
function initNavbar() {
    // Header BG onScroll---------------------------------------
    window.onscroll = function(){
        var header = document.getElementById('header');
        if (window.pageYOffset > 100) {
            header.style.background = 'linear-gradient(to bottom, #000, #1A1A19)';
        } else {
            header.style.background = 'linear-gradient(to bottom, #000, transparent)';
        }
    };
    // ---------------------------------------------------------
    
    const burgerBtn = document.getElementById('burger-btn');
    const burgerMenu = document.getElementById('burger-menu');
    const burgerCloseBtn = document.getElementById('burger-close-btn');

    burgerBtn.addEventListener('click', burgerBtnFunc);
    function burgerBtnFunc() {
        burgerMenu.style.display = "block";
        body.style.overflowY = "hidden";
    }

    burgerCloseBtn.addEventListener('click', burgerCloseBtnFunc);
    function burgerCloseBtnFunc() {
        burgerMenu.style.display = "none";
        body.style.overflowY = "visible";
    }
}

export { initNavbar };
// ---------------------------------------------------------