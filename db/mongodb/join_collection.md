# MongoDBでコレクション（テーブル）をJOIN


SQLの場合
'''
SELECT * FROM users 
LEFT JOIN groups AS grps ON users.groups = groups._id
'''

MongoDB 3.2 〜
'''
db.getCollection('users').aggregate([
  {
    $lookup: {
      from: "groups",
      localField: "groups",
      foreignField: "_id",
      as: "grps"
    }
  },
])
'''

MongoDB 3.6 〜
'''
db.getCollection('users').aggregate([
  {
    $lookup: {
      from: "groups",
      let: { user_groups: "$groups" },
      pipeline: [
        { $match:
          { $expr:
            { $and:
              [
                { $eq: [ "$_id",  "$$user_groups" ] },
              ]
            }
          }
        },
      ],
      as: "grps"
    }
  },
])
'''

