document.addEventListener("DOMContentLoaded", () => {
  fetch("/api/artists")
    .then((response) => response.json())
    .then((data) => {
      const artistsContainer = document.getElementById("artists");
      data.forEach((artist) => {
        const artistCard = document.createElement("div");
        artistCard.className = "artist-card";
        artistCard.innerHTML = `
                    <h2>${artist.Artist.Name}</h2>
                    <img src="${artist.Artist.Image}" alt="${
          artist.Artist.Name
        }" width="100">
                    <p>Created: ${artist.Artist.CreationDate}</p>
                    <p>First Album: ${artist.Artist.FirstAlbum}</p>
                    <p>Members: ${artist.Artist.Members.join(", ")}</p>
                `;
        artistsContainer.appendChild(artistCard);
      });
    })
    .catch((error) => console.error("Error:", error));
});

const searchForm = document.createElement("form");
searchForm.innerHTML = `
    <input type="text" id="search" placeholder="Search artists">
    <button type="submit">Search</button>
`;
document.body.insertBefore(searchForm, document.getElementById("artists"));

searchForm.addEventListener("submit", (e) => {
  e.preventDefault();
  const query = document.getElementById("search").value;
  fetch(`/api/search?q=${encodeURIComponent(query)}`)
    .then((response) => response.json())
    .then((data) => {
      const artistsContainer = document.getElementById("artists");
      artistsContainer.innerHTML = "";
      data.forEach((artist) => {
        // Create artist card (same as before)
      });
    })
    .catch((error) => console.error("Error:", error));
});
