export function setupBurgerMenu(burger_btn_id, burger_menu_id, burger_close_btn_id){
    const body = document.querySelector("body");
    const burger_btn = document.getElementById(burger_btn_id);
    const burger_menu = document.getElementById(burger_menu_id);
    const burger_close_btn = document.getElementById(burger_close_btn_id);

    if (!burger_btn || !burger_menu || !burger_close_btn) return;

    burger_btn.addEventListener('click', ()=> {
        burger_menu.style.display = "block";
        body.style.overflowY = "hidden";
    });

    burger_close_btn.addEventListener('click', () => {
        burger_menu.style.display = "none";
        body.style.overflowY = "visible";
    });
}