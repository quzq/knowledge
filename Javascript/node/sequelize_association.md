nodejsのorマッパー sequelize のaccociaton

# associationのまとめ
## belongsTo
1つの関連をもつ。
### 例) Players.belongsTo(Teams)
Playersに外部キー`teamId`が自動追加される。

## hasOne
1つの関連をもつ。
### 例) Players.hasOne(Profiles)
Profilesに外部キー`playerId`が自動追加される。include時、最初の１つのみ出力される。

## hasMany
複数の関連をもつ。

### 例) Players.hasMany(Profiles)
Profilesに外部キー`playerId`が自動追加される。include時、複数件出力される。

## belongsToMany
多対多の中間テーブルを持つ。

### 例1)Players.belongsToMany(Teams, { through: TeamPlayers }) 
中間テーブル'TeamPlayers'が作成され、そこに外部キー`teamId` `playerId`が自動追加される    
### 例2)Teams.belongsToMany(Players, { through: TeamPlayers })   
中間テーブル'TeamPlayers'が作成され、そこに外部キー`playerId` `teamId`が自動追加される  

## belongToとhasOneの違い
挙動はほぼ同じ。指定の順が変わっただけ。

# belongsTo のサンプル

```Javascript
// model(テーブル)を定義
// 主キー`id`や`createdAt`、`updatedAt`は勝手に定義されるので記述不要
const Players = sequelize.define('players', {
    name: Sequelize.STRING,
});
const Teams = sequelize.define('teams', {
    name: Sequelize.STRING,
});
Players.belongsTo(Teams, { foreignKey: 'team_id' })   // playersに外部キー'team_id'が作成される。foreignKeyを省略した場合、キャメルケースでid名'teamId'が自動生成される

// modelを元にDBテーブルの作成({force:true}オプションは、テーブルdrop&create)
await Players.sync({ force: true });
await Teams.sync({ force: true });

// レコードのcreate
const team1 = await Teams.create({ name: 'team1' });
const player1 = await Players.create({ name: 'player1', team_id: team1.id });

// レコード取得
const player = await Players.findById(player1.id, { include: Teams })

console.log(player.get({ plain: true }))
```

出力結果

```
{ id: 1,
    name: 'player1',
    createdat: 2019-08-29t02:36:21.464z,
    updatedat: 2019-08-29t02:36:21.464z,
    team_id: 1,
    team:
    { id: 1,
    name: 'team1',
    createdat: 2019-08-29t02:36:21.448z,
    updatedat: 2019-08-29t02:36:21.448z } }
```

以下でも、teams : players = 1 : n で取得できそうだが、できない

```Javascript
const team = await Teams.findById(team1.id, { include: Players})
```
出力結果
```
  failed: sequelizeeagerloadingerror: players is not associated to teams!
```

# hasOne のサンプル
```Javascript
// model(テーブル)を定義
const Players = sequelize.define('players', {
    name: Sequelize.STRING,
});
const Profiles= sequelize.define('profiles', {
    name: Sequelize.STRING,
});
Players.hasOne(Profiles,{foreignKey:'player_id'})

// modelを元にDBテーブルの作成
await Players.sync({ force: true });
await Profiles.sync({ force: true });

// レコードのcreate
const player1= await Players.create({ name: 'player1' });
const profile1= await Profiles.create({ name: 'profile1', player_id: player1.id });
const profile2= await Profiles.create({ name: 'profile2', player_id: player1.id });

// レコードの取得
const player = await Players.findById(player1.id, { include: Profiles})

console.log(player.get({ plain: true }))
```

出力結果

```
{ id: 1,
    name: 'player1',
    createdAt: 2019-08-29T03:59:30.432Z,
    updatedAt: 2019-08-29T03:59:30.432Z,
    profile:
    { id: 1,
    name: 'profile1',
    createdAt: 2019-08-29T03:59:30.448Z,
    updatedAt: 2019-08-29T03:59:30.448Z,
    player_id: 1 } }
```


# hasMany のサンプル

