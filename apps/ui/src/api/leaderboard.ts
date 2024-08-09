import { IRankScore as IRankScore } from "./user";

export interface ILeaderboard {
  Rank: number;
  Username: string;
  Score: number;
}

const BASE_URL = import.meta.env.VITE_API_BASE_URL;
const url = `${BASE_URL}/leaderboard`;

export async function getTopNLeaderboard(n: number) {
  const endpoint = `${url}?top=${n}`;
  try {
    const res = await fetch(endpoint, { method: "GET" });
    if (res.ok) {
      const data: ILeaderboard[] = await res.json();
      return data;
    }
    console.error(res.status, res.statusText);
    console.error(res.body);
  } catch (err) {
    console.error(err);
  }
}

export async function getUserRankScore(userId: string) {
  const endpoint = `${url}/users/${userId}`;
  try {
    const res = await fetch(endpoint, { method: "GET" });
    if (res.ok) {
      const data: IRankScore = await res.json();
      return data;
    }
    console.error(res.status, res.statusText);
    console.error(res.body);
  } catch (err) {
    console.error(err);
  }
}
