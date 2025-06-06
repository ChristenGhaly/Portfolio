export function displayInfo() {
    const body = document.querySelector("body");
    const info_btn = document.getElementById("info-btn");
    const close_btn = document.getElementById("close-btn-info");
    const info_box = document.getElementById("info-box"); 

    function openInfoBox(event) {
        event.preventDefault();
        info_box.style.display = "block";
        body.style.overflowY = "hidden";
    }

    function closeInfoBox() {
        info_box.style.display = "none";
        body.style.overflowY = "auto";
    }

    info_btn.addEventListener("click", openInfoBox);
    close_btn.addEventListener("click", closeInfoBox);
}

// function initNavbar() {
//     const header = document.getElementById('header');
    
//     window.onscroll = function(){
//         if (window.pageYOffset > 100) {
//             header.style.background = 'linear-gradient(to bottom, #000, #1A1A19)';
//         } else {
//             header.style.background = 'linear-gradient(to bottom, #000, transparent)';
//         }
//     };
// }

// export { initNavbar };