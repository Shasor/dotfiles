document.addEventListener("DOMContentLoaded", function () {
  const searchInput = document.getElementById("search");
  const suggestionsDropdown = document.getElementById("suggestions");
  const albumContainer = document.querySelector(".album-container");
  const albums = Array.from(albumContainer.querySelectorAll(".album"));

  searchInput.addEventListener("input", function () {
    const query = this.value.trim().toLowerCase();
    updateSearch(query);
  });

  document.addEventListener("click", function (event) {
    if (event.target !== searchInput && event.target !== suggestionsDropdown) {
      suggestionsDropdown.style.display = "none";
    }
  });

  function updateSearch(query) {
    if (query.length > 0) {
      fetch(`/search?q=${encodeURIComponent(query)}`)
        .then((response) => response.json())
        .then((suggestions) => {
          updateSuggestions(suggestions);
          filterAlbums(query, suggestions);
        })
        .catch((error) => console.error("Error:", error));
    } else {
      suggestionsDropdown.style.display = "none";
      albums.forEach((album) => (album.style.display = ""));
    }
  }

  function updateSuggestions(suggestions) {
    suggestionsDropdown.innerHTML = "";
    const uniqueSuggestions = new Map();

    suggestions.forEach((suggestion) => {
      const key = `${suggestion.value} - ${suggestion.type}`;
      if (!uniqueSuggestions.has(key)) {
        uniqueSuggestions.set(key, suggestion);
      }
    });

    uniqueSuggestions.forEach((suggestion, key) => {
      const div = document.createElement("div");
      div.textContent = key;
      div.addEventListener("click", () => {
        searchInput.value = suggestion.value;
        suggestionsDropdown.style.display = "none";
        updateSearch(suggestion.value.toLowerCase());
      });
      suggestionsDropdown.appendChild(div);
    });

    suggestionsDropdown.style.display =
      suggestions.length > 0 ? "block" : "none";
  }

  function filterAlbums(query, suggestions) {
    const locationMatches = new Set();
    suggestions.forEach((suggestion) => {
      if (suggestion.type === "location") {
        locationMatches.add(suggestion.value.toLowerCase());
      }
    });

    albums.forEach((album) => {
      const albumName = album.querySelector("h2").textContent.toLowerCase();
      const creationDate = album.querySelector("h3").textContent.toLowerCase();
      const albumId = parseInt(album.querySelector("a").href.split("/").pop());

      const shouldShow =
        albumName.includes(query) ||
        creationDate.includes(query) ||
        suggestions.some(
          (suggestion) =>
            suggestion.id === albumId &&
            (suggestion.type === "member" ||
              suggestion.type === "first album" ||
              (suggestion.type === "location" &&
                locationMatches.has(suggestion.value.toLowerCase())))
        );

      album.style.display = shouldShow ? "" : "none";
    });
  }
});
