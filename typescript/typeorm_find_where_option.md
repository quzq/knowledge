TypeORMの比較演算子一覧

# Equal ( = )
`SELECT * FROM "post" WHERE "title" = 'About #1'` を実行するコード：

```Typescript
const loadedPosts = await connection.getRepository(Post).find({
    title: "About #1"
})
```

# Not 
`SELECT * FROM "post" WHERE "title" != 'About #1'` を実行するコード：

```Typescript
import {Not} from "typeorm";

const loadedPosts = await connection.getRepository(Post).find({
    title: Not("About #1")
})
```
# LessThan ( < )
`SELECT * FROM "post" WHERE "likes" < 10` を実行するコード：
```Typescript
import {LessThan} from "typeorm";

const loadedPosts = await connection.getRepository(Post).find({
    likes: LessThan(10)
});
```

# LessThanOrEqual ( <= )
`SELECT * FROM "post" WHERE "likes" <= 10` を実行するコード：
```Typescript
import {LessThanOrEqual} from "typeorm";

const loadedPosts = await connection.getRepository(Post).find({
    likes: LessThanOrEqual(10)
});
```

# MoreThan ( > )
`SELECT * FROM "post" WHERE "likes" > 10` を実行するコード：
```Typescript
import {MoreThan} from "typeorm";

const loadedPosts = await connection.getRepository(Post).find({
    likes: MoreThan(10)
});
```

# MoreThanOrEqual ( >= )
`SELECT * FROM "post" WHERE "likes" >= 10` を実行するコード：
```Typescript
import {MoreThanOrEqual} from "typeorm";

const loadedPosts = await connection.getRepository(Post).find({
    likes: MoreThanOrEqual(10)
});
```

# Like
`SELECT * FROM "post" WHERE "title" LIKE '%out #%'` を実行するコード：
```Typescript
import {Like} from "typeorm";

const loadedPosts = await connection.getRepository(Post).find({
    title: Like("%out #%")
});
```

# Between
`SELECT * FROM "post" WHERE "likes" BETWEEN 1 AND 10`
```Typescript
import {Between} from "typeorm";

const loadedPosts = await connection.getRepository(Post).find({
    likes: Between(1, 10)
});
```

# In
`SELECT * FROM "post" WHERE "title" IN ('About #2','About #3')`
```Typescript
import {In} from "typeorm";

const loadedPosts = await connection.getRepository(Post).find({
    title: In(["About #2", "About #3"])
});
```

# Any
`SELECT * FROM "post" WHERE "title" = ANY(['About #2','About #3'])`
will execute following query (Postgres notation):
```Typescript
import {Any} from "typeorm";

const loadedPosts = await connection.getRepository(Post).find({
    title: Any(["About #2", "About #3"])
});
```

# IsNull
`SELECT * FROM "post" WHERE "title" IS NULL`
```Typescript
import {IsNull} from "typeorm";

const loadedPosts = await connection.getRepository(Post).find({
    title: IsNull()
});
```

# Raw
`SELECT * FROM "post" WHERE "likes" = "dislikes" - 4`
```Typescript
import {Raw} from "typeorm";

const loadedPosts = await connection.getRepository(Post).find({
    likes: Raw("dislikes - 4")
});
```

# 📖参考
https://github.com/typeorm/typeorm/blob/master/docs/find-options.md