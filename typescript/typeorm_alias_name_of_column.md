typeormのモデル定義で、DB物理名と別名のモデル名をつける

# tl;dr
`@Entity`や`@Column`デコレータの引数のオブジェクトに指定すると、DB物理名を指定できます。

# サンプル
`@Entity`や`@Column`デコレータの引数のオブジェクトの`name`属性にDBの物理名を指定しています。

```Typescript
import { Entity, PrimaryGeneratedColumn, Column } from "typeorm";

@Entity({ name: "m_user" })
export class MPsIdIbl {
  @PrimaryGeneratedColumn({name: "id"})
  readonly id: number;

  @Column({ name: "user_name" })
  userName: string;

}
```

