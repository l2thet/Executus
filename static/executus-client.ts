export const executusServerURL = "http://localhost:8080";
export const executusGetMusicListURL = `${executusServerURL}/api/listavailablemusic`;

export const getSongURL = (name: string) => {
  return `${executusServerURL}/api/servesong?name=${name}`;
};

export const getMusicList = async () => {
  const response = await fetch(executusGetMusicListURL);
  const musicList = await response.json();
  return musicList;
};