```Javascript
// model(テーブル)を定義
const Players = sequelize.define('players', {
    name: Sequelize.STRING,
});
const Profiles= sequelize.define('profiles', {
    name: Sequelize.STRING,
});
Players.hasMany(Profiles,{foreignKey:'player_id'})

// modelを元にDBテーブルの作成
await Players.sync({ force: true });
await Profiles.sync({ force: true });

// レコードのcreate
const player1= await Players.create({ name: 'player1' });
const profile1= await Profiles.create({ name: 'profile1', player_id: player1.id });
const profile2= await Profiles.create({ name: 'profile2', player_id: player1.id });

//　レコードの取得
const player = await Players.findById(player1.id, { include: Profiles})

// hasOneだと最初の１件しか取得できないがhasManyだと全て取得できる
console.log(player.get({ plain: true }))
```

出力結果

```
{ id: 1,
    name: 'player1',
    createdAt: 2019-08-29T04:13:05.396Z,
    updatedAt: 2019-08-29T04:13:05.396Z,
    profiles:
    [ { id: 1,
        name: 'profile1',
        createdAt: 2019-08-29T04:13:05.412Z,
        updatedAt: 2019-08-29T04:13:05.412Z,
        player_id: 1 },
    { id: 2,
        name: 'profile2',
        createdAt: 2019-08-29T04:13:05.428Z,
        updatedAt: 2019-08-29T04:13:05.428Z,
        player_id: 1 } ] }
```


# belongsToMany のサンプル
```Javascript
// model(テーブル)を定義
// 主キー'id'や'createdAt'、'updatedAt'は勝手に定義されるので記述不要
const Players = sequelize.define('players', {
    name: Sequelize.STRING,
});
const Teams = sequelize.define('teams', {
    name: Sequelize.STRING,
});
const TeamPlayers = sequelize.define('team_players', {
    name: Sequelize.STRING,
});

// 中間テーブル'TeamPlayers'を作成し、そこに'team_id'を追加する
Players.belongsToMany(Teams, { through: TeamPlayers, foreignKey: 'team_id' }) 
// 中間テーブル'TeamPlayers'を作成し、そこに'team_id'を追加する
Teams.belongsToMany(Players, { through: TeamPlayers, foreignKey: 'player_id' })

await TeamPlayers.sync({ force: true });
await Players.sync({ force: true });
await Teams.sync({ force: true });

// レコードのcreate
const team1 = await Teams.create({ name: 'team1' });
const team2 = await Teams.create({ name: 'team2' });
const player1 = await Players.create({ name: 'player1' });
player1.addTeam(team1,{ status: 'started' })
player1.addTeam(team2,{ status: 'started' })
const player2 = await Players.create({ name: 'player2' });
player2.addTeam(team1,{ status: 'started' })
player2.addTeam(team2,{ status: 'started' })

// レコード取得
const player = await Players.findById(player1.id, { include: Teams })
console.log(player.get({ plain: true }))

const team= await Teams.findById(team1.id, { include: Players})
console.log(team.get({ plain: true }))
```

出力結果

```
{ id: 1,
  name: 'player1',
  createdAt: 2019-08-29T05:56:11.200Z,
  updatedAt: 2019-08-29T05:56:11.200Z,
  teams:
   [ { id: 1,
       name: 'team1',
       createdAt: 2019-08-29T05:56:11.168Z,
       updatedAt: 2019-08-29T05:56:11.168Z,
       team_players: [Object] },
     { id: 2,
       name: 'team2',
       createdAt: 2019-08-29T05:56:11.184Z,
       updatedAt: 2019-08-29T05:56:11.184Z,
       team_players: [Object] } ] }

{ id: 1,
  name: 'team1',
  createdAt: 2019-08-29T05:56:11.168Z,
  updatedAt: 2019-08-29T05:56:11.168Z,
  players:
   [ { id: 1,
       name: 'player1',
       createdAt: 2019-08-29T05:56:11.200Z,
       updatedAt: 2019-08-29T05:56:11.200Z,
       team_players: [Object] },
     { id: 2,
       name: 'player2',
       createdAt: 2019-08-29T05:56:11.217Z,
       updatedAt: 2019-08-29T05:56:11.217Z,
       team_players: [Object] } ] }
```




