import { useEffect, useState } from "react";
import { getTopNLeaderboard } from "./api/leaderboard";
import "./App.css";
import LeaderboardTable from "./components/LeaderboardTable";

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

  const [leaderboards, setLeaderboards] = useState<any[]>([]);

  return (
    <>
      <h2>Top 10 Leaderboard</h2>
      <LeaderboardTable
        rows={leaderboards}
        columns={["Rank", "Name", "Score"]}
      />
    </>
  );
}

export default App;
