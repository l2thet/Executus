const executusServerURL = "http://localhost:8080";
const executusGetMusicListURL = `${executusServerURL}/api/listavailablemusic`;

export const getMusicList = async () => {
  const response = await fetch(executusGetMusicListURL);
  const musicList = await response.json();
  return musicList;
};