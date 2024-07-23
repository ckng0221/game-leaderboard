export interface IUser {
  ID: string;
  username: string;
  role: string;
}
export interface IRankScore {
  Score: number;
  Rank: number;
}

const BASE_URL = import.meta.env.VITE_API_BASE_URL;
const url = `${BASE_URL}/users`;

export async function getUsers() {
  const endpoint = url;
  try {
    const res = await fetch(endpoint, { method: "GET" });
    if (res.ok) {
      return res.json();
    }
    console.error(res.status, res.statusText);
    console.error(res.body);
  } catch (err) {
    console.error(err);
  }
}

export async function getUsersById(id: string) {
  const endpoint = `${url}/${id}`;
  try {
    const res = await fetch(endpoint, { method: "GET" });
    if (res.ok) {
      const data: IUser = await res.json();
      return data;
    }
    console.error(res.status, res.statusText);
    console.error(res.body);
  } catch (err) {
    console.error(err);
  }
}
