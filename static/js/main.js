// import { initNavbar } from "./header";

// document.addEventListener("DOMContentLoaded", () => {
//     console.log("Website loaded!");
//     initNavbar();
//     initFooter();
// });

// Slick Slider [Tools]
$('.tools-flex').slick({
    autoplay: true,
    autoplaySpeed: 0,
    dots: false,
    arrows: false,
    infinite: true,
    speed: 2000,
    slidesToShow: 8,
    slidesToScroll: 1,
    cssEase: 'linear'
});

// Initialize AOS
AOS.init({
	once: false,
	easing: 'ease',
	offset: 20,
	duration: 1000
});