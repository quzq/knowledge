typeORM で、複合主キーによるリレーション

@Entity({ name: 'm_psdspedittbl' })
export class MPsDspEditTbl {
  @PrimaryGeneratedColumn()
  readonly id: number;

  @Column({ name: 'edtjancode', type: 'char', length: 30 })
  edtJanCode: string;

  @UpdateDateColumn({ name: 'updateat' })
  updatedDate: Date;

  // relation
  @OneToOne((type) => MPsItemTbl)
  @JoinColumn([
    // {name: {ローカル物理カラム名}, referencedColumnName: {モデルクラスメンバー名}}
    { name: 'edtjancode', referencedColumnName: 'itmJanCode' },
    //{ name: 'locale_id', referencedColumnName: 'locale_id' },
  ])
  @JoinTable()
  item: MPsItemTbl;
}


    const mPsDspEditTbl = await cn.getRepository(MPsDspEditTbl).find({ relations: ['item'], where: { edtStrCode: strCode, edtCtgCode: ctgCode, edtMkDate: mkDate } });



https://orkhan.gitbook.io/typeorm/docs/relations