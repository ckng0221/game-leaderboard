import FormControl from "@mui/material/FormControl";
import InputLabel from "@mui/material/InputLabel";
import MenuItem from "@mui/material/MenuItem";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import { useEffect, useState } from "react";
import {
  getTopNLeaderboard,
  getUserRankScore,
  ILeaderboard,
} from "./api/leaderboard";
import "./App.css";
import LeaderboardTable from "./components/LeaderboardTable";
import { getUsers, IRankScore, IUser } from "./api/user";
import { Button } from "@mui/material";

function App() {
  useEffect(() => {
    // fetch leaderboard
    async function fetchData() {
      const data = await getTopNLeaderboard(10);
      if (data) {
        setLeaderboards(data);
      }
    }

    fetchData();
  }, []);

  useEffect(() => {
    async function fetchUsers() {
      const data = await getUsers();
      if (data) {
        setUsers(data);
      }
    }

    fetchUsers();
  }, []);

  const [leaderboards, setLeaderboards] = useState<ILeaderboard[]>([]);
  const [users, setUsers] = useState<IUser[]>([]);
  const [currentUserId, setCurrentUserId] = useState<string>("");
  const [currentUserScoreRank, setCurrentUserScoreRank] = useState<IRankScore>({
    Score: "-",
    Rank: "-",
  });

  async function handleChangeUser(e: SelectChangeEvent<string>) {
    setCurrentUserId(e.target.value as string);
    const rankScore = await getUserRankScore(e.target.value);
    console.log(rankScore);
    if (!rankScore) {
      setCurrentUserScoreRank({ Rank: "-", Score: "-" });
    } else {
      setCurrentUserScoreRank(rankScore);
    }
  }

  return (
    <>
      <h2>Top 10 Leaderboard</h2>
      <LeaderboardTable
        rows={leaderboards}
        columns={["Rank", "Name", "Score"]}
      />
      <br />
      {/* Current User */}
      <div>
        <div className="grid grid-cols-4 gap-3">
          <div>User</div>
          <div>Score</div>
          <div>Rank</div>
          <div>Play</div>
          <FormControl fullWidth>
            <InputLabel id="demo-simple-select-label">User</InputLabel>
            <Select
              labelId="demo-simple-select-label"
              id="demo-simple-select"
              value={currentUserId}
              label="User"
              onChange={handleChangeUser}
            >
              {users.map((user) => {
                return (
                  <MenuItem key={user.ID} value={user.ID}>
                    {user.username}
                  </MenuItem>
                );
              })}
            </Select>
          </FormControl>
          <div>{currentUserScoreRank.Score}</div>
          <div>{currentUserScoreRank.Rank}</div>
          <div>
            <Button variant="contained">Play Game</Button>
          </div>
        </div>
      </div>
    </>
  );
}

export default App;
