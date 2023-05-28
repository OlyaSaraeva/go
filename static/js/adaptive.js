document.addEventListener('DOMContentLoaded', function () {
  console.log('Готов!');
  initEventsListener();
});

function initEventsListener() {
  let burger = document.getElementById('burger');
  burger.addEventListener('click', function () {
    
    openMenu()
  })
}

function openMenu() {
  let burgerMenu = document.getElementById('burgerMenu');
  burgerMenu.classList.toggle('burger__menu')
  burgerMenu.classList.toggle('burger-menu')
}