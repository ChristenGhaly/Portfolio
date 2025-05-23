export function setupBurgerMenu(burger_btn_id, burger_menu_id, burger_close_btn_id){
    const body = document.querySelector("body");
    const burger_btn = document.getElementById(burger_btn_id);
    const burger_menu = document.getElementById(burger_menu_id);
    const burger_close_btn = document.getElementById(burger_close_btn_id);

    if (!burger_btn || !burger_menu || !burger_close_btn) return;

    function openBurgerMenu() {
        burger_menu.classList.add("active");
        // burger_btn.style.display = "none";
        body.style.overflowY = "hidden";
    }
    
    function closeBurgerMenu() {
        burger_menu.classList.remove("active");
        // burger_btn.style.display = "block";
        body.style.overflowY = "visible";
    }

    burger_btn.addEventListener('click', openBurgerMenu);
    burger_close_btn.addEventListener('click', closeBurgerMenu);

    window.addEventListener('resize', () => {
        if (window.innerWidth >= 768){
            closeBurgerMenu();
        }

        // if (window.innerWidth <= 450){
        //     burger_menu.style.width = "100vw";
        // } else {
        //     burger_menu.style.width = "60vw";
        // }
    });

    // Close Burger_Menu when click outside the menu
    // document.addEventListener('click', (event) => {
    //     const isMenuClicked = burger_menu.contains(event.target);
    //     const isCloseBtnClicked = burger_btn.contains(event.target);
    //     const isMenuActive = burger_menu.classList.contains('active')

    //     if ( !isMenuClicked && !isCloseBtnClicked && isMenuActive) {
    //         closeBurgerMenu();
    //     }
    // });
}