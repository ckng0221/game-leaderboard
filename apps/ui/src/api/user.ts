export interface IUser {
  ID: string;
  username: string;
  role: string;
}
export interface IRankScore {
  Score: number | string;
  Rank: number | string;
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

export async function createTestUser() {
  const endpoint = `${url}?last=true`;
  const headers = new Headers();
  headers.append("Content-Type", "application/json");

  // get the last user
  let userId;
  try {
    const res = await fetch(endpoint, {
      headers: headers,
      method: "GET",
    });

    if (res.ok) {
      const users: IUser[] = await res.json();
      userId = users?.[0]?.ID || 0;
      userId = Number(userId);
    } else {
      console.error(res.status, res.statusText);
      console.error(res.body);
      return;
    }
  } catch (err) {
    console.error(err);
    return;
  }

  const usernameNew = `Player${userId + 1}`;

  try {
    const payload = [{ username: usernameNew }];
    const res = await fetch(url, {
      headers: headers,
      method: "POST",
      body: JSON.stringify(payload),
    });

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
