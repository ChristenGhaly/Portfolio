function initNavbar() {
    window.onscroll = function(){
        var header = document.getElementById('header');
        if (window.pageYOffset > 100) {
            header.style.background = 'linear-gradient(to bottom, #000, #1A1A19)';
        } else {
            header.style.background = 'linear-gradient(to bottom, #000, transparent)';
        }
    };
}

export { initNavbar };