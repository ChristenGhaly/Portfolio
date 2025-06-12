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
    info_box.addEventListener("click", (event) => {
        if (event.target == info_box) {
            closeInfoBox();
        }
    });
}