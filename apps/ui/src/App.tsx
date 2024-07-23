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
import { Button, CircularProgress } from "@mui/material";
import { createGameplay } from "./api/gameplay";
import { toast } from "react-hot-toast";

function App() {
  const [leaderboards, setLeaderboards] = useState<ILeaderboard[]>([]);
  const [users, setUsers] = useState<IUser[]>([]);
  const [currentUserId, setCurrentUserId] = useState<string>("");
  const [currentUserScoreRank, setCurrentUserScoreRank] = useState<IRankScore>({
    Score: "-",
    Rank: "-",
  });
  const [leaderboardState, setLeaderboardState] = useState(0);

  useEffect(() => {
    // fetch leaderboard
    async function fetchData() {
      const data = await getTopNLeaderboard(10);
      if (!data) {
        toast.error("Failed to fetch leaderboard");
        return;
      }
      setLeaderboards(data);
    }

    fetchData();
  }, [leaderboardState]);

  useEffect(() => {
    async function fetchUsers() {
      const data = await getUsers();
      if (!data) {
        toast.error("Failed to fetch users");
        return;
      }
      setUsers(data);
    }

    fetchUsers();
  }, []);

  async function handleChangeUser(e: SelectChangeEvent<string>) {
    setCurrentUserId(e.target.value as string);
    const rankScore = await getUserRankScore(e.target.value);
    if (!rankScore) {
      setCurrentUserScoreRank({ Rank: "-", Score: "-" });
    } else {
      setCurrentUserScoreRank(rankScore);
    }
  }

  async function handlePlayGame() {
    if (!currentUserId) {
      toast.error("Please choose user first");
      return;
    }

    const data = await createGameplay({ user_id: currentUserId });
    if (!data) return;
    const rankScore = await getUserRankScore(currentUserId);
    if (!rankScore) {
      setCurrentUserScoreRank({ Rank: "-", Score: "-" });
    } else {
      setCurrentUserScoreRank(rankScore);
    }
    setLeaderboardState((prev) => prev + 1);
    toast.success("Gameplay score +10");
  }

  return (
    <>
      <h2 className="mb-8 font-bold text-xl">Game Leaderboard</h2>
      {leaderboards.length == 0 ? (
        <CircularProgress />
      ) : (
        <LeaderboardTable
          rows={leaderboards}
          columns={["Rank", "Name", "Score"]}
        />
      )}
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
            <Button variant="contained" onClick={handlePlayGame}>
              Play Game
            </Button>
          </div>
        </div>
      </div>
    </>
  );
}

export default App;
