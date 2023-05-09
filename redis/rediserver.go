package rediserver

import (
	"fmt"
	"goredis/config"
	"goredis/mypkg"
	"strconv"

	"github.com/fragmentization/mahonia"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var rds *redis.Client

var zsetkey string

var roomgameidconfig int

// 定义一个RedisSingleObj结构体
type RedisSingleObj struct {
	Redis_host string
	Redis_port uint16
	Redis_auth string
	Database   int
	Db         *redis.Client
}

// 结构体InitSingleRedis方法: 用于初始化redis数据库
func (r *RedisSingleObj) InitSingleRedis() (err error) {
	// Redis连接格式拼接
	redisAddr := fmt.Sprintf("%s:%d", r.Redis_host, r.Redis_port)
	// Redis 连接对象: NewClient将客户端返回到由选项指定的Redis服务器。
	r.Db = redis.NewClient(&redis.Options{
		Addr:     redisAddr,    // redis服务ip:port
		Password: r.Redis_auth, // redis的认证密码
		DB:       r.Database,   // 连接的database库
		PoolSize: 10000,        // 连接池
	})
	logrus.Info("Connecting Redis : ", redisAddr)

	// 验证是否连接到redis服务端
	res, err := r.Db.Ping().Result()
	if err != nil {
		logrus.Error("Connect Failed! Err: ", err.Error())
		return err
	} else {
		logrus.Info("Connect Successful! Ping => ", res)
		return nil
	}
}

func RedisInit() {
	dbredis := config.Config_Redis
	// 实例化RedisSingleObj结构体
	conn := &mypkg.RedisSingleObj{
		Redis_host: dbredis.RedisAddr,
		Redis_port: uint16(dbredis.RedisPort),
		Redis_auth: dbredis.RedisPass,
		Database:   dbredis.RedisName,
	}

	// 初始化连接 Single Redis 服务端
	err := conn.InitSingleRedis()
	if err != nil {
		logrus.Error("Init Redis Error :", err.Error())
		return
	}
	rds = conn.Db
}

func keyexist() (isexist bool) {
	zsetkey = GetRedisKey(roomgameidconfig)
	_, e := rds.Exists(zsetkey).Result()
	if e != nil {
		return false
	}
	return true
}

// 排行榜加入/更新数据
func AddRank(roomgameid int, members redis.Z) {
	// ZADD
	zsetkey = GetRedisKey(roomgameid)
	num, err := rds.ZAdd(zsetkey, members).Result()
	if err != nil {
		logrus.Error("zadd failed, err:", err.Error())
		return
	}
	logrus.Info("zadd succ num=", num)
}

// 踢出排行榜
func RmvRank(roomgameid int, members redis.Z) {
	zsetkey = GetRedisKey(roomgameid)
	// if rds.ZRem(zsetkey, members.Member).Val() > 0 {
	// 	logrus.Info(members.Member, "had removed from rank")
	// }
}

func GetRedisKey(roomgameid int) (key string) {
	// var t time.Time = time.Now()
	// timestr := fmt.Sprintf("%d%d%d", t.Year(), t.Month(), t.Day())
	round := GetMatchRound(roomgameid)
	zsetkeyc := fmt.Sprintf("tbmatch_rank:%d:%d", roomgameid, round)
	return zsetkeyc
}

func GetRankNum(z redis.Z) (rank int, err error) {
	// zsetkey := "tbmatch_rank:20221129"
	if !keyexist() {
		return
	}
	ranknum, err := rds.ZRevRank(zsetkey, z.Member.(string)).Result()
	if err != nil {
		logrus.Error("获取Rank失败,err:", err.Error())
		return
	}
	return int(ranknum) + 1, nil
}

/*
获取第多少名的信息
*/
func GetMatchN(n int) (numid int, rank int, score float64, err error) {
	// key := "tbmatch_rank:20221129"
	if !keyexist() {
		return
	}
	topnret, err := rds.ZRevRangeWithScores(zsetkey, (int64)(n-1), (int64)(n-1)).Result()
	if err != nil {
		logrus.Error("获取top", n, "失败,err:", err.Error())
		return -1, -1, -1, err
	}
	if len(topnret) <= 0 {
		return -1, -1, -1, err
	}
	num, e := GetRankNum(topnret[0])
	if e != nil {
		logrus.Error("获取top", n, "失败,err:", e.Error())
		return -1, -1, -1, e
	}
	numidp, _ := strconv.Atoi(topnret[0].Member.(string))
	// fmt.Printf("名次=%d,id=%s\n", num, topnret[0].Member)
	return numidp, num, topnret[0].Score, nil
}

func GetMatchNTest(n int) (numid int, rank int, score float64, err error) {
	key := "tbmatch_rank:38025:15"
	// if !keyexist() {
	// 	return
	// }
	topnret, err := rds.ZRevRangeWithScores(key, (int64)(n-1), (int64)(n-1)).Result()
	if err != nil {
		logrus.Error("获取top", n, "失败,err:", err.Error())
		return -1, -1, -1, err
	}
	if len(topnret) <= 0 {
		return -1, -1, -1, err
	}
	num, e := GetRankNumTest(topnret[0])
	if e != nil {
		logrus.Error("获取ranknum", n, "失败,err:", e.Error())
		return -1, -1, -1, e
	}
	numidp, _ := strconv.Atoi(topnret[0].Member.(string))
	// fmt.Printf("名次=%d,id=%s\n", num, topnret[0].Member)
	return numidp, num, topnret[0].Score, nil
}

func GetRankNumTest(z redis.Z) (rank int, err error) {
	// zsetkey := "tbmatch_rank:20221129"
	// if !keyexist() {
	// 	return
	// }
	key := "tbmatch_rank:38025:15"
	ranknum, err := rds.ZRevRank(key, z.Member.(string)).Result()
	if err != nil {
		logrus.Error("获取Rank失败,err:", err.Error())
		return
	}
	return int(ranknum) + 1, nil
}

/*
获取topN的分数与排名
*/
func GetMatchTopN(n int) (info []redis.Z, err error) {
	// key := "tbmatch_rank:20221129"
	if !keyexist() {
		logrus.Error("key not exist:", zsetkey)
		return
	}
	// key := "tbmatch_rank:38025:12"
	topnret, err := rds.ZRevRangeWithScores(zsetkey, 0, int64(n-1)).Result()
	if err != nil {
		logrus.Error("获取top", n, "失败,err:", err.Error())
		return nil, err
	}

	return topnret, nil
}

/*
获取topN的分数与排名
*/
func GetMatchTopNTest(n int) (info []redis.Z, err error) {
	// key := "tbmatch_rank:20221129"
	// if !keyexist() {
	// 	logrus.Error("key not exist:", zsetkey)
	// 	return
	// }
	key := "tbmatch_rank:38025:15"
	topnret, err := rds.ZRevRangeWithScores(key, 0, int64(n-1)).Result()
	if err != nil {
		logrus.Error("获取top", n, "失败,err:", err.Error())
		return nil, err
	}

	return topnret, nil
}

/*
自己的排名
*/
func GetMemberMatchInfo(numid int) (score float64, rank int64, err error) {
	// key := "tbmatch_rank:20221129"
	if !keyexist() {
		return
	}
	inranker, errf := rds.ZLexCount(zsetkey, fmt.Sprintf("(%s", strconv.Itoa(numid)), fmt.Sprintf("(%s", strconv.Itoa(numid))).Result()
	if errf != nil {
		logrus.Error("ZLexCount fail numid=", numid)
		return
	}
	ranknum, err1 := rds.ZRevRank(zsetkey, strconv.Itoa(numid)).Result()
	rankscore, _ := rds.ZScore(zsetkey, strconv.Itoa(numid)).Result()
	fmt.Printf("ranknum=%v,rankscore=%v,inranker=%v\n", ranknum, rankscore, inranker)
	if err1 != nil {
		//改玩家不存在于排行榜
		logrus.Warn("玩家不存在于排行榜,numid=", numid)
		return rankscore, ranknum, err1
	} else {
		return rankscore, ranknum + 1, err1
	}
}

func ConvertToString(src string, srcCode string, tagCode string) string {

	srcCoder := mahonia.NewDecoder(srcCode)

	srcResult := srcCoder.ConvertString(src)

	tagCoder := mahonia.NewDecoder(tagCode)

	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)

	result := string(cdata)

	return result

}

