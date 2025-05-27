import { initMap } from "./map.js";
import { setupBurgerMenu } from "./burgermenu.js";

window.addEventListener('DOMContentLoaded', () => {
  initMap();
});

document.addEventListener('DOMContentLoaded', () => {
    setupBurgerMenu("burger-btn-id", "burger-menu-id", "burger-close-btn-id");
});

document.addEventListener("DOMContentLoaded", () => {
    const links = document.querySelectorAll(".menu a, .burger-menu a");
    const currentPage = window.location.pathname;

    links.forEach(link => {
        if (link.href.startsWith("https://www.linkedin.com")) return;
        
        if (currentPage === "/") {
            link.classList.toggle("active", link.href.endsWith("/"));
        } else {
            link.classList.toggle("active", link.href.includes(currentPage));
        }
    });
});

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
    cssEase: 'linear',
    responsive: [
        {
            breakpoint: 768,
            settings: {
                slidesToShow: 5
            }
        },
        {
            breakpoint: 480,
            settings: {
                slidesToShow: 4
            }
        }
    ]
});

// Slick Slider [Projects]
$('.slides').slick({
    dots: true,
    centerMode: true,
    centerPadding: '60px',
    slidesToShow: 3,
    autoplay: true,
    autoplaySpeed: 2000,
    responsive: [
        {
            breakpoint: 1100,
            settings: {
                centerPadding: '40px',
                slidesToShow: 2
            }
        },
        {
            breakpoint: 768,
            settings: {
                arrows: false,
                centerPadding: '40px',
                slidesToShow: 2
            }
        },
        {
            breakpoint: 576,
            settings: {
                arrows: false,
                slidesToShow: 1
            }
        }
    ]
});

// Initialize AOS
AOS.init({
	once: false,
	easing: 'ease',
	offset: 20,
	duration: 1000
});