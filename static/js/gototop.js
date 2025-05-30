export function goToTop (){
    const gototop_btn = document.getElementById("go-to-top");

    window.addEventListener("scroll", () => {
        if (document.body.scrollTop > 500 || document.documentElement.scrollTop > 500) {
            gototop_btn.style.display = "block";
        } else {
            gototop_btn.style.display = "none";
        }
    });

    gototop_btn.addEventListener("click", () => {
        window.scrollTo({ top: 0, behavior: "smooth" });
    });
}