const artistsList = document.getElementById('artistsList');
const searchInput = document.getElementById('searchInput');

async function getArtists() {
  const response = await fetch('/artists/data');
  const artists = await response.json();
  return artists;
}

function displayArtists(artists) {
  artistsList.innerHTML = ''; 

  for (const artist of artists) {
    const listItem = document.createElement('li');  

    // Создаем ссылку и добавляем её в li
    const artistLink = document.createElement('a');
    artistLink.href = `/artist/${artist.id}`;
    listItem.appendChild(artistLink);

    // Добавляем изображение в ссылку
    if (artist.image) {
      const artistImage = document.createElement('img');
      artistImage.src = artist.image;
      artistImage.alt = artist.name;
      artistLink.appendChild(artistImage);
    }

    // Добавляем имя артиста в ссылку
    const artistName = document.createElement('p');
    artistName.textContent = artist.name;
    artistLink.appendChild(artistName);

    artistsList.appendChild(listItem);
  }
}

getArtists().then(displayArtists);
