document.addEventListener("DOMContentLoaded", function () {
  // Éléments DOM
  const elements = {
    creationDate: {
      min: document.querySelector("#creationDateMin"),
      max: document.querySelector("#creationDateMax"),
      minInput: document.querySelector("#creationDateMinInput"),
      maxInput: document.querySelector("#creationDateMaxInput"),
    },
    firstAlbum: {
      min: document.querySelector("#firstAlbumMin"),
      max: document.querySelector("#firstAlbumMax"),
      minInput: document.querySelector("#firstAlbumMinInput"),
      maxInput: document.querySelector("#firstAlbumMaxInput"),
    },
    location: {
      input: document.querySelector("#locationInput"),
      list: document.querySelector("#locationsList"),
    },
    applyFiltersBtn: document.querySelector("#applyFilters"),
  };

  // Fonctions utilitaires
  const utils = {
    getParsed: (currentFrom, currentTo) => {
      const from = parseInt(currentFrom.value, 10);
      const to = parseInt(currentTo.value, 10);
      return [from, to];
    },
    fillSlider: (from, to, sliderColor, rangeColor, controlSlider) => {
      const rangeDistance = to.max - to.min;
      const fromPosition = from.value - to.min;
      const toPosition = to.value - to.min;
      controlSlider.style.background = `linear-gradient(
                to right,
                ${sliderColor} 0%,
                ${sliderColor} ${(fromPosition / rangeDistance) * 100}%,
                ${rangeColor} ${(fromPosition / rangeDistance) * 100}%,
                ${rangeColor} ${(toPosition / rangeDistance) * 100}%, 
                ${sliderColor} ${(toPosition / rangeDistance) * 100}%, 
                ${sliderColor} 100%)`;
    },
  };

  // Contrôleurs de slider
  const sliderControllers = {
    controlFromInput: (fromSlider, fromInput, toInput, controlSlider) => {
      const [from, to] = utils.getParsed(fromInput, toInput);
      utils.fillSlider(fromInput, toInput, "#C6C6C6", "#25daa5", controlSlider);
      if (from > to) {
        fromSlider.value = to;
        fromInput.value = to;
      } else {
        fromSlider.value = from;
      }
    },
    controlToInput: (toSlider, fromInput, toInput, controlSlider) => {
      const [from, to] = utils.getParsed(fromInput, toInput);
      utils.fillSlider(fromInput, toInput, "#C6C6C6", "#25daa5", controlSlider);
      if (from <= to) {
        toSlider.value = to;
        toInput.value = to;
      } else {
        toInput.value = from;
        toSlider.value = from;
      }
    },
    controlFromSlider: (fromSlider, toSlider, fromInput) => {
      const [from, to] = utils.getParsed(fromSlider, toSlider);
      utils.fillSlider(fromSlider, toSlider, "#C6C6C6", "#25daa5", toSlider);
      if (from > to) {
        fromSlider.value = to;
        fromInput.value = to;
      } else {
        fromInput.value = from;
      }
    },
    controlToSlider: (fromSlider, toSlider, toInput) => {
      const [from, to] = utils.getParsed(fromSlider, toSlider);
      utils.fillSlider(fromSlider, toSlider, "#C6C6C6", "#25daa5", toSlider);
      if (from <= to) {
        toSlider.value = to;
        toInput.value = to;
      } else {
        toInput.value = from;
        toSlider.value = from;
      }
    },
  };

  // Configuration des écouteurs d'événements pour les sliders
  const setupSliderListeners = (type) => {
    const { min, max, minInput, maxInput } = elements[type];
    min.oninput = () => sliderControllers.controlFromSlider(min, max, minInput);
    max.oninput = () => sliderControllers.controlToSlider(min, max, maxInput);
    minInput.oninput = () =>
      sliderControllers.controlFromInput(min, minInput, maxInput, max);
    maxInput.oninput = () =>
      sliderControllers.controlToInput(max, minInput, maxInput, min);
  };

  // Initialisation des sliders
  const initializeSliders = () => {
    setupSliderListeners("creationDate");
    setupSliderListeners("firstAlbum");
    utils.fillSlider(
      elements.creationDate.min,
      elements.creationDate.max,
      "#C6C6C6",
      "#25daa5",
      elements.creationDate.max
    );
    utils.fillSlider(
      elements.firstAlbum.min,
      elements.firstAlbum.max,
      "#C6C6C6",
      "#25daa5",
      elements.firstAlbum.max
    );
  };

  // Gestion de l'autocomplétion des locations
  const setupLocationAutocomplete = () => {
    elements.location.input.addEventListener("input", function () {
      const query = this.value;
      if (query.length < 2) {
        elements.location.list.innerHTML = "";
        elements.location.list.classList.add("hidden");
        return;
      }

      fetch(`/locations-autocomplete?q=${encodeURIComponent(query)}`)
        .then((response) => response.json())
        .then((data) => {
          elements.location.list.innerHTML = "";
          data.forEach((location) => {
            const div = document.createElement("div");
            div.textContent = location;
            div.addEventListener("click", function () {
              elements.location.input.value = this.textContent;
              elements.location.list.classList.add("hidden");
            });
            elements.location.list.appendChild(div);
          });
          elements.location.list.classList.remove("hidden");
        });
    });

    document.addEventListener("click", function (e) {
      if (
        e.target !== elements.location.input &&
        e.target !== elements.location.list
      ) {
        elements.location.list.classList.add("hidden");
      }
    });
  };

  // Application des filtres
  function applyFilters() {
    const filters = {
      creationDateMin: document.getElementById("creationDateMinInput").value,
      creationDateMax: document.getElementById("creationDateMaxInput").value,
      firstAlbumMin: document.getElementById("firstAlbumMinInput").value,
      firstAlbumMax: document.getElementById("firstAlbumMaxInput").value,
      memberCounts: Array.from(
        document.querySelectorAll('input[type="checkbox"]:checked')
      )
        .map((checkbox) => checkbox.value)
        .join(","),
      locations: document.getElementById("locationInput").value,
    };

    const queryString = new URLSearchParams(filters).toString();
    window.location.href = `/filtered-results?${queryString}`;
  }

  // Initialisation
  initializeSliders();
  setupLocationAutocomplete();
  elements.applyFiltersBtn.addEventListener("click", applyFilters);
});
