export interface IGameplay {
  score: number;
  user_id: string;
  CreatedAt: string;
  UpdatedAt: string;
}

interface IGamePlayCreate {
  user_id: string;
}

const BASE_URL = import.meta.env.VITE_API_BASE_URL;
const url = `${BASE_URL}/gameplays`;

export async function getGameplays() {
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

export async function getGameplayById(id: string) {
  const endpoint = `${url}/${id}`;
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

export async function createGameplay(payload: IGamePlayCreate) {
  const endpoint = url;
  const headers = new Headers();
  headers.append("Content-Type", "application/json");
  try {
    const res = await fetch(endpoint, {
      method: "POST",
      headers: headers,
      body: JSON.stringify(payload),
    });
    if (res.ok) {
      return res.json();
    }
    console.error(res.status, res.statusText);
    console.error(res.body);
  } catch (err) {
    console.error(err);
  }
}
