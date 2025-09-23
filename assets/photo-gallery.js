document.addEventListener("DOMContentLoaded", async () => {
  const grid = document.getElementById("image-grid");
  const see_more_button = document.getElementById("see-more-button");
  const pageSize = 4;
  let page = 0;
  let photos = [];

  // Load the full list of image filenames from JSON
  const response = await fetch("/USAA_champs.json");
  photos = await response.json();

  // Function to load the next batch
  function loadMore() {
    const start = page * pageSize;
    const end = start + pageSize;
    const batch = photos.slice(start, end);

    batch.forEach(name => {
      const img = document.createElement("img");
      img.src = `/images/photo_gallery${name}`;
      img.alt = "Photo";
      img.loading = "lazy";
      grid.appendChild(img);
    });

    page++;

    // Hide button if no more images
    if (end >= photos.length) {
      button.style.display = "none";
    }
  }

  // Load first batch on page load
  loadMore();

  button.addEventListener("click", loadMore);
});
