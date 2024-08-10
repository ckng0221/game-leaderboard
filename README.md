```mermaid
  flowchart LR;
      Client<-->ui
      subgraph Game Leaderboard App
      ui{{UI}}<-->api;
      api{{API}}<-->db[(DB)];
      redis-->api
      api-.->rabbitmq[[Message Queue]]
      rabbitmq-.->leaderboard
      leaderboard{{Leaderboard}}-->redis[Redis]
      end
```