func GameidInit() {
	gameconfig := config.Config_Game
	roomgameidconfig = gameconfig.RoomGameid
}

func MatchRoundAdd() {
	roundkey := getRoundKey(roomgameidconfig)
	value, err1 := rds.Get(roundkey).Result()
	if err1 != nil {
		logrus.Error("get roundkey fail err1=", err1.Error())
	}
	vint, _ := strconv.Atoi(value)
	vint++
	rds.Set(roundkey, vint, 0)
	logrus.Info("roundkey:", roundkey)
}
func getRoundKey(roomgameid int) string {
	roundkey := fmt.Sprintf("tb_match_round:%d", roomgameid)
	return roundkey
}
func GetMatchRound(roomgameid int) int {
	roundkey := getRoundKey(roomgameid)
	logrus.Info("roundkey=", roundkey)
	value, err1 := rds.Get(roundkey).Result()
	if err1 != nil {
		logrus.Error("get roundkey fail err1=", err1.Error())
		return -1
	}
	vint, _ := strconv.Atoi(value)
	return vint
}

func TestRedis() {
	testkey := "xtask_hd_rl_38000_1"
	_, err := rds.ZAdd(testkey, redis.Z{Score: 1675353600, Member: "48000808"}).Result()
	if err != nil {
		logrus.Error("zadd failed, err:", err.Error())
		return
	}
	// logrus.Info("zadd succ num=", num)
}
