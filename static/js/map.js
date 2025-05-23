export function initMap() {
  const map = L.map('map').setView([45.5017, -73.5673], 13);

  L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy; OpenStreetMap contributors'
  }).addTo(map);

  L.marker([45.5017, -73.5673]).addTo(map);

  // L.marker([45.5017, -73.5673]).addTo(map)
  //   .bindPopup('Location')
  //   .openPopup();
}